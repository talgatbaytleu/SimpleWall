package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"poster/internal/dal"
	"poster/internal/router"
	"poster/pkg/apperrors"
	"poster/pkg/models"
	"poster/pkg/utils"
)

var s3url = "http://localhost:8086/" // SHOULD BE IN .env!!!

type PostServiceInterface interface {
	CreatePost(imageFile io.ReadCloser, description string, user_idStr string) error
	RetrievePost(post_idStr string) (*string, error)
	UpdatePost(
		imageFile io.ReadCloser,
		description string,
		user_idStr string,
		post_idStr string,
	) error
	RemovePost(user_idStr string, post_idStr string) error
}

type postService struct {
	postDal dal.PostDalInterface
}

func NewPostService(postDal dal.PostDalInterface) *postService {
	return &postService{postDal: postDal}
}

func (s *postService) CreatePost(
	imageFile io.ReadCloser,
	description string,
	user_idStr string,
) error {
	user_id, err := strconv.Atoi(user_idStr)
	if err != nil {
		return err
	}
	// Creating a bucket on S3, ignore if it already exist
	s3reqUrl := s3url + "bucket" + user_idStr
	_, err = router.SendRequest("PUT", &s3reqUrl, nil)

	// Creating image on S3
	image_url := utils.GenUniqueString()
	s3reqUrl = s3url + "bucket" + user_idStr + image_url

	resp, err := router.SendRequest("PUT", &s3reqUrl, &imageFile)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return apperrors.ErrS3NotCreated
	}

	defer func() {
		if err != nil {
			router.SendRequest("DELETE", &s3reqUrl, &imageFile)
		}
	}()

	err = s.postDal.InsertPost(user_id, description, image_url)
	if err != nil {
		return err
	}

	return nil
}

func (s *postService) RetrievePost(post_idStr string) (*string, error) {
	post_id, err := strconv.Atoi(post_idStr)
	if err != nil {
		return nil, err
	}

	jsonPostPtr, err := s.postDal.SelectPost(post_id)
	if err != nil {
		return nil, err
	}

	return jsonPostPtr, nil
}

func (s *postService) UpdatePost(
	imageFile io.ReadCloser,
	description string,
	user_idStr string,
	post_idStr string,
) error {
	user_id, err := strconv.Atoi(user_idStr)
	if err != nil {
		return err
	}

	post_id, err := strconv.Atoi(post_idStr)
	if err != nil {
		return err
	}

	// Getting image URL for S3
	jsonPost, err := s.postDal.SelectPost(post_id)
	if err != nil {
		return err
	}

	var postStruct models.PostType
	err = json.Unmarshal([]byte(*jsonPost), &postStruct)

	if postStruct.UserID != user_id {
		return apperrors.ErrNotAllowed
	}

	// Creating new imageURL
	image_url := utils.GenUniqueString()
	s3reqUrl := s3url + user_idStr + image_url

	// Creating new image on S3
	resp, err := router.SendRequest("PUT", &s3reqUrl, &imageFile)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return apperrors.ErrS3NotCreated
	}

	// Rollback S3 and DB data if any error occurres
	defer func() {
		if err != nil {
			router.SendRequest("DELETE", &s3reqUrl, &imageFile)
			s.postDal.UpdateTable(user_id, post_id, postStruct.Description, postStruct.ImageLink)
		}
	}()

	// Updating DB
	err = s.postDal.UpdateTable(user_id, post_id, description, image_url)
	if err != nil {
		return err
	}

	// Deleting image from S3
	s3reqUrl = s3url + user_idStr + postStruct.ImageLink
	resp, err = router.SendRequest("DELETE", &s3reqUrl, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return apperrors.ErrS3NotDeleted
	}

	return nil
}

func (s *postService) RemovePost(user_idStr string, post_idStr string) error {
	user_id, err := strconv.Atoi(user_idStr)
	if err != nil {
		return err
	}

	post_id, err := strconv.Atoi(post_idStr)
	if err != nil {
		return err
	}

	// Getting image ImageLink
	jsonPost, err := s.postDal.SelectPost(post_id)
	if err != nil {
		return err
	}

	var postStruct models.PostType
	err = json.Unmarshal([]byte(*jsonPost), &postStruct)

	s3reqUrl := s3url + user_idStr + postStruct.ImageLink

	// Saving old image for Rollback
	oldImageResp, err := router.SendRequest("GET", &s3reqUrl, nil)
	if err != nil {
		return nil
	}

	// if postStruct.UserID != user_id {
	// 	return apperrors.ErrNotAllowed
	// }

	// Deleting image from S3
	resp, err := router.SendRequest("DELETE", &s3reqUrl, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return apperrors.ErrS3NotDeleted
	}

	// Rollback S3 image
	defer func() {
		if err != nil {
			router.SendRequest("PUT", &s3reqUrl, &oldImageResp.Body)
		}
	}()

	err = s.postDal.DeletePost(user_id, post_id)
	if err != nil {
		return err
	}

	return nil
}
