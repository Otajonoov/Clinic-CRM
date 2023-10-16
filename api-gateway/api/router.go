package api

import (
	_ "gitlab.com/clinic-crm/api-gateway/api/docs"
	v1 "gitlab.com/clinic-crm/api-gateway/api/handlers/v1"
	"gitlab.com/clinic-crm/api-gateway/config"
	"gitlab.com/clinic-crm/api-gateway/pkg/logger"
	"gitlab.com/clinic-crm/api-gateway/services"
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct { 	 	
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// @title           MedicalCRM api
// @version         1.0
// @description     This is MedicalCRM server api. Created by Otajonov Quvonchbek
// @termsOfService  1 term MedicalCRM
// @in header
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "App is running...",
		})
	})

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	api := router.Group("/v1")
	// Patient...
	api.POST("/patient-create", handlerV1.PatientCreate)
	api.GET("/patient-get", handlerV1.PatientGet)
	api.GET("/patient-find", handlerV1.PatientsFind)
	api.POST("/patient-update/:id", handlerV1.PatientUpdate)
	api.DELETE("/patient-delete/:id", handlerV1.PatientDelete)

	// Doctor...
	api.POST("/doctor-create", handlerV1.DoctorCreate)
	api.GET("/doctor-get", handlerV1.DoctorGet)
	api.GET("/doctor-find", handlerV1.DoctorsFind)
	api.POST("/doctor-update/:id", handlerV1.DoctorUpdate)
	api.DELETE("/doctor-delete/:id", handlerV1.DoctorDelete)
	api.GET("/doctor-type-get", handlerV1.DoctorTypeGet)

	// Labs...
	api.POST("/lab-create", handlerV1.LabCreate)
	api.GET("/lab-get", handlerV1.LabGet)
	api.GET("/lab-find", handlerV1.LabsFind)
	api.POST("/lab-update/:id", handlerV1.LabUpdate)
	api.DELETE("lab-delete/:id", handlerV1.LabDelete)

	// Aparat...
	api.POST("/aparat-create", handlerV1.AparatCreate)
	api.GET("/aparat-get", handlerV1.AparatGet)
	api.GET("/aparat-find", handlerV1.AparatsFind)
	api.POST("/aparat-update/:id", handlerV1.AparatsUpdate)
	api.DELETE("aparat-delete/:id", handlerV1.AparatsDelete)

	// Lab Category
	api.POST("/lab-category-create", handlerV1.LabCategoryCreate)
	api.GET("/lab-category-get", handlerV1.LabCategoryGet)
	api.GET("/lab-category-find", handlerV1.LabCategoryFind)
	api.POST("/lab-category-update/:id", handlerV1.LabCategoryUpdate)
	api.DELETE("lab-category-delete/:id", handlerV1.LabCategoryDelete)

	// Aparat Category
	api.POST("/aparat-category-create", handlerV1.AparatCategoryCreate)
	api.GET("/aparat-category-get", handlerV1.AparatCategoryGet)
	api.GET("/aparat-category-find", handlerV1.AparatCategoryFind)
	api.POST("/aparat-category-update/:id", handlerV1.AparatCategoryUpdate)
	api.DELETE("aparat-category-delete/:id", handlerV1.AparatCategoryDelete)

	// Lab sub Category
	api.POST("/lab-sub-category-create", handlerV1.LabSubCategoryCreate)
	api.GET("/lab-sub-category-get", handlerV1.LabSubCategoryGet)
	api.GET("/lab-sub-category-find", handlerV1.LabSubCategoryFind)
	api.POST("/lab-sub-category-update/:id", handlerV1.LabSubCategoryUpdate)
	api.DELETE("lab-sub-category-delete/:id", handlerV1.LabSubCategoryDelete)

	// Aparat sub Category
	api.POST("/aparat-sub-category-create", handlerV1.AparatSubCategoryCreate)
	api.GET("/aparat-sub-category-get", handlerV1.AparatSubCategoryGet)
	api.GET("/aparat-sub-category-find", handlerV1.AparatSubCategoryFind)
	api.POST("/aparat-sub-category-update/:id", handlerV1.AparatSubCategoryUpdate)
	api.DELETE("aparat-sub-category-delete/:id", handlerV1.AparatSubCategoryDelete)
	// Doctor report
	api.POST("/doctor-report-create", handlerV1.DoctorReportCreate)
	api.GET("/doctor-report-get", handlerV1.DoctorReportGet)
	api.GET("/doctor-report-find", handlerV1.DoctorReportsFind)
	api.DELETE("doctor-report-delete/:id", handlerV1.DoctorReportDelete)

	// Aparat analysis
	api.POST("/aparat-analysis-create", handlerV1.AparatAnalysisCreate)
	api.GET("/aparat-analysis-get", handlerV1.AparatAnalysisGet)
	api.DELETE("aparat-analysis-delete/:id", handlerV1.AparatAnalysisDelete)

	// Lab analysis
	api.POST("/lab-analysis-create", handlerV1.LabAnalysisCreate)
	api.GET("/lab-analysis-get", handlerV1.LabAnalysisGet)
	api.DELETE("/lab-analysis-delete/:id", handlerV1.LabAnalysisDelete)

	// Sqlad
	api.POST("/sqlad-create", handlerV1.SqladCreate)
	api.GET("/sqlad-get", handlerV1.SqladGet)
	api.GET("/low-stock", handlerV1.LowStock)
	api.POST("/sqlad-update/:id", handlerV1.SqladUpdate)
	api.DELETE("/sqlad-delete/:id", handlerV1.SqladDelete)

	// Queue
	api.POST("/queue-create", handlerV1.PatientQueueCreate)
	api.GET("/queue-get", handlerV1.PatientQueueGet)
	api.GET("/queue-check-get", handlerV1.CheckPatientQueue)
	api.GET("/queue-find", handlerV1.PatientQueuesFind)
	api.POST("/queue-update/:id", handlerV1.PatientQueueUpdate)

	// Cashbox
	api.POST("/cashbox-create", handlerV1.CashboxCreate)
	api.GET("/cashbox-print", handlerV1.CashboxPrint)
	api.GET("/cashbox-find", handlerV1.CashboxFind)
	api.GET("/cashbox-get", handlerV1.CashboxGet)
	api.POST("/cashbox-update/:id", handlerV1.CashboxUpdate)
	api.DELETE("cashbox-delete/:id", handlerV1.CashboxDelete)

	// Payment history
	api.POST("/payment-create", handlerV1.CreatePaymentHistory)
	api.GET("/payment-get", handlerV1.GetPaymentHistory)
	api.GET("/payment-find", handlerV1.FindPaymentHistory)
	api.DELETE("payment-delete/:id", handlerV1.DeletePaymentHistory)

	api.Static("/media", "./media")
	api.POST("/media/photo", handlerV1.UploadMedia)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
