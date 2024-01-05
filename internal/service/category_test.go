// category_service_test.go

package service_test

import (
	"context"
	"spun/internal/model"
	"spun/internal/service"
	"spun/mocks"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestCreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock the repository and audit trail service
	mockRepo := mocks.NewMockCategoryRepository(ctrl)

	// Set up the expected category and parameters
	newCategory := &model.Category{
		Name:        "Test Category",
		Description: "Test Description",
	}

	mockRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func(c *model.Category) (*model.Category, error) {
		// Simulate setting the ID of the new category
		newCategory.ID = 123
		return newCategory, nil
	}).Times(1)

	// Initialize the CategoryService with the mock repository and audit trail service
	categoryService := service.NewCategoryService(mockRepo, nil)

	// Create the parameters for creating a category
	params := &service.CreateCategoryParam{
		Name:        "Test Category",
		Description: "Test Description",
	}

	// Call the method you want to test
	result, err := categoryService.CreateCategory(context.Background(), params)

	// Assert the results
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.ID != 123 || result.Name != newCategory.Name || result.Description != newCategory.Description {
		t.Errorf("Expected category with ID 123 and correct name and description, got %v", result)
	}
}

func TestViewCategory(t *testing.T) {
	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock CategoryRepository
	mockRepo := mocks.NewMockCategoryRepository(ctrl)

	// Set up expected calls and return values
	expectedCategory := &model.Category{ID: 123, Name: "Test Category"}
	mockRepo.EXPECT().GetByID(int64(123)).Return(expectedCategory, nil)

	// Initialize the CategoryService with the mock CategoryRepository
	categoryService := service.NewCategoryService(mockRepo, nil) // passing nil for svcAudit as it's not used here

	// Call the method you want to test
	param := service.ViewCategoryParam{IDParam: service.IDParam{ID: 123}}
	result, err := categoryService.ViewCategory(context.Background(), &param)

	// Assert the results
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.ID != expectedCategory.ID || result.Name != expectedCategory.Name {
		t.Errorf("Expected category %v, got %v", expectedCategory, result)
	}
}
