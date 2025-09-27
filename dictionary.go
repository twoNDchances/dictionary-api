package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Dictionary struct {
	Id        uint   `gorm:"primarykey" uri:"id" json:"id"`
	Key       string `gorm:"size:255,uniqueIndex" uri:"key" json:"key"`
	Value     *string `uri:"value" json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *Dictionary) isExist() (bool, error) {
	var result Dictionary
    tx := DB.Where(d).First(&result)
    if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
        return false, nil
    }
    if tx.Error != nil {
        return false, tx.Error
    }
    return true, nil
}

func (d *Dictionary) create() (*Dictionary, int, error) {
    exist, err := d.isExist()
    if err != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("error when check exists before creating: %w", err)
    }
    if exist {
        return nil, http.StatusBadRequest, fmt.Errorf("dictionary already exists")
    }
    tx := DB.Create(d)
    if tx.Error != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("error when creating dictionary: %w", tx.Error)
    }
    return d, http.StatusOK, nil
}

func (d *Dictionary) list() ([]Dictionary, int, error) {
    var dicts []Dictionary
    tx := DB.Find(&dicts)
    if tx.Error != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("error when listing dictionaries: %s", tx.Error.Error())
    }
    return dicts, http.StatusOK, nil
}

func (d *Dictionary) show() (*Dictionary, int, error) {
	exist, err := d.isExist()
    if err != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("error when check exists before showing: %w", err)
    }
    if !exist {
        return nil, http.StatusNotFound, fmt.Errorf("dictionary not exists")
    }
    var dict Dictionary
    tx := DB.Where(d).First(&dict)
    if tx.Error != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("error when showing dictionary: %s", tx.Error.Error())
    }
    return &dict, http.StatusOK, nil
}

func (d *Dictionary) update() (*Dictionary, int, error) {
	exist, err := d.isExist()
    if err != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("error when check exists before updating: %w", err)
    }
    if !exist {
        return nil, http.StatusNotFound, fmt.Errorf("dictionary not exists")
    }
    tx := DB.Model(d).Updates(*d)
    if tx.Error != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("error when updating dictionary: %s", tx.Error.Error())
    }
    return d, http.StatusOK, nil
}

func (d *Dictionary) delete() (*Dictionary, int, error) {
	exist, err := d.isExist()
    if err != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("error when check exists before deleting: %w", err)
    }
    if !exist {
        return nil, http.StatusNotFound, fmt.Errorf("dictionary not exists")
    }
    tx := DB.Delete(d)
    if tx.Error != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("error when deleting dictionary: %s", tx.Error.Error())
    }
    return d, http.StatusOK, nil
}
