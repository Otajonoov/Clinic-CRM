package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/clinic-crm/api-gateway/api/models"
	"gitlab.com/clinic-crm/api-gateway/genproto/doctor"
	"gitlab.com/clinic-crm/api-gateway/genproto/lab"
	"gitlab.com/clinic-crm/api-gateway/genproto/patient"
	p "gitlab.com/clinic-crm/api-gateway/genproto/patient"
	"gitlab.com/clinic-crm/api-gateway/pkg/logger"
	l "gitlab.com/clinic-crm/api-gateway/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary 	create patient
// @Description This api can patient registr
// @Tags 		Patient
// @Accept 		json
// @Produce 	json
// @Param body 	body models.CreatePatientModel true "Body"
// @Success 	201 {object} models.PatientModel
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/patient-create [post]
func (h *handlerV1) PatientCreate(c *gin.Context) {
	var body models.CreatePatientModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().PatientCreate(ctx, &p.Patient{
		Id:                 uuid.New().String(),
		FirstName:          body.FirstName,
		LastName:           body.LastName,
		Patronymic:         body.Patronymic,
		DateOfBirth:        body.DateOfBirth,
		MainPhoneNumber:    body.MainPhoneNumber,
		OtherPhoneNumber:   body.OtherPhoneNumber,
		AdvertisingChannel: body.AdvertisingChannel,
		Respublic:          body.Respublic,
		Region:             body.Region,
		District:           body.District,
		PassportInfo:       body.PassportInfo,
		Discount:           body.Discount,
		Condition:          body.Condition,
		Gender:             body.Gender,
		DoctorId:           body.DoctorId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("fail to create patient", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.PatientModel{
		Id:                 response.Id,
		ClientId:           response.ClientId,
		FirstName:          response.FirstName,
		LastName:           response.LastName,
		Patronymic:         response.Patronymic,
		DateOfBirth:        response.DateOfBirth,
		MainPhoneNumber:    response.MainPhoneNumber,
		OtherPhoneNumber:   response.OtherPhoneNumber,
		AdvertisingChannel: response.AdvertisingChannel,
		Respublic:          response.Respublic,
		Region:             response.Region,
		District:           response.District,
		PassportInfo:       response.PassportInfo,
		Discount:           response.Discount,
		Condition:          response.Condition,
		Gender:             response.Gender,
		DoctorId:           response.DoctorId,
		CreatedAt:          response.CreatedAt,
		UpdatedAt:          response.UpdatedAt,
	})
}

// @Summary 	get patient
// @Description This api can get patient
// @Tags 		Patient
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.GetPatientReqModel false "Filter"
// @Success 	200 {object} models.PatientModel
// @Failure     400         {object}  models.ResponseError
// @Failure     500         {object}  models.ResponseError
// @Router 		/v1/patient-get [get]
func (h *handlerV1) PatientGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().PatientGet(ctx, &p.GetPatientReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
	})
	if err != nil {
		h.log.Error("Error getting patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.PatientModel{
		Id:                 response.Id,
		ClientId:           response.ClientId,
		FirstName:          response.FirstName,
		LastName:           response.LastName,
		Patronymic:         response.Patronymic,
		DateOfBirth:        response.DateOfBirth,
		MainPhoneNumber:    response.MainPhoneNumber,
		OtherPhoneNumber:   response.OtherPhoneNumber,
		AdvertisingChannel: response.AdvertisingChannel,
		Respublic:          response.Respublic,
		Region:             response.Region,
		District:           response.District,
		PassportInfo:       response.PassportInfo,
		Discount:           response.Discount,
		Condition:          response.Condition,
		Gender:             response.Gender,
		DoctorId:           response.DoctorId,
		CreatedAt:          response.CreatedAt,
		UpdatedAt:          response.UpdatedAt,
	})
}

