package models

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Dictionary struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Key       string    `gorm:"size:255;uniqueIndex" json:"key"`
	Value     *string   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func getByID(id uint) (*Dictionary, error) {
	var dict Dictionary
	tx := DB.First(&dict, id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dict, nil
}

func getByKey(key string) (*Dictionary, error) {
	var dict Dictionary
	tx := DB.Where("key = ?", key).First(&dict)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dict, nil
}

func (d *Dictionary) Create() (*Dictionary, int, error) {
	exist, err := getByKey(d.Key)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error when checking existing key: %w", err)
	}
	if exist != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("dictionary with key '%s' already exists", d.Key)
	}

	if err := DB.Create(d).Error; err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error when creating dictionary: %w", err)
	}
	return d, http.StatusCreated, nil
}

func List() ([]Dictionary, int, error) {
	var dicts []Dictionary
	if err := DB.Find(&dicts).Error; err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error when listing dictionaries: %w", err)
	}
	return dicts, http.StatusOK, nil
}

func Show(id uint) (*Dictionary, int, error) {
	dict, err := getByID(id)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error when showing dictionary: %w", err)
	}
	if dict == nil {
		return nil, http.StatusNotFound, fmt.Errorf("dictionary not found")
	}
	return dict, http.StatusOK, nil
}

func (d *Dictionary) Update(id uint) (*Dictionary, int, error) {
	exist, err := getByID(id)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error when checking existing dictionary: %w", err)
	}
	if exist == nil {
		return nil, http.StatusNotFound, fmt.Errorf("dictionary not found")
	}

	if d.Key != "" && d.Key != exist.Key {
		dup, err := getByKey(d.Key)
		if err != nil {
			return nil, http.StatusInternalServerError, fmt.Errorf("error when checking duplicate key: %w", err)
		}
		if dup != nil {
			return nil, http.StatusBadRequest, fmt.Errorf("dictionary with key '%s' already exists", d.Key)
		}
	}

	if err := DB.Model(exist).Updates(d).Error; err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error when updating dictionary: %w", err)
	}

	return exist, http.StatusOK, nil
}

func Delete(id uint) (int, error) {
	dict, err := getByID(id)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error when checking dictionary before delete: %w", err)
	}
	if dict == nil {
		return http.StatusNotFound, fmt.Errorf("dictionary not found")
	}

	if err := DB.Delete(dict).Error; err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error when deleting dictionary: %w", err)
	}
	return http.StatusOK, nil
}
