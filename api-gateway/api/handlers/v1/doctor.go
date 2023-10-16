package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/clinic-crm/api-gateway/api/models"
	"gitlab.com/clinic-crm/api-gateway/genproto/doctor"
	patient "gitlab.com/clinic-crm/api-gateway/genproto/patient"
	"gitlab.com/clinic-crm/api-gateway/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

// @Summary 	Create doctor
// @Description This api can registr doctor
// @Tags 		Doctor
// @Accept 		json
// @Produce 	json
// @Param body 	body models.CreateDoctorModel true "Body"
// @Success 201 {object} models.DoctorResp
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/doctor-create [post]
func (h *handlerV1) DoctorCreate(c *gin.Context) {
	var body models.CreateDoctorModel

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

	response, err := h.serviceManager.DoctorService().DoctorCreate(ctx, &doctor.Doctor{
		Id:          uuid.New().String(),
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Gender:      body.Gender,
		WorkTime:    body.WorkTime,
		Price:       body.Price,
		Cpecialety:  body.Specialty,
		RoomNumber:  body.RoomNumber,
		PhoneNumber: body.PhoneNumber,
	})
	if err != nil {
		h.log.Error("Error creating doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.DoctorResp{
		Id:          response.Id,
		FirstName:   response.FirstName,
		LastName:    response.LastName,
		Gender:      response.Gender,
		WorkTime:    response.WorkTime,
		Price:       response.Price,
		Specialty:   response.Cpecialety,
		RoomNumber:  response.RoomNumber,
		PhoneNumber: response.PhoneNumber,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Summary 		Get doctor
// @Description 	This api can get doctor
// @Tags 			Doctor
// @Accept 			json
// @Produce         json
// @Param 			filter query models.GetDoctorReq false "Filter"
// @Success         200			{object}  models.DoctorResp
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/doctor-get [get]
func (h *handlerV1) DoctorGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().DoctorGet(ctx, &doctor.GetDoctorReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})

	if err != nil {
		h.log.Error("Error getting doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.DoctorResp{
		Id:          response.Id,
		FirstName:   response.FirstName,
		LastName:    response.LastName,
		Gender:      response.Gender,
		WorkTime:    response.WorkTime,
		Price:       response.Price,
		Specialty:   response.Cpecialety,
		RoomNumber:  response.RoomNumber,
		PhoneNumber: response.PhoneNumber,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Summary 	Find doctors
// @Description This api can find doctors
// @Tags 		Doctor
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.DoctorsFindReq false "Filter"
// @Success 	200 {object} models.DoctorsResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/doctor-find [get]
func (h *handlerV1) DoctorsFind(c *gin.Context) {
	var (
		DoctorsResp models.DoctorsResp
	)
	req, err := doctorParams(c)
	if err != nil {
		h.log.Error("Error finding doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().DoctorsFind(ctx, &doctor.DoctorsFindReq{
		Limit:  req.Limit,
		Page:   req.Page,
		Search: req.Search,
	})
	if err != nil {
		h.log.Error("Error finding doctors", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, doctor := range response.Doctors {
		DoctorsResp.Doctors = append(DoctorsResp.Doctors, &models.DoctorResp{
			Id:          doctor.Id,
			FirstName:   doctor.FirstName,
			LastName:    doctor.LastName,
			Gender:      doctor.Gender,
			WorkTime:    doctor.WorkTime,
			Price:       doctor.Price,
			Specialty:   doctor.Cpecialety,
			RoomNumber:  doctor.RoomNumber,
			PhoneNumber: doctor.PhoneNumber,
			CreatedAt:   doctor.CreatedAt,
			UpdatedAt:   doctor.UpdatedAt,
		})
	}
	DoctorsResp.Count = response.Count

	c.JSON(http.StatusCreated, DoctorsResp)
}

func doctorParams(c *gin.Context) (*models.DoctorsFindReq, error) {
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

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	return &models.DoctorsFindReq{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}, nil
}

// @Summary 	Update doctor
// @Description This api can update doctor
// @Tags 		Doctor
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param body 	body models.CreateDoctorModel true "Body"
// @Success 	200 {object} models.DoctorResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/doctor-update/{id}  [post]
func (h *handlerV1) DoctorUpdate(c *gin.Context) {
	var body models.CreateDoctorModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error updating doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().DoctorUpdate(ctx, &doctor.Doctor{
		Id:          c.Param("id"),
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Gender:      body.Gender,
		WorkTime:    body.WorkTime,
		Price:       body.Price,
		Cpecialety:  body.Specialty,
		RoomNumber:  body.RoomNumber,
		PhoneNumber: body.PhoneNumber,
	})
	if err != nil {
		h.log.Error("Error updating doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.DoctorResp{
		Id:          response.Id,
		FirstName:   response.FirstName,
		LastName:    response.LastName,
		Gender:      response.Gender,
		WorkTime:    response.WorkTime,
		Price:       response.Price,
		Specialty:   response.Cpecialety,
		RoomNumber:  response.RoomNumber,
		PhoneNumber: response.PhoneNumber,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Summary 	Delete doctor
// @Description This api can delete doctor
// @Tags 		Doctor
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/doctor-delete/{id}  [delete]
func (h *handlerV1) DoctorDelete(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.DoctorService().DoctorDelete(ctx, &doctor.DoctorId{
		DoctorId: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	Create doctor report
// @Description This api can registr doctor report
// @Tags 		Doctor-report
// @Accept 		json
// @Produce 	json
// @Param body 	body models.DoctorReportsModel true "Body"
// @Success 201 {object} models.DoctorReportsModelRes
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/doctor-report-create [post]
func (h *handlerV1) DoctorReportCreate(c *gin.Context) {
	var body models.DoctorReportsModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating doctor reports", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().DoctorReportCreate(ctx, &doctor.DoctorReport{
		Id:       uuid.New().String(),
		ClientId: body.ClientId,
		DoctorId: body.DoctorId,
		Text:     body.Text,
	})
	if err != nil {
		h.log.Error("Error creating doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.DoctorReportsModelRes{
		Id:        response.Id,
		ClientId:  response.ClientId,
		DoctorId:  response.DoctorId,
		Text:      response.Text,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 		Get doctor report
// @Description 	This api can get doctor reports
// @Tags 			Doctor-report
// @Accept 			json
// @Produce         json
// @Param 			filter query models.GetDoctorReportReq false "Filter"
// @Success         200			{object}  models.DoctorReportsModelRes
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/doctor-report-get [get]
func (h *handlerV1) DoctorReportGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().DoctorReportGet(ctx, &doctor.GetDoctorReport{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})

	if err != nil {
		h.log.Error("Error getting doctor report", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.DoctorReportsModelRes{
		Id:        response.Id,
		ClientId:  response.ClientId,
		DoctorId:  response.DoctorId,
		Text:      response.Text,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Summary 	Find doctors report
// @Description This api can find doctors
// @Tags 		Doctor-report
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.DoctorReportsFindReq false "Filter"
// @Success 	200 {object} models.DoctorReportsResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/doctor-report-find [get]
func (h *handlerV1) DoctorReportsFind(c *gin.Context) {
	var (
		DoctorReportsResp models.DoctorReportsResp
	)
	req, err := doctorParams(c)
	if err != nil {
		h.log.Error("Error finding doctor reports", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().DoctorReportsFind(ctx, &doctor.DoctorReportsFindReq{
		Limit:  req.Limit,
		Page:   req.Page,
		Search: req.Search,
	})
	if err != nil {
		h.log.Error("Error finding doctor reports", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, reports := range response.DoctorReports {
		DoctorReportsResp.Reports = append(DoctorReportsResp.Reports, &models.DoctorReportsModelRes{
			Id:        reports.Id,
			ClientId:  reports.ClientId,
			DoctorId:  reports.DoctorId,
			Text:      reports.Text,
			CreatedAt: reports.CreatedAt,
			UpdatedAt: reports.UpdatedAt,
		})
	}
	DoctorReportsResp.Count = response.Count

	c.JSON(http.StatusCreated, DoctorReportsResp)
}

// @Summary 	Delete doctor report
// @Description This api can delete doctor
// @Tags 		Doctor-report
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/doctor-report-delete/{id}  [delete]
func (h *handlerV1) DoctorReportDelete(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.DoctorService().DoctorReportDelete(ctx, &doctor.ReportId{
		ReportId: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	Create sqlad product info
// @Description This api can registr sqlad product info
// @Tags 		Sqlad
// @Accept 		json
// @Produce 	json
// @Param body 	body models.SqladReqModel true "Body"
// @Success 201 {object} models.SqladRespModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router 		/v1/sqlad-create [post]
func (h *handlerV1) SqladCreate(c *gin.Context) {
	var body models.SqladReqModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating sqlad info", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().SqladCreate(ctx, &doctor.SqladReq{
		Id:             uuid.New().String(),
		Name:           body.Name,
		Count:          body.Count,
		Price:          body.Price,
		LowStock:       body.LowStock,
		ExpirationDate: body.ExpirationDate,
		Provider:       body.Provider,
	})
	if err != nil {
		h.log.Error("Error creating sqlad info", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.SqladRespModel{
		Id:             response.Id,
		Name:           response.Name,
		Count:          response.Count,
		Price:          response.Price,
		LowStock:       response.LowStock,
		ExpirationDate: response.ExpirationDate,
		Provider:       response.Provider,
		CreatedAt:      response.CreatedAt,
		UpdatedAt:      response.UpdatedAt,
	})
}

// @Summary 		Get sqlad product info
// @Description 	This api can get sqlad
// @Tags 			Sqlad
// @Accept 			json
// @Produce         json
// @Param 			filter query models.SqladGetReqModel false "Filter"
// @Success         200			{object}  models.SqladRespModel
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/sqlad-get [get]
func (h *handlerV1) SqladGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().SqladGet(ctx, &doctor.SqladGetReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})

	if err != nil {
		h.log.Error("Error getting sqlad", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SqladRespModel{
		Id:             response.Id,
		Name:           response.Name,
		Count:          response.Count,
		Price:          response.Price,
		LowStock:       response.LowStock,
		ExpirationDate: response.ExpirationDate,
		Provider:       response.Provider,
		CreatedAt:      response.CreatedAt,
		UpdatedAt:      response.UpdatedAt,
	})
}

// @Summary 	Update sqlad product info
// @Description This api can update doctor report
// @Tags 		Sqlad
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param body 	body models.SqladReqModel true "Body"
// @Success 	200 {object} models.SqladRespModel
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/sqlad-update/{id}  [post]
func (h *handlerV1) SqladUpdate(c *gin.Context) {
	var body models.SqladReqModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error updating sqlad info", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().SqladUpdate(ctx, &doctor.SqladReq{
		Id:             c.Param("id"),
		Name:           body.Name,
		Count:          body.Count,
		Price:          body.Price,
		LowStock:       body.LowStock,
		ExpirationDate: body.ExpirationDate,
		Provider:       body.Provider,
	})
	if err != nil {
		h.log.Error("Error updating sqlad info", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SqladRespModel{
		Id:             response.Id,
		Name:           response.Name,
		Count:          response.Count,
		Price:          response.Price,
		LowStock:       response.LowStock,
		ExpirationDate: response.ExpirationDate,
		Provider:       response.Provider,
		CreatedAt:      response.CreatedAt,
		UpdatedAt:      response.UpdatedAt,
	})
}

// @Summary 	Delete sqlad info
// @Description This api can delete product in sqlad
// @Tags 		Sqlad
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/sqlad-delete/{id}  [delete]
func (h *handlerV1) SqladDelete(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.DoctorService().SqladDelete(ctx, &doctor.SqladId{
		Id: c.Param("id"),
	})
	if err != nil {
		h.log.Error("Error deleting product in sqlad", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	Find low stochs
// @Description Retrieves a list of low stock
// @Tags 		Sqlad
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} models.LowStocksRespModel
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/low-stock [get]
func (h *handlerV1) LowStock(c *gin.Context) {
	var (
		LowStocksResp models.LowStocksRespModel
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().LowStock(ctx, &emptypb.Empty{})
	if err != nil {
		h.log.Error("Error finding low stocks", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, product := range response.LowStock {
		LowStocksResp.LowStocks = append(LowStocksResp.LowStocks, &models.SqladRespModel{
			Id:             product.Id,
			Name:           product.Name,
			Count:          product.Count,
			Price:          product.Price,
			LowStock:       product.LowStock,
			ExpirationDate: product.ExpirationDate,
			Provider:       product.Provider,
			CreatedAt:      product.CreatedAt,
			UpdatedAt:      product.UpdatedAt,
		})
	}
	LowStocksResp.Count = response.Count

	c.JSON(http.StatusCreated, LowStocksResp)
}

// @Summary 		Get patients info for doctor page
// @Description 	This api can get doctor page
// @Tags 			Doctor-page
// @Accept 			json
// @Produce         json
// @Param 			filter query models.DocPageFilterReq false "Filter"
// @Success         200			{object}  models.DocPageFilterResModel
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/doctor-page-filter [get]
func (h *handlerV1) DoctorPageFilter(c *gin.Context) {
	var body models.DocPageFilterReq
	var resp models.DocPageFilterResModel

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	respQueue, err := h.serviceManager.PatientService().FindQueue(ctx, &patient.QueueFilter{
		ServiceId:   body.ServiceId,
		ServiceType: body.ServiceType,
		Page:        body.Page,
		Limit:       body.Limit,
	})
	if err != nil {
		h.log.Error("Error gitting patient for doctor page", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, patientResponse := range respQueue.Queues {

		respPatient, err := h.serviceManager.PatientService().PatientGet(ctx, &patient.GetPatientReq{
			Field: "client_id",
			Value: strconv.Itoa(int(patientResponse.ClientId)),
		})
		if err != nil {
			h.log.Error("Error gitting patient for doctor page", logger.Error(err))
			c.JSON(http.StatusInternalServerError, models.ResponseError{
				Message: err.Error(),
			})
			return
		}

		fullName := respPatient.FirstName + respPatient.LastName
		resp.PatientInfo = append(resp.PatientInfo, &models.PatientInfo{
			QueueNumber:   patientResponse.QueueNumber,
			FullName:      fullName,
			PhoneNumber:   respPatient.MainPhoneNumber,
			DateLastVisit: respPatient.UpdatedAt,
		})
	}
}

// @Summary 		Get doctor type
// @Description 	This api can get doctor type
// @Tags 			Doctor
// @Accept 			json
// @Produce         json
// @Success         200			{object}  models.DoctorTypes
// @Failure         400         {object}  models.ResponseError
// @Failure         500         {object}  models.ResponseError
// @Router          /v1/doctor-type-get [get]
func (h *handlerV1) DoctorTypeGet(c *gin.Context) {
	var (
		DoctorTypes models.DoctorTypes
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.DoctorService().DoctorTypeGet(ctx, &emptypb.Empty{})

	if err != nil {
		h.log.Error("Error getting doctor type", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, res := range response.DoctorTypes {
		DoctorTypes.DoctorTypes = append(DoctorTypes.DoctorTypes, &models.DoctorType{
			DoctorType: res.DoctorType,
		})
	}
	DoctorTypes.Count = response.Count

	c.JSON(http.StatusOK, DoctorTypes)
}