// @Summary 	update patient
// @Description This api can update patient
// @Tags 		Patient
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Param body 	body models.CreatePatientModel true "UpdatePatientModel"
// @Success 	200 {object} models.PatientModel
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/patient-update/{id} [post]
func (h *handlerV1) PatientUpdate(c *gin.Context) {
	var body models.CreatePatientModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error getting doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().PatientUpdate(ctx, &p.Patient{
		Id:                 c.Param("id"),
		FirstName:          body.FirstName,
		LastName:           body.LastName,
		Patronymic:         body.Patronymic,
		DateOfBirth:        body.DateOfBirth,
		MainPhoneNumber:    body.MainPhoneNumber,
		OtherPhoneNumber:   body.OtherPhoneNumber,
		AdvertisingChannel: body.AdvertisingChannel,
		Respublic:          body.Respublic,
		Region:             body.Region,
		District:           body.District,
		PassportInfo:       body.PassportInfo,
		Discount:           body.Discount,
		Condition:          body.Condition,
		Gender:             body.Gender,
		DoctorId:           body.DoctorId,
	})
	if err != nil {
		h.log.Error("Error getting doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.PatientModel{
		Id:                 response.Id,
		ClientId:           response.ClientId,
		FirstName:          response.FirstName,
		LastName:           response.LastName,
		Patronymic:         response.Patronymic,
		DateOfBirth:        response.DateOfBirth,
		MainPhoneNumber:    response.MainPhoneNumber,
		OtherPhoneNumber:   response.OtherPhoneNumber,
		AdvertisingChannel: response.AdvertisingChannel,
		Respublic:          response.Respublic,
		Region:             response.Region,
		District:           response.District,
		PassportInfo:       response.PassportInfo,
		Discount:           response.Discount,
		Condition:          response.Condition,
		Gender:             response.Gender,
		DoctorId:           response.DoctorId,
		CreatedAt:          response.CreatedAt,
		UpdatedAt:          response.UpdatedAt,
	})
}

// @Summary 	delete patient
// @Description This api can delete patient
// @Tags 		Patient
// @Accept		json
// @Produce 	json
// @Param 		filter query models.PatientId false "Filter"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router		/v1/patient-delete/{id} [delete]
func (h *handlerV1) PatientDelete(c *gin.Context) {
	// id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	// if err != nil {
	// 	h.log.Error("Error parsing patient id", logger.Error(err))
	// 	c.JSON(http.StatusInternalServerError, models.ResponseError{
	// 		Message: err.Error(),
	// 	})
	// 	return
	// }

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	_, err := h.serviceManager.PatientService().PatientDelete(ctx, &p.PatientId{
		//Id: id,
	})
	if err != nil {
		h.log.Error("Error deleting patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully",
	})
}

