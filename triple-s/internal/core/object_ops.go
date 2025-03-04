package core

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
)

func CreateObject(dirPath, bucket, object string, r *http.Request) (*os.File, int, error) {
	objectFile, err := os.Create(dirPath + bucket + "/" + object)
	if err != nil {
		return nil, 0, err
	}
	objectSize, err := io.Copy(objectFile, r.Body)
	if err != nil {
		return nil, 0, err
	}

	return objectFile, int(objectSize), nil
}

func DeleteObjectAndMeta(dirPath, bucket, object string) error {
	var ObjectExisted bool = false

	csvObjectsFile, err := os.Open(dirPath + bucket + "/" + "objects.csv")
	if err != nil {
		return err
	}

	csvObjectsReader := csv.NewReader((csvObjectsFile))
	defer csvObjectsFile.Close()

	csvObjectsRecords, err := csvObjectsReader.ReadAll()
	if err != nil {
		return err
	}

	var filteredRecords [][]string
	for _, row := range csvObjectsRecords {
		if row[0] == object {
			// remove file
			err := os.Remove(dirPath + bucket + "/" + object)
			if err != nil {
				if os.IsNotExist(err) {
					return ErrObjectNotExist
				}
				return err
			}
			ObjectExisted = true
			continue
		}

		filteredRecords = append(filteredRecords, row)
	}

	if !ObjectExisted {
		return ErrObjectNotExist
		// err := os.Remove(dirPath + bucket + "/" + object)
		// 	if err != nil {
		//       if os.IsNotExist(err){
		//         return ErrObjectNotExist
		//       }
		// 		return err
		// 	}
	}

	if len(filteredRecords) == 1 {
		err := os.Remove(dirPath + bucket + "/" + "objects.csv")
		if err != nil {
			return err
		}
		return nil
	}
	csvObjectsFile, err = os.OpenFile(
		dirPath+bucket+"/"+"objects.csv",
		os.O_WRONLY|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}

	csvObjectsWriter := csv.NewWriter(csvObjectsFile)

	err = csvObjectsWriter.WriteAll(filteredRecords)
	if err != nil {
		return err
	}

	return nil
}

func CheckObjectExist(dirPath, bucket, object string) error {
	_, err := os.Stat(dirPath + bucket + "/" + object)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrObjectNotExist
		} else {
			return err
		}
	}
	return nil
}

func CheckBlackList(endpoint string, blackList []string) error {
	for _, v := range blackList {
		if endpoint == v {
			return ErrObjectNameNotAllowed
		}
	}

	return nil
}
