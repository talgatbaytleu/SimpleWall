package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"poster/internal/dal"
	"poster/pkg/apperrors"
	"poster/pkg/models"
	"poster/pkg/utils"
)

type PostServiceInterface interface {
	CreatePost(imageFile io.ReadCloser, description string, user_idStr string) error
	RetrievePost(post_idStr string) (*io.ReadCloser, *string, error)
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
	_, err = SendRequest("PUT", &s3reqUrl, http.NoBody)

	// Creating image on S3
	image_url := utils.GenUniqueString()
	s3reqUrl = s3url + "bucket" + user_idStr + image_url

	resp, err := SendRequest("PUT", &s3reqUrl, imageFile)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return apperrors.ErrS3NotCreated
	}

	defer func() {
		if err != nil {
			SendRequest("DELETE", &s3reqUrl, imageFile)
		}
	}()

	err = s.postDal.InsertPost(user_id, description, image_url)
	if err != nil {
		return err
	}

	return nil
}

func (s *postService) RetrievePost(post_idStr string) (*io.ReadCloser, *string, error) {
	post_id, err := strconv.Atoi(post_idStr)
	if err != nil {
		return nil, nil, err
	}

	// Getting ImageLink from DB
	jsonPostPtr, err := s.postDal.SelectPost(post_id)
	if err != nil {
		return nil, nil, err
	}

	// Unmarshal to struct
	var postStruct models.PostType
	err = json.Unmarshal([]byte(*jsonPostPtr), &postStruct)

	// Getting imageFile from S3
	user_id := strconv.Itoa(postStruct.UserID)
	s3reqUrl := s3url + user_id + postStruct.ImageLink
	resp, err := SendRequest("GET", &s3reqUrl, nil)

	return &resp.Body, &postStruct.ImageLink, nil
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
	resp, err := SendRequest("PUT", &s3reqUrl, imageFile)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return apperrors.ErrS3NotCreated
	}

	// Rollback S3 and DB data if any error occurres
	defer func() {
		if err != nil {
			SendRequest("DELETE", &s3reqUrl, imageFile)
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
	resp, err = SendRequest("DELETE", &s3reqUrl, http.NoBody)
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
	oldImageResp, err := SendRequest("GET", &s3reqUrl, http.NoBody)
	if err != nil {
		return nil
	}

	// if postStruct.UserID != user_id {
	// 	return apperrors.ErrNotAllowed
	// }

	// Deleting image from S3
	resp, err := SendRequest("DELETE", &s3reqUrl, http.NoBody)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return apperrors.ErrS3NotDeleted
	}

	// Rollback S3 image
	defer func() {
		if err != nil {
			SendRequest("PUT", &s3reqUrl, oldImageResp.Body)
		}
	}()

	err = s.postDal.DeletePost(user_id, post_id)
	if err != nil {
		return err
	}

	return nil
}