// @Summary 	find patients
// @Description This api can find patient
// @Tags 		Patient
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.PatientsFindModel false "Filter"
// @Success 	200 {object} models.PatientsResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
// @Router 		/v1/patient-find [get]
func (h *handlerV1) PatientsFind(c *gin.Context) {
	var (
		PatientsResp models.PatientsResp
	)
	req, err := patientParams(c)
	if err != nil {
		h.log.Error("Error finding patients", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().PatientsFind(ctx, &p.PatientsFindReq{
		Limit:  req.Limit,
		Page:   req.Page,
		Search: req.Search,
	})
	if err != nil {
		h.log.Error("Error finding patients", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, patient := range response.Patients {
		PatientsResp.Patients = append(PatientsResp.Patients, &models.PatientModel{
			Id:                 patient.Id,
			ClientId:           patient.ClientId,
			FirstName:          patient.FirstName,
			LastName:           patient.LastName,
			Patronymic:         patient.Patronymic,
			DateOfBirth:        patient.DateOfBirth,
			MainPhoneNumber:    patient.MainPhoneNumber,
			OtherPhoneNumber:   patient.OtherPhoneNumber,
			AdvertisingChannel: patient.AdvertisingChannel,
			Respublic:          patient.Respublic,
			Region:             patient.Region,
			District:           patient.District,
			PassportInfo:       patient.PassportInfo,
			Discount:           patient.Discount,
			Condition:          patient.Condition,
			Gender:             patient.Gender,
			CreatedAt:          patient.CreatedAt,
			UpdatedAt:          patient.UpdatedAt,
		})
	}
	PatientsResp.Count = response.Count

	c.JSON(http.StatusCreated, PatientsResp)
}

func patientParams(c *gin.Context) (*models.PatientsFindModel, error) {
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

	return &models.PatientsFindModel{
		Limit:  int64(limit),
		Page:   int64(page),
		Search: c.Query("search"),
	}, nil
}

// // @Summary 	filter patients
// // @Description This api can filter patient
// // @Tags 		Filter
// // @Accept 		json
// // @Produce 	json
// // @Param 		filter query models.GetPatientReqModel false "Filter"
// // @Success 	200 {object} models.PatientFilterModel
// // @Failure 	400 {object} models.ResponseError
// // @Failure 	500 {object} models.ResponseError
// // @Router 		/v1/patient-filter [get]
// func (h *handlerV1) FindPatientWithFilter(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
// 	defer cancel()
// 	resPatient, err := h.serviceManager.PatientService().PatientGet(ctx, &p.GetPatientReq{
// 		Field: c.Query("field"),
// 		Value: c.Query("value"),
// 	})
// 	if err != nil {
// 		h.log.Error("Error getting patient", logger.Error(err))
// 		c.JSON(http.StatusInternalServerError, models.ResponseError{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	resAnalysis, err := h.serviceManager.PatientService().PatientAnalysisGet(ctx, &p.PatientPhoneNumber{
// 		PhoneNumber: resPatient.MainPhoneNumber,
// 	})
// 	if err != nil {
// 		h.log.Error("Error getting analysiss", logger.Error(err))
// 		c.JSON(http.StatusInternalServerError, models.ResponseError{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	resDoctorRes, err := h.serviceManager.PatientService().DoctorReportGet(ctx, &p.PatientId{
// 		//Id: resPatient.ClientId,
// 	})
// 	if err != nil {
// 		h.log.Error("Error getting doctor reports", logger.Error(err))
// 		c.JSON(http.StatusInternalServerError, models.ResponseError{
// 			Message: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, &models.PatientFilterModel{
// 		PatientModel: models.PatientModel{
// 			Id:                 resPatient.Id,
// 			ClientId:           resPatient.ClientId,
// 			FirstName:          resPatient.FirstName,
// 			LastName:           resPatient.LastName,
// 			Patronymic:         resPatient.Patronymic,
// 			DateOfBirth:        resPatient.DateOfBirth,
// 			MainPhoneNumber:    resPatient.MainPhoneNumber,
// 			OtherPhoneNumber:   resPatient.OtherPhoneNumber,
// 			AdvertisingChannel: resPatient.AdvertisingChannel,
// 			Respublic:          resPatient.Respublic,
// 			Region:             resPatient.Region,
// 			District:           resPatient.District,
// 			PassportInfo:       resPatient.PassportInfo,
// 			Discount:           resPatient.Discount,
// 			Condition:          resPatient.Condition,
// 			Gender:             resPatient.Gender,
// 			CreatedAt:          resPatient.CreatedAt,
// 			UpdatedAt:          resPatient.UpdatedAt,
// 		},
// 		AnalysisInfo: models.AnalysisInfo{
// 			Id:                resAnalysis.Id,
// 			ClientPhoneNumber: resPatient.MainPhoneNumber,
// 			AnalysName:        resAnalysis.AnalysName,
// 			AnalysUrl:         resAnalysis.AnalysUrl,
// 			CreatedAt:         resAnalysis.CreatedAt,
// 			UpdatedAt:         resAnalysis.UpdatedAt,
// 		},
// 	})
// }

// @Router 		/v1/queue-create [post]
// @Summary 	create patient queue
// @Description create patient queue
// @Tags 		Queue
// @Accept 		json
// @Produce 	json
// @Param body 	body models.CreatePatientQueueReq true "Body"
// @Success 	201 {object} models.PatientQueueResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
func (h *handlerV1) PatientQueueCreate(c *gin.Context) {
	var body models.CreatePatientQueueReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().CreatePatientQueue(ctx, &p.CreatePatientQueueReq{
		Id:          uuid.New().String(),
		ClientId:    body.ClientId,
		ServiceId:   body.ServiceId,
		ServiceType: body.ServiceType,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("fail to create patient", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.PatientQueueResp{
		Id:          response.Id,
		ClientId:    response.ClientId,
		QueueNumber: response.QueueNumber,
		ServiceId:   response.ServiceId,
		ServiceType: response.ServiceType,
		TurnPassed:  response.TurnPassed,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Router 		/v1/queue-get [get]
// @Summary 	get patient queue
// @Description This api can get patient queue
// @Tags 		Queue
// @Accept 		json
// @Produce 	json
// @Param 		filter 		query models.PatientQueueId false "Filter"
// @Success 	200 		{object} models.PatientQueueResp
// @Failure     400         {object}  models.ResponseError
// @Failure     500         {object}  models.ResponseError
func (h *handlerV1) PatientQueueGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().GetPatientQueue(ctx, &p.PaymentHistoryId{
		Id: c.Query("id"),
	})
	if err != nil {
		h.log.Error("Error getting patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.PatientQueueResp{
		Id:          response.Id,
		ClientId:    response.ClientId,
		QueueNumber: response.QueueNumber,
		ServiceId:   response.ServiceId,
		ServiceType: response.ServiceType,
		TurnPassed:  response.TurnPassed,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Router 		/v1/queue-check-get [get]
// @Summary 	check patient queue
// @Description This api can check patient queue
// @Tags 		Queue
// @Accept 		json
// @Produce 	json
// @Param 		filter 		query models.CheckQueueReq false "Filter"
// @Success 	200 		{object} models.QueueNumber
// @Failure     400         {object}  models.ResponseError
// @Failure     500         {object}  models.ResponseError
func (h *handlerV1) CheckPatientQueue(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().CheckServiceQueue(ctx, &p.CheckQueueReq{
		ServiceId:   c.Query("service_id"),
		ServiceType: c.Query("service_type"),
	})
	if err != nil {
		h.log.Error("Error getting patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.QueueNumber{
		QueueNumber: response.QueueNumber,
	})
}

// @Router 		/v1/queue-update/{id} [post]
// @Summary 	update patient queue
// @Description This api can update patient queue
// @Tags 		Queue
// @Accept 		json
// @Produce 	json
// @Param body 	body models.UpdateQueueReq true "UpdatePatientModel"
// @Success 	200 {object} models.PatientQueueResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
func (h *handlerV1) PatientQueueUpdate(c *gin.Context) {
	var body models.UpdateQueueReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error getting doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().UpdateQueue(ctx, &p.UpdateQueueReq{
		ClientId:    body.ClientId,
		ServiceId:   body.ServiceId,
		ServiceType: body.ServiceType,
	})
	if err != nil {
		h.log.Error("Error getting doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.PatientQueueResp{
		Id:          response.Id,
		ClientId:    response.ClientId,
		QueueNumber: response.QueueNumber,
		ServiceId:   response.ServiceId,
		ServiceType: response.ServiceType,
		TurnPassed:  response.TurnPassed,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Router 		/v1/queue-find [get]
// @Summary 	Find patient queues
// @Description This api can find patient queues
// @Tags 		Queue
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.QueueFilter false "Filter"
// @Success 	200 {object} models.QueuesResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
func (h *handlerV1) PatientQueuesFind(c *gin.Context) {
	var (
		QueuesResp models.QueuesResp
	)
	req, err := patientQueueParams(c)
	if err != nil {
		h.log.Error("Error finding doctor reports", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.PatientService().FindQueue(ctx, &patient.QueueFilter{
		Limit:       req.Limit,
		Page:        req.Page,
		ClientId:    req.ClientId,
		ServiceId:   req.ServiceId,
		ServiceType: req.ServiceType,
	})
	if err != nil {
		h.log.Error("Error finding doctor reports", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, queue := range response.Queues {
		QueuesResp.Queues = append(QueuesResp.Queues, &models.PatientQueueResp{
			Id:          queue.Id,
			ClientId:    queue.ClientId,
			QueueNumber: queue.QueueNumber,
			ServiceId:   queue.ServiceId,
			ServiceType: queue.ServiceType,
			TurnPassed:  queue.TurnPassed,
			CreatedAt:   queue.CreatedAt,
			UpdatedAt:   queue.UpdatedAt,
		})
	}
	QueuesResp.Count = int(response.Count)

	c.JSON(http.StatusCreated, QueuesResp)
}

func patientQueueParams(c *gin.Context) (*models.QueueFilter, error) {
	var (
		limit     int = 10
		page      int = 1
		client_id int = 0
		err       error
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

	if c.Query("client_id") != "" {
		client_id, err = strconv.Atoi(c.Query("client_id"))
		if err != nil {
			return nil, err
		}
	}
	return &models.QueueFilter{
		ServiceId:   c.Query("service_id"),
		ServiceType: c.Query("service_type"),
		Limit:       int64(limit),
		Page:        int64(page),
		ClientId:    int64(client_id),
	}, nil
}

// @Router 		/v1/cashbox-create [post]
// @Summary 	create cashbox
// @Description create cashbox
// @Tags 		Cashbox
// @Accept 		json
// @Produce 	json
// @Param body 	body models.CreateCashboxReq true "Body"
// @Success 	201 {object} models.CashboxResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
func (h *handlerV1) CashboxCreate(c *gin.Context) {
	var body models.CreateCashboxReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().CreateCashbox(ctx, &p.CreateCashboxReq{
		Id:          uuid.New().String(),
		ClientId:    body.ClientId,
		IsPayed:     body.IsPayed,
		CashCount:   0,
		PaymentType: body.PaymentType,
		DoctorsIds:  body.DoctorsIds,
		LabsIds:     body.LabsIds,
		AparatsIds:  body.AparatsIds,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("fail to create patient", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.CashboxResp{
		Id:         response.Id,
		ClientId:   int(response.ClientId),
		Summa:      int(response.Summa),
		IsPayed:    response.IsPayed,
		CashCount:  int(response.CashCount),
		DoctorsIds: response.DoctorsIds,
		LabsIds:    response.LabsIds,
		AparatsIds: response.AparatsIds,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	})
}

// @Router 		/v1/cashbo-find [get]
// @Summary 	Find cashbox
// @Description This api can find cashbox
// @Tags 		Cashbox
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.FindCashboxReq false "Filter"
// @Success 	200 {object} models.FindCashboxResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
func (h *handlerV1) CashboxFind(c *gin.Context) {
	var (
		CashboxesResp models.FindCashboxResp
	)
	req, err := cashboxParams(c)
	if err != nil {
		h.log.Error("Error finding doctor reports", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.PatientService().FindCashbox(ctx, &patient.FindCashboxReq{
		Limit:    int64(req.Limit),
		Page:     int64(req.Page),
		ClientId: int64(req.ClientId),
		Search:   req.Search,
		FromDate: req.FromDate,
		ToDate:   req.ToDate,
	})
	if err != nil {
		h.log.Error("Error finding doctor reports", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, queue := range response.Cashboxes {
		CashboxesResp.Cashboxes = append(CashboxesResp.Cashboxes, &models.CashboxResp{
			Id:         queue.Id,
			ClientId:   int(queue.ClientId),
			Summa:      int(queue.Summa),
			IsPayed:    queue.IsPayed,
			CashCount:  int(queue.CashCount),
			DoctorsIds: queue.DoctorsIds,
			LabsIds:    queue.LabsIds,
			AparatsIds: queue.AparatsIds,
			CreatedAt:  queue.CreatedAt,
			UpdatedAt:  queue.UpdatedAt,
		})
	}
	CashboxesResp.Count = int(response.Count)

	c.JSON(http.StatusCreated, CashboxesResp)
}

func cashboxParams(c *gin.Context) (*models.FindCashboxReq, error) {
	var (
		limit     int = 10
		page      int = 1
		client_id int = 0
		err       error
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

	if c.Query("client_id") != "" {
		client_id, err = strconv.Atoi(c.Query("client_id"))
		if err != nil {
			return nil, err
		}
	}
	return &models.FindCashboxReq{
		FromDate: c.Query("from_date"),
		ToDate:   c.Query("to_date"),
		Search:   c.Query("search"),
		Limit:    limit,
		Page:     page,
		ClientId: client_id,
	}, nil
}

// @Router 		/v1/cashbox-get [get]
// @Summary 	get patient cashbox
// @Description This api can get patient queue
// @Tags 		Cashbox
// @Accept 		json
// @Produce 	json
// @Param 		filter 		query models.PatientQueueId false "Filter"
// @Success 	200 		{object} models.CashboxResp
// @Failure     400         {object}  models.ResponseError
// @Failure     500         {object}  models.ResponseError
func (h *handlerV1) CashboxGet(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().GetCashbox(ctx, &p.GetCashboxReq{
		CashboxId: c.Query("id"),
	})
	if err != nil {
		h.log.Error("Error getting patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CashboxResp{
		Id:         response.Id,
		ClientId:   int(response.ClientId),
		Summa:      int(response.Summa),
		IsPayed:    response.IsPayed,
		CashCount:  int(response.CashCount),
		DoctorsIds: response.DoctorsIds,
		LabsIds:    response.LabsIds,
		AparatsIds: response.AparatsIds,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	})
}

// @Router 		/v1/cashbox-print [get]
// @Summary 	get patient cashbox print
// @Description This api can get patient queue
// @Tags 		Cashbox
// @Accept 		json
// @Produce 	json
// @Param 		filter 		query models.PatientQueueId false "Filter"
// @Success 	200 		{object} models.CashboxesPrinterResp
// @Failure     400         {object}  models.ResponseError
// @Failure     500         {object}  models.ResponseError
func (h *handlerV1) CashboxPrint(c *gin.Context) {
	var (
		result models.CashboxesPrinterResp
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().GetCashbox(ctx, &p.GetCashboxReq{
		CashboxId: c.Query("id"),
	})

	userResp, err := h.serviceManager.PatientService().PatientGet(ctx, &p.GetPatientReq{
		Field: "client_id",
		Value: strconv.Itoa(int(response.ClientId)),
	})

	if err != nil {
		h.log.Error("Error getting patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, aparatId := range response.AparatsIds {
		aparat, err := h.serviceManager.LabService().AparatGet(ctx, &lab.AparatGetReq{Field: "id", Value: aparatId})
		if err != nil {
			h.log.Error("Error getting patient", logger.Error(err))
			c.JSON(http.StatusInternalServerError, models.ResponseError{
				Message: err.Error(),
			})
			return
		}
		doctorInfo, err := h.serviceManager.DoctorService().DoctorGet(ctx, &doctor.GetDoctorReq{Field: "id", Value: aparat.DoctorId})
		if err != nil {
			h.log.Error("Error getting patient", logger.Error(err))
			c.JSON(http.StatusInternalServerError, models.ResponseError{
				Message: err.Error(),
			})
			return
		}

		result.Cashboxes = append(result.Cashboxes, &models.CashboxPrinterResp{
			ImageUrl:    "https://www.impulse-clinic.com/wp-content/uploads/2019/08/logo_new-2.png",
			CashCount:   int(response.CashCount),
			FirstName:   userResp.FirstName,
			LastName:    userResp.LastName,
			ServiceType: aparat.Type,
			DoctorName:  doctorInfo.FirstName + " " + doctorInfo.LastName,
			RoomNumber:  aparat.RoomNumber,
			Summa:       int(response.Summa),
			CreatedAt:   response.CreatedAt,
		})
		result.Count += 1
	}

	for _, doctorId := range response.DoctorsIds {
		doctor, err := h.serviceManager.DoctorService().DoctorGet(ctx, &doctor.GetDoctorReq{Field: "id", Value: doctorId})
		if err != nil {
			h.log.Error("Error getting patient", logger.Error(err))
			c.JSON(http.StatusInternalServerError, models.ResponseError{
				Message: err.Error(),
			})
			return
		}

		result.Cashboxes = append(result.Cashboxes, &models.CashboxPrinterResp{
			ImageUrl:    "https://www.impulse-clinic.com/wp-content/uploads/2019/08/logo_new-2.png",
			CashCount:   int(response.CashCount),
			FirstName:   userResp.FirstName,
			LastName:    userResp.LastName,
			ServiceType: doctor.Cpecialety,
			DoctorName:  doctor.FirstName + " " + doctor.LastName,
			RoomNumber:  doctor.RoomNumber,
			Summa:       int(response.Summa),
			CreatedAt:   response.CreatedAt,
		})
		result.Count += 1
	}

	for _, labId := range response.LabsIds {
		lab, err := h.serviceManager.LabService().LabGet(ctx, &lab.LabGetReq{Field: "id", Value: labId})
		if err != nil {
			h.log.Error("Error getting patient", logger.Error(err))
			c.JSON(http.StatusInternalServerError, models.ResponseError{
				Message: err.Error(),
			})
			return
		}
		doctorInfo, err := h.serviceManager.DoctorService().DoctorGet(ctx, &doctor.GetDoctorReq{Field: "id", Value: lab.DoctorId})
		if err != nil {
			h.log.Error("Error getting patient", logger.Error(err))
			c.JSON(http.StatusInternalServerError, models.ResponseError{
				Message: err.Error(),
			})
			return
		}

		result.Cashboxes = append(result.Cashboxes, &models.CashboxPrinterResp{
			ImageUrl:    "https://www.impulse-clinic.com/wp-content/uploads/2019/08/logo_new-2.png",
			CashCount:   int(response.CashCount),
			FirstName:   userResp.FirstName,
			LastName:    userResp.LastName,
			ServiceType: lab.Type,
			DoctorName:  doctorInfo.FirstName + " " + doctorInfo.LastName,
			RoomNumber:  lab.RoomNumber,
			Summa:       int(response.Summa),
			CreatedAt:   response.CreatedAt,
		})
		result.Count += 1
	}

	c.JSON(http.StatusCreated, result)
}

// @Router 		/v1/cashbox-update/{id} [post]
// @Summary 	update cashbox
// @Description This api can update cashbox
// @Tags 		Cashbox
// @Accept 		json
// @Produce 	json
// @Param body 	body models.UpdateCashboxReq true "UpdatePatientModel"
// @Success 	200 {object} models.CashboxResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
func (h *handlerV1) CashboxUpdate(c *gin.Context) {
	var body models.UpdateCashboxReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error getting doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().UpdateCashbox(ctx, &p.UpdateCashboxReq{
		Id:          body.Id,
		IsPayed:     body.IsPayed,
		PaymentType: body.PaymentType,
	})
	if err != nil {
		h.log.Error("Error getting doctor", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.CashboxResp{
		Id:         response.Id,
		ClientId:   int(response.ClientId),
		Summa:      int(response.Summa),
		IsPayed:    response.IsPayed,
		CashCount:  int(response.CashCount),
		DoctorsIds: response.DoctorsIds,
		LabsIds:    response.LabsIds,
		AparatsIds: response.AparatsIds,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	})
}

// @Router 		/v1/cashbox-delete/{id}  [delete]
// @Summary 	Delete cashbox
// @Description This api can delete cashbox
// @Tags 		Cashbox
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
func (h *handlerV1) CashboxDelete(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.PatientService().DeleteCashbox(ctx, &patient.GetCashboxReq{
		CashboxId: c.Param("id"),
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

// @Router 		/v1/payment-create [post]
// @Summary 	create payment history
// @Description create payment history
// @Tags 		Payment history
// @Accept 		json
// @Produce 	json
// @Param body 	body models.CreatePaymentHistoryReq true "Body"
// @Success 	201 {object} models.PaymentHistoryResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
func (h *handlerV1) CreatePaymentHistory(c *gin.Context) {
	var body models.CreatePaymentHistoryReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error creating patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().CreatePaymentHistory(ctx, &p.CreatePaymentHistoryReq{
		Id:          uuid.New().String(),
		ClientId:    body.ClientId,
		Summa:       body.Summa,
		CashboxId:   body.CashboxId,
		PaymentType: body.PaymentType,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("fail to create patient", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.PaymentHistoryResp{
		Id:          response.Id,
		ClientId:    response.ClientId,
		Summa:       response.Summa,
		PaymentType: response.PaymentType,
		CashboxId:   response.CashboxId,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Router 		/v1/payment-get [get]
// @Summary 	get payment history
// @Description This api can get payment history
// @Tags 		Payment history
// @Accept 		json
// @Produce 	json
// @Param 		filter 		query models.PaymentHistoryId false "Filter"
// @Success 	200 		{object} models.PaymentHistoryResp
// @Failure     400         {object}  models.ResponseError
// @Failure     500         {object}  models.ResponseError
func (h *handlerV1) GetPaymentHistory(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(time.Duration(h.cfg.CtxTimeout)))
	defer cancel()

	response, err := h.serviceManager.PatientService().GetPaymentHistory(ctx, &p.PaymentHistoryId{
		Id: c.Query("id"),
	})
	if err != nil {
		h.log.Error("Error getting patient", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.PaymentHistoryResp{
		Id:          response.Id,
		ClientId:    response.ClientId,
		Summa:       response.Summa,
		PaymentType: response.PaymentType,
		CashboxId:   response.CashboxId,
		CreatedAt:   response.CreatedAt,
		UpdatedAt:   response.UpdatedAt,
	})
}

// @Router 		/v1/payment-find [get]
// @Summary 	Find cashbox
// @Description This api can find cashbox
// @Tags 		Payment history
// @Accept 		json
// @Produce 	json
// @Param 		filter query models.PaymentHistoryFilter false "Filter"
// @Success 	200 {object} models.PaymentHistoriesResp
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
func (h *handlerV1) FindPaymentHistory(c *gin.Context) {
	var (
		paymentHistoryResp models.PaymentHistoriesResp
	)
	req, err := paymentHistoryParams(c)
	if err != nil {
		h.log.Error("Error finding doctor reports", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := h.serviceManager.PatientService().FindPaymentHistory(ctx, &patient.PaymentHistoryFilter{
		Limit:    int64(req.Limit),
		Page:     int64(req.Page),
		ClientId: int64(req.ClientId),
		FromDate: req.FromDate,
		ToDate:   req.ToDate,
	})
	if err != nil {
		h.log.Error("Error finding doctor reports", logger.Error(err))
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Message: err.Error(),
		})
		return
	}

	for _, history := range response.PaymentHistory {
		paymentHistoryResp.PaymentHistories = append(paymentHistoryResp.PaymentHistories, &models.PaymentHistoryResp{
			Id:          history.Id,
			ClientId:    history.ClientId,
			Summa:       history.Summa,
			PaymentType: history.PaymentType,
			CashboxId:   history.CashboxId,
			CreatedAt:   history.CreatedAt,
			UpdatedAt:   history.UpdatedAt,
		})
	}
	paymentHistoryResp.Count = int(response.Count)

	c.JSON(http.StatusCreated, paymentHistoryResp)
}

func paymentHistoryParams(c *gin.Context) (*models.PaymentHistoryFilter, error) {
	var (
		limit     int = 10
		page      int = 1
		client_id int = 0
		err       error
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

	if c.Query("client_id") != "" {
		client_id, err = strconv.Atoi(c.Query("client_id"))
		if err != nil {
			return nil, err
		}
	}
	return &models.PaymentHistoryFilter{
		FromDate: c.Query("from_date"),
		ToDate:   c.Query("to_date"),
		Limit:    limit,
		Page:     page,
		ClientId: client_id,
	}, nil
}

// @Router 		/v1/payment-delete/{id}  [delete]
// @Summary 	Delete payment history
// @Description This api can delete payment history
// @Tags 		Payment history
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "ID"
// @Success 	200 {object} models.ResponseOK
// @Failure 	400 {object} models.ResponseError
// @Failure 	500 {object} models.ResponseError
func (h *handlerV1) DeletePaymentHistory(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	_, err := h.serviceManager.PatientService().DeletePaymentHistory(ctx, &patient.PaymentHistoryId{
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
