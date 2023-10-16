package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/clinic-crm/api-gateway/api/models"
	"gitlab.com/clinic-crm/api-gateway/genproto/lab"
	"gitlab.com/clinic-crm/api-gateway/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary 	Create lab
// @Description This api can registr lab
// @Tags 		Lab
// @Accept 		json
// @Produce 	json
// @Param body 	body models.LabModel true "Body"
// @Success 201 {object} models.LabModelResp
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/lab-create [post]
func (h *handlerV1) LabCreate(c *gin.Context) {
	var body models.LabModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabCreate(ctx, &lab.LabCreateReq{
		Id:            uuid.New().String(),
		Name:          body.Name,
		Price:         body.Price,
		Type:          body.Type,
		SubCategoryId: body.SubCategoryId,
	})
	if err != nil {
		h.log.Error("Error creating lab", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.LabModelResp{
		Id:            response.Id,
		Name:          response.Name,
		Price:         response.Price,
		Type:          response.Type,
		SubCategoryId: response.SubCategoryId,
		CreatedAt:     response.CreatedAt,
		UpdatedAt:     response.UpdatedAt,
	})
}

// @Summary 		Get lab
// @Description 	This api can get lab
// @Tags 			Lab
// @Accept 			json
// @Produce         json
// @Param 			filter query models.LabGetReq false "Filter"
// @Success         200			{object}  models.LabModelResp
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/lab-get [get]
func (h *handlerV1) LabGet(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabGet(ctx, &lab.LabGetReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})
	if err != nil {
		h.log.Error("Error getting lab", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.LabModelResp{
		Id:            response.Id,
		Name:          response.Name,
		Price:         response.Price,
		Type:          response.Type,
		SubCategoryId: response.SubCategoryId,
		CreatedAt:     response.CreatedAt,
		UpdatedAt:     response.UpdatedAt,
	})
}

// @Summary 	Find labs
// @Description This api can find labs
// @Tags 		Lab
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.LabsFindReq false "Filter"
// @Success 	200 {object} models.LabsResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/lab-find [get]
func (h *handlerV1) LabsFind(c *gin.Context) {
	var LabsResp models.LabsResp

	req, err := labParams(c)
	if err != nil {
		h.log.Error("Error finding labs", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabsFind(ctx, &lab.LabsFindReq{
		Limit:  req.Limit,
		Page:   req.Page,
		Search: req.Search,
	})
	if err != nil {
		h.log.Error("Error finding labs", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, lab := range response.Labs {
		LabsResp.Labs = append(LabsResp.Labs, &models.LabModelResp{
			Id:            lab.Id,
			Name:          lab.Name,
			Price:         lab.Price,
			Type:          lab.Type,
			SubCategoryId: lab.SubCategoryId,
			CreatedAt:     lab.CreatedAt,
			UpdatedAt:     lab.UpdatedAt,
		})
	}
	LabsResp.Count = response.Count

	c.JSON(http.StatusCreated, LabsResp)
}

func labParams(c *gin.Context) (*models.LabsFindReq, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Param("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	return &models.LabsFindReq{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}, nil
}

// @Summary 	Update lab
// @Description This api can update lab
// @Tags 		Lab
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param   	body body models.LabModel true "Body"
// @Success 	200 {object} models.LabModelResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/lab-update/{id}  [post]
func (h *handlerV1) LabUpdate(c *gin.Context) {
	var body models.LabModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error updating lab", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabUpdate(ctx, &lab.LabUpdateReq{
		Id:    c.Param("id"),
		Name:  body.Name,
		Price: body.Price,
		Type:  body.Type,
	})
	if err != nil {
		h.log.Error("Error updating lab", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.LabModelResp{
		Id:            response.Id,
		Name:          response.Name,
		Price:         response.Price,
		Type:          response.Type,
		SubCategoryId: response.SubCategoryId,
		CreatedAt:     response.CreatedAt,
		UpdatedAt:     response.UpdatedAt,
	})
}

// @Summary 	Delete lab
// @Description This api can delete lab
// @Tags 		Lab
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/lab-delete/{id}  [delete]
func (h *handlerV1) LabDelete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.LabService().LabDelete(ctx, &lab.LabId{
		Id: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting lab", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	Create aparat
// @Description This api can registr aparat
// @Tags 		Aparat
// @Accept 		json
// @Produce 	json
// @Param body 	body models.CreateAparat true "Body"
// @Success 201 {object} models.AparatModelResp
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/aparat-create [post]
func (h *handlerV1) AparatCreate(c *gin.Context) {
	var body models.CreateAparat

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating aparat", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatCreate(ctx, &lab.AparatCreateReq{
		Id:            uuid.New().String(),
		Name:          body.Name,
		Price:         body.Price,
		Type:          body.Type,
		SubCategoryId: body.SubCategoryId,
	})
	if err != nil {
		h.log.Error("Error creating aparat", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.AparatModelResp{
		Id:            response.Id,
		Name:          response.Name,
		Price:         response.Price,
		Type:          response.Type,
		SubCategoryId: response.SubCategoryId,
		CreatedAt:     response.CreatedAt,
		UpdatedAt:     response.UpdatedAt,
	})
}

// @Summary 		Get aparat
// @Description 	This api can get lab
// @Tags 			Aparat
// @Accept 			json
// @Produce         json
// @Param 			filter query models.AparatGetReq false "Filter"
// @Success         200			{object}  models.AparatModelResp
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/aparat-get [get]
func (h *handlerV1) AparatGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatGet(ctx, &lab.AparatGetReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})
	if err != nil {
		h.log.Error("Error getting aparats", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.AparatModelResp{
		Id:            response.Id,
		Name:          response.Name,
		Price:         response.Price,
		Type:          response.Type,
		SubCategoryId: response.SubCategoryId,
		CreatedAt:     response.CreatedAt,
		UpdatedAt:     response.UpdatedAt,
	})
}

// @Summary 	Find aparat
// @Description This api can find labs
// @Tags 		Aparat
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.AparatFindReq false "Filter"
// @Success 	200 {object} models.AparatsResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/aparat-find [get]
func (h *handlerV1) AparatsFind(c *gin.Context) {
	var AparatsResp models.AparatsResp

	req, err := labParams(c)
	if err != nil {
		h.log.Error("Error finding aparats", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatsFind(ctx, &lab.AparatsFindReq{
		Limit:  req.Limit,
		Page:   req.Page,
		Search: req.Search,
	})
	if err != nil {
		h.log.Error("Error finding aparats", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, lab := range response.Aparats {
		AparatsResp.Aparats = append(AparatsResp.Aparats, &models.AparatModelResp{
			Id:            lab.Id,
			Name:          lab.Name,
			Price:         lab.Price,
			Type:          lab.Type,
			SubCategoryId: lab.SubCategoryId,
			CreatedAt:     lab.CreatedAt,
			UpdatedAt:     lab.UpdatedAt,
		})
	}
	AparatsResp.Count = response.Count

	c.JSON(http.StatusCreated, AparatsResp)
}

// @Summary 	Update aparat
// @Description This api can update lab
// @Tags 		Aparat
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param   	body body models.UpdateAparat true "Body"
// @Success 	200 {object} models.AparatModelResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/aparat-update/{id}  [post]
func (h *handlerV1) AparatsUpdate(c *gin.Context) {
	var body models.UpdateAparat

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error updating lab", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatsUpdate(ctx, &lab.AparatUpdateReq{
		Id:    c.Param("id"),
		Name:  body.Name,
		Price: body.Price,
		Type:  body.Type,
	})
	if err != nil {
		h.log.Error("Error updating lab", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.LabModelResp{
		Id:        response.Id,
		Name:      response.Name,
		Price:     response.Price,
		Type:      response.Type,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 	Delete aparat
// @Description This api can delete lab
// @Tags 		Aparat
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/aparat-delete/{id}  [delete]
func (h *handlerV1) AparatsDelete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.LabService().AparatsDelete(ctx, &lab.AparatId{
		Id: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting lab", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	Create lab category
// @Description This api can registr lab categoty
// @Tags 		Lab Category
// @Accept 		json
// @Produce 	json
// @Param body 	body models.CategoryModel true "Body"
// @Success 201 {object} models.CategoryModelResp
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/lab-category-create [post]
func (h *handlerV1) LabCategoryCreate(c *gin.Context) {
	var body models.CategoryModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating lab category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabCategoryCreate(ctx, &lab.Category{
		Id:   uuid.New().String(),
		Name: body.Name,
	})
	if err != nil {
		h.log.Error("Error creating lab category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CategoryModelResp{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 		Get lab category
// @Description 	This api can get lab category
// @Tags 			Lab Category
// @Accept 			json
// @Produce         json
// @Param 			filter query models.CategoryGetReqModel false "Filter"
// @Success         200			{object}  models.CategoryModelResp
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/lab-category-get [get]
func (h *handlerV1) LabCategoryGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabCategoryGet(ctx, &lab.CategoryGetReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})
	if err != nil {
		h.log.Error("Error getting lab category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CategoryModelResp{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 	Find lab category
// @Description This api can find lab categories
// @Tags 		Lab Category
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.CategoryFindReqModel false "Filter"
// @Success 	200 {object} models.CategoriesResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/lab-category-find [get]
func (h *handlerV1) LabCategoryFind(c *gin.Context) {
	var CategoriesResp models.CategoriesResp

	req, err := labParams(c)
	if err != nil {
		h.log.Error("Error finding lab category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabCategoryFind(ctx, &lab.CategoryFindReq{
		Limit:  req.Limit,
		Page:   req.Page,
		Search: req.Search,
	})
	if err != nil {
		h.log.Error("Error finding lab category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, lab := range response.Info {
		CategoriesResp.Category = append(CategoriesResp.Category, &models.CategoryModelResp{
			Id:        lab.Id,
			Name:      lab.Name,
			CreatedAt: lab.CreatedAt,
			UpdatedAt: lab.UpdatedAt,
		})
	}
	CategoriesResp.Count = response.Count

	c.JSON(http.StatusCreated, CategoriesResp)
}

// @Summary 	Update lab category
// @Description This api can update lab catigory
// @Tags 		Lab Category
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param   	body body models.CategoryModel true "Body"
// @Success 	200 {object} models.CategoryModelResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/lab-category-update/{id}  [post]
func (h *handlerV1) LabCategoryUpdate(c *gin.Context) {
	var body models.CategoryModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error updating lab category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabCategoryUpdate(ctx, &lab.Category{
		Id:   c.Param("id"),
		Name: body.Name,
	})
	if err != nil {
		h.log.Error("Error updating lab category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CategoryModelResp{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 	Delete lab category
// @Description This api can delete lab category
// @Tags 		Lab Category
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/lab-category-delete/{id}  [delete]
func (h *handlerV1) LabCategoryDelete(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.LabService().LabCategoryDelete(ctx, &lab.CategoryId{
		Id: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting lab category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	Create aparat category
// @Description This api can registr aparat categoty
// @Tags 		Aparat Category
// @Accept 		json
// @Produce 	json
// @Param body 	body models.CategoryModel true "Body"
// @Success 201 {object} models.CategoryModelResp
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/aparat-category-create [post]
func (h *handlerV1) AparatCategoryCreate(c *gin.Context) {
	var body models.CategoryModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating aparat category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatCategoryCreate(ctx, &lab.Category{
		Id:   uuid.New().String(),
		Name: body.Name,
	})
	if err != nil {
		h.log.Error("Error creating aparat category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CategoryModelResp{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 		Get aparat category
// @Description 	This api can get aparat category
// @Tags 			Aparat Category
// @Accept 			json
// @Produce         json
// @Param 			filter query models.CategoryGetReqModel false "Filter"
// @Success         200			{object}  models.CategoryModelResp
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/aparat-category-get [get]
func (h *handlerV1) AparatCategoryGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatCategoryGet(ctx, &lab.CategoryGetReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})
	if err != nil {
		h.log.Error("Error getting aparat category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CategoryModelResp{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 	Find aparat category
// @Description This api can find aparat categories
// @Tags 		Aparat Category
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.CategoryFindReqModel false "Filter"
// @Success 	200 {object} models.CategoriesResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/aparat-category-find [get]
func (h *handlerV1) AparatCategoryFind(c *gin.Context) {
	var CategoriesResp models.CategoriesResp

	req, err := labParams(c)
	if err != nil {
		h.log.Error("Error finding aparat category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatCategoryFind(ctx, &lab.CategoryFindReq{
		Limit:  req.Limit,
		Page:   req.Page,
		Search: req.Search,
	})
	if err != nil {
		h.log.Error("Error finding aparat category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, lab := range response.Info {
		CategoriesResp.Category = append(CategoriesResp.Category, &models.CategoryModelResp{
			Id:        lab.Id,
			Name:      lab.Name,
			CreatedAt: lab.CreatedAt,
			UpdatedAt: lab.UpdatedAt,
		})
	}
	CategoriesResp.Count = response.Count

	c.JSON(http.StatusCreated, CategoriesResp)
}

// @Summary 	Update aparat category
// @Description This api can update aparat catigory
// @Tags 		Aparat Category
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param   	body body models.CategoryModel true "Body"
// @Success 	200 {object} models.CategoryModelResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/aparat-category-update/{id}  [post]
func (h *handlerV1) AparatCategoryUpdate(c *gin.Context) {
	var body models.CategoryModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error updating aparat category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatCategoryUpdate(ctx, &lab.Category{
		Id:   c.Param("id"),
		Name: body.Name,
	})
	if err != nil {
		h.log.Error("Error updating aparat category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CategoryModelResp{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 	Delete aparat category
// @Description This api can delete aparat category
// @Tags 		Aparat Category
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/aparat-category-delete/{id}  [delete]
func (h *handlerV1) AparatCategoryDelete(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.LabService().AparatCategoryDelete(ctx, &lab.CategoryId{
		Id: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting aparat category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	Create Lab sub category
// @Description This api can registr Lab sub category create
// @Tags 		Lab sub category
// @Accept 		json
// @Produce 	json
// @Param body 	body models.SubCategoryModel true "Body"
// @Success 201 {object} models.SubCategoryResModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/lab-sub-category-create [post]
func (h *handlerV1) LabSubCategoryCreate(c *gin.Context) {
	var body models.SubCategoryModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating lab sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabSubCategoryCreate(ctx, &lab.SubCategory{
		Id:         uuid.New().String(),
		Name:       body.Name,
		CategoryId: body.CategoryId,
	})
	if err != nil {
		h.log.Error("Error creating lab sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.SubCategoryResModel{
		Id:         response.Id,
		Name:       response.Name,
		CategoryId: response.CategoryId,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	})
}

// @Summary 		Get Lab sub category
// @Description 	This api can get Lab sub category create
// @Tags 			Lab sub category
// @Accept 			json
// @Produce         json
// @Param 			filter query models.CategoryGetReqModel false "Filter"
// @Success         200			{object}  models.SubCategoryResModel
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/lab-sub-category-get [get]
func (h *handlerV1) LabSubCategoryGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabSubCategoryGet(ctx, &lab.CategoryGetReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})
	if err != nil {
		h.log.Error("Error getting lab sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.SubCategoryResModel{
		Id:         response.Id,
		Name:       response.Name,
		CategoryId: response.CategoryId,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	})
}

// @Summary 	Find Lab sub category
// @Description This api can find lab sub categories
// @Tags 		Lab sub category
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.CategoryFindReqModel false "Filter"
// @Success 	200 {object} models.SubCategoriesResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/lab-sub-category-find [get]
func (h *handlerV1) LabSubCategoryFind(c *gin.Context) {
	var CategoriesResp models.SubCategoriesResp

	req, err := labParams(c)
	if err != nil {
		h.log.Error("Error finding lab sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabSubCategoryFind(ctx, &lab.SubCategoryFindReq{
		Limit:      req.Limit,
		Page:       req.Page,
		CategoryId: req.Search,
	})
	if err != nil {
		h.log.Error("Error finding lab sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, lab := range response.Info {
		CategoriesResp.Category = append(CategoriesResp.Category, &models.SubCategoryResModel{
			Id:         lab.Id,
			Name:       lab.Name,
			CategoryId: lab.CategoryId,
			CreatedAt:  lab.CreatedAt,
			UpdatedAt:  lab.UpdatedAt,
		})
	}
	CategoriesResp.Count = response.Count

	c.JSON(http.StatusCreated, CategoriesResp)
}

// @Summary 	Update Lab sub category
// @Description This api can update Lab sub category
// @Tags 		Lab sub category
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param   	body body models.UpdateSubCategory true "Body"
// @Success 	200 {object} models.CategoryModelResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/lab-sub-category-update/{id}  [post]
func (h *handlerV1) LabSubCategoryUpdate(c *gin.Context) {
	var body models.UpdateSubCategory

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error updating lab sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabSubCategoryUpdate(ctx, &lab.SubCategory{
		Id:   c.Param("id"),
		Name: body.Name,
	})
	if err != nil {
		h.log.Error("Error updating lab sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CategoryModelResp{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 	Delete Lab sub category
// @Description This api can delete Lab sub category
// @Tags 		Lab sub category
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/lab-sub-category-delete/{id}  [delete]
func (h *handlerV1) LabSubCategoryDelete(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.LabService().LabSubCategoryDelete(ctx, &lab.CategoryId{
		Id: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting lab sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	Create Aparat sub category
// @Description This api can registr Aparat sub category create
// @Tags 		Aparat sub category
// @Accept 		json
// @Produce 	json
// @Param body 	body models.SubCategoryModel true "Body"
// @Success 201 {object} models.SubCategoryResModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/aparat-sub-category-create [post]
func (h *handlerV1) AparatSubCategoryCreate(c *gin.Context) {
	var body models.SubCategoryModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating aparat sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatSubCategoryCreate(ctx, &lab.SubCategory{
		Id:         uuid.New().String(),
		Name:       body.Name,
		CategoryId: body.CategoryId,
	})
	if err != nil {
		h.log.Error("Error creating aparat sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.SubCategoryResModel{
		Id:         response.Id,
		Name:       response.Name,
		CategoryId: response.CategoryId,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	})
}

// @Summary 		Get Aparat sub category
// @Description 	This api can get Aparat sub category create
// @Tags 			Aparat sub category
// @Accept 			json
// @Produce         json
// @Param 			filter query models.CategoryGetReqModel false "Filter"
// @Success         200			{object}  models.SubCategoryResModel
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/aparat-sub-category-get [get]
func (h *handlerV1) AparatSubCategoryGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatSubCategoryGet(ctx, &lab.CategoryGetReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})
	if err != nil {
		h.log.Error("Error getting aparat sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.SubCategoryResModel{
		Id:         response.Id,
		Name:       response.Name,
		CategoryId: response.CategoryId,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	})
}

// @Summary 	Find Aparat sub category
// @Description This api can find Aparat sub categories
// @Tags 		Aparat sub category
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.CategoryFindReqModel false "Filter"
// @Success 	200 {object} models.SubCategoriesResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/aparat-sub-category-find [get]
func (h *handlerV1) AparatSubCategoryFind(c *gin.Context) {
	var CategoriesResp models.SubCategoriesResp

	req, err := labParams(c)
	if err != nil {
		h.log.Error("Error finding aparat sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatSubCategoryFind(ctx, &lab.SubCategoryFindReq{
		Limit:      req.Limit,
		Page:       req.Page,
		CategoryId: req.Search,
	})
	if err != nil {
		h.log.Error("Error finding aparat sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, lab := range response.Info {
		CategoriesResp.Category = append(CategoriesResp.Category, &models.SubCategoryResModel{
			Id:         lab.Id,
			Name:       lab.Name,
			CategoryId: lab.CategoryId,
			CreatedAt:  lab.CreatedAt,
			UpdatedAt:  lab.UpdatedAt,
		})
	}
	CategoriesResp.Count = response.Count

	c.JSON(http.StatusCreated, CategoriesResp)
}

// @Summary 	Update Aparat sub category
// @Description This api can update Aparat sub category
// @Tags 		Aparat sub category
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param   	body body models.UpdateSubCategory true "Body"
// @Success 	200 {object} models.CategoryModelResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/aparat-sub-category-update/{id}  [post]
func (h *handlerV1) AparatSubCategoryUpdate(c *gin.Context) {
	var body models.UpdateSubCategory

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error updating aparat sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatSubCategoryUpdate(ctx, &lab.SubCategory{
		Id:   c.Param("id"),
		Name: body.Name,
	})
	if err != nil {
		h.log.Error("Error updating aparat sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CategoryModelResp{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 	Delete Aparat sub category
// @Description This api can delete Aparat sub category
// @Tags 		Aparat sub category
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/aparat-sub-category-delete/{id}  [delete]
func (h *handlerV1) AparatSubCategoryDelete(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.LabService().AparatSubCategoryDelete(ctx, &lab.CategoryId{
		Id: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting aparat sub category", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	Create Aparat analysis
// @Description This api can registr Aparat analysis
// @Tags 		Aparat analysis
// @Accept 		json
// @Produce 	json
// @Param body 	body models.AnalysisReqModel true "Body"
// @Success 201 {object} models.AnalysisRespModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/aparat-analysis-create [post]
func (h *handlerV1) AparatAnalysisCreate(c *gin.Context) {
	var body models.AnalysisReqModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating aparat analysis", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatAnalysisCreate(ctx, &lab.AnalysisReq{
		Id:          uuid.New().String(),
		ClientId:    body.ClientId,
		AparatId:    body.AparatId,
		AnalysisUrl: body.AnalysisUrl,
	})
	if err != nil {
		h.log.Error("Error creating aparat analysis", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.AnalysisRespModel{
		Id:          response.Id,
		ClientId:    response.ClientId,
		AparatId:    response.AparatId,
		AnalysisUrl: response.AnalysisUrl,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Summary 		Get Aparat analysis
// @Description 	This api can get Aparat analysis
// @Tags 			Aparat analysis
// @Accept 			json
// @Produce         json
// @Param 			filter query models.AnalysisGetReqModel false "Filter"
// @Success         200			{object}  models.AnalysisRespModel
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/aparat-analysis-get [get]
func (h *handlerV1) AparatAnalysisGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().AparatAnalysisGet(ctx, &lab.AnalysisGetReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})
	if err != nil {
		h.log.Error("Error getting aparat analysis", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.AnalysisRespModel{
		Id:          response.Id,
		ClientId:    response.ClientId,
		AparatId:    response.AparatId,
		AnalysisUrl: response.AnalysisUrl,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Summary 	Delete Aparat analysis
// @Description This api can delete Aparat analysis
// @Tags 		Aparat analysis
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/aparat-analysis-delete/{id}  [delete]
func (h *handlerV1) AparatAnalysisDelete(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.LabService().AparatAnalysisDelete(ctx, &lab.AparatId{
		Id: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting aparat analysis", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	Create Lab analysis
// @Description This api can registr Lab analysis
// @Tags 		Lab analysis
// @Accept 		json
// @Produce 	json
// @Param body 	body models.AnalysisReqModel true "Body"
// @Success 201 {object} models.AnalysisRespModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/lab-analysis-create [post]
func (h *handlerV1) LabAnalysisCreate(c *gin.Context) {
	var body models.AnalysisReqModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating aparat analysis", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabAnalysisCreate(ctx, &lab.AnalysisReq{
		Id:          uuid.New().String(),
		ClientId:    body.ClientId,
		AparatId:    body.AparatId,
		AnalysisUrl: body.AnalysisUrl,
	})
	if err != nil {
		h.log.Error("Error creating lab analysis", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.AnalysisRespModel{
		Id:          response.Id,
		ClientId:    response.ClientId,
		AparatId:    response.AparatId,
		AnalysisUrl: response.AnalysisUrl,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Summary 		Get Lab analysis
// @Description 	This api can get Lab analysis
// @Tags 			Lab analysis
// @Accept 			json
// @Produce         json
// @Param 			filter query models.AnalysisGetReqModel false "Filter"
// @Success         200			{object}  models.AnalysisRespModel
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/lab-analysis-get [get]
func (h *handlerV1) LabAnalysisGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.LabService().LabAnalysisGet(ctx, &lab.AnalysisGetReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})
	if err != nil {
		h.log.Error("Error getting lab analysis", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.AnalysisRespModel{
		Id:          response.Id,
		ClientId:    response.ClientId,
		AparatId:    response.AparatId,
		AnalysisUrl: response.AnalysisUrl,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Summary 	Delete Lab analysis
// @Description This api can delete Lab analysis
// @Tags 		Lab analysis
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/lab-analysis-delete/{id}  [delete]
func (h *handlerV1) LabAnalysisDelete(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.LabService().LabAnalysisDelete(ctx, &lab.AparatId{
		Id: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting lab analysis", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}
