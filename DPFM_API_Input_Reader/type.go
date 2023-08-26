package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string   `json:"connection_key"`
	Result           bool     `json:"result"`
	RedisKey         string   `json:"redis_key"`
	Filepath         string   `json:"filepath"`
	APIStatusCode    int      `json:"api_status_code"`
	RuntimeSessionID string   `json:"runtime_session_id"`
	BusinessPartner  int      `json:"business_partner"`
	ServiceLabel     string   `json:"service_label"`
	APIType          string   `json:"api_type"`
	Header           Header   `json:"MaintenanceOrder"`
	APISchema        string   `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	Deleted          bool     `json:"deleted"`
}

type Header struct {
	MaintenanceOrder               string  `json:"MaintenanceOrder"`
	MaintOrderRoutingNumber        *string `json:"MaintOrderRoutingNumber"`
	MaintenanceOrderType           *string `json:"MaintenanceOrderType"`
	MaintenanceOrderDesc           *string `json:"MaintenanceOrderDesc"`
	MaintOrdBasicStart             *string `json:"MaintOrdBasicStart"`
	MaintOrdBasicEnd               *string `json:"MaintOrdBasicEnd"`
	MaintOrdBasicStartDate         *string `json:"MaintOrdBasicStartDate"`
	MaintOrdBasicStartTime         *string `json:"MaintOrdBasicStartTime"`
	MaintOrdBasicEndDate           *string `json:"MaintOrdBasicEndDate"`
	MaintOrdBasicEndTime           *string `json:"MaintOrdBasicEndTime"`
	MaintOrdSchedldBscStrt         *string `json:"MaintOrdSchedldBscStrt"`
	MaintOrdSchedldBscEnd          *string `json:"MaintOrdSchedldBscEnd"`
	ScheduledBasicStartDate        *string `json:"ScheduledBasicStartDate"`
	ScheduledBasicStartTime        *string `json:"ScheduledBasicStartTime"`
	ScheduledBasicEndDate          *string `json:"ScheduledBasicEndDate"`
	ScheduledBasicEndTime          *string `json:"ScheduledBasicEndTime"`
	MaintenanceNotification        *string `json:"MaintenanceNotification"`
	OrdIsNotSchedldAutomatically   *string `json:"OrdIsNotSchedldAutomatically"`
	MainWorkCenterInternalID       *string `json:"MainWorkCenterInternalID"`
	MainWorkCenterTypeCode         *string `json:"MainWorkCenterTypeCode"`
	MainWorkCenter                 *string `json:"MainWorkCenter"`
	MainWorkCenterPlant            *string `json:"MainWorkCenterPlant"`
	WorkCenterInternalID           *string `json:"WorkCenterInternalID"`
	WorkCenterTypeCode             *string `json:"WorkCenterTypeCode"`
	WorkCenter                     *string `json:"WorkCenter"`
	MaintenancePlanningPlant       *string `json:"MaintenancePlanningPlant"`
	MaintenancePlant               *string `json:"MaintenancePlant"`
	Assembly                       *string `json:"Assembly"`
	MaintOrdProcessPhaseCode       *string `json:"MaintOrdProcessPhaseCode"`
	MaintOrdProcessSubPhaseCode    *string `json:"MaintOrdProcessSubPhaseCode"`
	BusinessArea                   *string `json:"BusinessArea"`
	CompanyCode                    *string `json:"CompanyCode"`
	CostCenter                     *string `json:"CostCenter"`
	ReferenceElement               *string `json:"ReferenceElement"`
	FunctionalArea                 *string `json:"FunctionalArea"`
	AdditionalDeviceData           *string `json:"AdditionalDeviceData"`
	Equipment                      *string `json:"Equipment"`
	EquipmentName                  *string `json:"EquipmentName"`
	FunctionalLocation             *string `json:"FunctionalLocation"`
	MaintenanceOrderPlanningCode   *string `json:"MaintenanceOrderPlanningCode"`
	MaintenancePlannerGroup        *string `json:"MaintenancePlannerGroup"`
	MaintenanceActivityType        *string `json:"MaintenanceActivityType"`
	MaintPriority                  *string `json:"MaintPriority"`
	MaintPriorityType              *string `json:"MaintPriorityType"`
	OrderProcessingGroup           *string `json:"OrderProcessingGroup"`
	ProfitCenter                   *string `json:"ProfitCenter"`
	ResponsibleCostCenter          *string `json:"ResponsibleCostCenter"`
	MaintenanceRevision            *string `json:"MaintenanceRevision"`
	SerialNumber                   *string `json:"SerialNumber"`
	SuperiorProjectNetwork         *string `json:"SuperiorProjectNetwork"`
	OperationSystemCondition       *string `json:"OperationSystemCondition"`
	Project                        *string `json:"Project"`
	ControllingObjectClass         *string `json:"ControllingObjectClass"`
	MaintenanceOrderInternalID     *string `json:"MaintenanceOrderInternalID"`
	MaintenanceObjectList          *string `json:"MaintenanceObjectList"`
	MaintObjectLocAcctAssgmtNmbr   *string `json:"MaintObjectLocAcctAssgmtNmbr"`
	AssetLocation                  *string `json:"AssetLocation"`
	AssetRoom                      *string `json:"AssetRoom"`
	PlantSection                   *string `json:"PlantSection"`
	BasicSchedulingType            *string `json:"BasicSchedulingType"`
	LatestAcceptableCompletionDate *string `json:"LatestAcceptableCompletionDate"`
	MaintOrdPersonResponsible      *string `json:"MaintOrdPersonResponsible"`
	LastChange                     *string `json:"LastChange"`
	ControllingSettlementProfile   *string `json:"ControllingSettlementProfile"`
	SystemStatusText               *string `json:"SystemStatusText"`
	UserStatusText                 *string `json:"UserStatusText"`
	TechnicalObject                *string `json:"TechnicalObject"`
	TechnicalObjectLabel           *string `json:"TechnicalObjectLabel"`
	TechObjIsEquipOrFuncnlLoc      *string `json:"TechObjIsEquipOrFuncnlLoc"`
}
type ObjectListItem struct {
	MaintenanceOrder            string  `json:"MaintenanceOrder"`
	MaintenanceOrderObjectList  int     `json:"MaintenanceOrderObjectList"`
	MaintenanceObjectListItem   int     `json:"MaintenanceObjectListItem"`
	Equipment                   *string `json:"Equipment"`
	MaintenanceNotification     *string `json:"MaintenanceNotification"`
	Assembly                    *string `json:"Assembly"`
	Product                     *string `json:"Product"`
	SerialNumber                *string `json:"SerialNumber"`
	UniqueItemIdentifier        *string `json:"UniqueItemIdentifier"`
	FunctionalLocation          *string `json:"FunctionalLocation"`
	MaintObjectListItemSequence *string `json:"MaintObjectListItemSequence"`
}
type OperationComponent struct {
	MaenanceOrder                  string  `json:"MaenanceOrder"`
	MaenanceOrderComponent         string  `json:"MaenanceOrderComponent"`
	Reservation                    *string `json:"Reservation"`
	ReservationItem                *string `json:"ReservationItem"`
	ReservationType                *string `json:"ReservationType"`
	MaOrderRoutingNumber           *string `json:"MaOrderRoutingNumber"`
	MaOrderOperationCounter        *string `json:"MaOrderOperationCounter"`
	Product                        *string `json:"Product"`
	MaOrdOperationComponentText    *string `json:"MaOrdOperationComponentText"`
	MaOrdOpCompRequiredQuantity    *string `json:"MaOrdOpCompRequiredQuantity"`
	BaseUnit                       *string `json:"BaseUnit"`
	QuantityInUnitOfEntry          *string `json:"QuantityInUnitOfEntry"`
	UnitOfEntry                    *string `json:"UnitOfEntry"`
	RequirementDate                *string `json:"RequirementDate"`
	RequirementTime                *string `json:"RequirementTime"`
	Supplier                       *string `json:"Supplier"`
	Plant                          *string `json:"Plant"`
	StorageLocation                *string `json:"StorageLocation"`
	MaOrdOpCompItemCategory        *string `json:"MaOrdOpCompItemCategory"`
	GoodsMovementType              *string `json:"GoodsMovementType"`
	ReservationIsFinallyIssued     *int    `json:"ReservationIsFinallyIssued"`
	ProductGroup                   *string `json:"ProductGroup"`
	ProductTypeCode                *string `json:"ProductTypeCode"`
	ServicePerformer               *string `json:"ServicePerformer"`
	PerformancePeriodStartDateTime *string `json:"PerformancePeriodStartDateTime"`
	PerformancePeriodStartDate     *string `json:"PerformancePeriodStartDate"`
	PerformancePeriodEndDate       *string `json:"PerformancePeriodEndDate"`
	PerformancePeriodEndDateTime   *string `json:"PerformancePeriodEndDateTime"`
	PerformancePeriodStartTime     *string `json:"PerformancePeriodStartTime"`
	PerformancePeriodEndTime       *string `json:"PerformancePeriodEndTime"`
	LeanServiceDuration            *string `json:"LeanServiceDuration"`
	LeanServiceDurationUnit        *string `json:"LeanServiceDurationUnit"`
	DistributionFunction           *string `json:"DistributionFunction"`
	SrvcSchedgIsAlignedWthOpWrkCtr *string `json:"SrvcSchedgIsAlignedWthOpWrkCtr"`
	MaOrderCompDebitCreditCode     *string `json:"MaOrderCompDebitCreditCode"`
	GoodsMovementIsAllowed         *int    `json:"GoodsMovementIsAllowed"`
	MaenanceOrderComponentBatch    *string `json:"MaenanceOrderComponentBatch"`
	QuantityIsFixed                *int    `json:"QuantityIsFixed"`
	MaOrdOpComponentGLAccount      *string `json:"MaOrdOpComponentGLAccount"`
	MaOrdOpCompCostingRelevancy    *string `json:"MaOrdOpCompCostingRelevancy"`
	MaCompAltvProdUsgeRateInPct    *string `json:"MaCompAltvProdUsgeRateInPct"`
	MaOrderOpComponentSortText     *string `json:"MaOrderOpComponentSortText"`
	MaOrdOpCompIsBulkProduct       *int    `json:"MaOrdOpCompIsBulkProduct"`
	MaterialProvisionType          *string `json:"MaterialProvisionType"`
	MaOrdOpCompAssgdWBSElmnt       *string `json:"MaOrdOpCompAssgdWBSElmnt"`
	MaOrderOpComponentPrice        *string `json:"MaOrderOpComponentPrice"`
	MaOrdOpComponentCurrency       *string `json:"MaOrdOpComponentCurrency"`
	MatlCompIsMarkedForBackflush   *int    `json:"MatlCompIsMarkedForBackflush"`
	PurchasingGroup                *string `json:"PurchasingGroup"`
	DeliveryTimeInDays             *string `json:"DeliveryTimeInDays"`
	MaOrdOpCompGdsRecipientName    *string `json:"MaOrdOpCompGdsRecipientName"`
	MaOrdOpCompUnloadingPtTxt      *string `json:"MaOrdOpCompUnloadingPtTxt"`
	GoodsReceiptDurationInWorkDays *string `json:"GoodsReceiptDurationInWorkDays"`
	PurchasingInfoRecord           *string `json:"PurchasingInfoRecord"`
	OperationLeadTimeOffset        *string `json:"OperationLeadTimeOffset"`
	OpsLeadTimeOffsetUnit          *string `json:"OpsLeadTimeOffsetUnit"`
	MaOrdOpCompRequisitioner       *string `json:"MaOrdOpCompRequisitioner"`
	MaOrdOpCompProcmtTrckgNmbr     *string `json:"MaOrdOpCompProcmtTrckgNmbr"`
	ResponsiblePurchaseOrg         *string `json:"ResponsiblePurchaseOrg"`
	MaOrdOpCompSpecialStockType    *string `json:"MaOrdOpCompSpecialStockType"`
	VariableSizeDimension1         *string `json:"VariableSizeDimension1"`
	VariableSizeDimensionUnit      *string `json:"VariableSizeDimensionUnit"`
	VariableSizeCompFormulaKey     *string `json:"VariableSizeCompFormulaKey"`
	VariableSizeDimension2         *string `json:"VariableSizeDimension2"`
	NumberOfVariableSizeItem       *int    `json:"NumberOfVariableSizeItem"`
	VariableSizeDimension3         *string `json:"VariableSizeDimension3"`
	VariableSizeItemQuantity       *string `json:"VariableSizeItemQuantity"`
	VariableSizeComponentUnit      *string `json:"VariableSizeComponentUnit"`
	RqmtDateIsEnteredManually      *int    `json:"RqmtDateIsEnteredManually"`
	SupplierProduct                *string `json:"SupplierProduct"`
	MaOrdCompPurOutlineAgrmtItm    *string `json:"MaOrdCompPurOutlineAgrmtItm"`
	MaOrderComponenternalID        *string `json:"MaOrderComponenternalID"`
	PurchaseRequisition            *string `json:"PurchaseRequisition"`
	PurchaseRequisitionItem        *string `json:"PurchaseRequisitionItem"`
	OverallLimitAmount             *string `json:"OverallLimitAmount"`
	ExpectedOverallLimitAmount     *string `json:"ExpectedOverallLimitAmount"`
}
type Operation struct {
	MaintenanceOrder               string  `json:"MaintenanceOrder"`
	MaintenanceOrderOperation      string  `json:"MaintenanceOrderOperation"`
	MaintenanceOrderSubOperation   string  `json:"MaintenanceOrderSubOperation"`
	OperationControlKey            *string `json:"OperationControlKey"`
	OperationWorkCenterInternalID  *string `json:"OperationWorkCenterInternalID"`
	WorkCenter                     *string `json:"WorkCenter"`
	Plant                          *string `json:"Plant"`
	OperationStandardTextCode      *string `json:"OperationStandardTextCode"`
	OperationDescription           *string `json:"OperationDescription"`
	MaintOrderRoutingNumber        *string `json:"MaintOrderRoutingNumber"`
	MaintenanceOrderRoutingNode    *string `json:"MaintenanceOrderRoutingNode"`
	SuperiorOperationInternalID    *string `json:"SuperiorOperationInternalID"`
	OperationWorkCenterTypeCode    *string `json:"OperationWorkCenterTypeCode"`
	Language                       *string `json:"Language"`
	NumberOfTimeTickets            *string `json:"NumberOfTimeTickets"`
	OperationPurgInfoRecdSearchTxt *string `json:"OperationPurgInfoRecdSearchTxt"`
	OperationSupplier              *string `json:"OperationSupplier"`
	CostElement                    *string `json:"CostElement"`
	OperationPurchasingInfoRecord  *string `json:"OperationPurchasingInfoRecord"`
	PurchasingOrganization         *string `json:"PurchasingOrganization"`
	PurchasingGroup                *string `json:"PurchasingGroup"`
	ProductGroup                   *string `json:"ProductGroup"`
	OpPurchaseOutlineAgreement     *string `json:"OpPurchaseOutlineAgreement"`
	OpPurchaseOutlineAgreementItem *string `json:"OpPurchaseOutlineAgreementItem"`
	OperationRequisitionerName     *string `json:"OperationRequisitionerName"`
	OperationTrackingNumber        *string `json:"OperationTrackingNumber"`
	NumberOfCapacities             *int    `json:"NumberOfCapacities "`
	OperationWorkPercent           *int    `json:"OperationWorkPercent "`
	OperationCalculationControl    *string `json:"OperationCalculationControl"`
	ActivityType                   *string `json:"ActivityType"`
	OperationSystemCondition       *string `json:"OperationSystemCondition"`
	OperationGoodsRecipientName    *string `json:"OperationGoodsRecipientName"`
	OperationUnloadingPointName    *string `json:"OperationUnloadingPointName"`
	OperationPersonResponsible     *string `json:"OperationPersonResponsible"`
	DeliveryTimeInDays             *string `json:"DeliveryTimeInDays"`
	MaintOrderOperationDuration    *string `json:"MaintOrderOperationDuration"`
	MaintOrdOperationDurationUnit  *string `json:"MaintOrdOperationDurationUnit"`
	OpBscStartDateConstraintType   *string `json:"OpBscStartDateConstraintType"`
	OpBscEndDateConstraintType     *string `json:"OpBscEndDateConstraintType"`
	MaintOrdOperationWorkDuration  *string `json:"MaintOrdOperationWorkDuration"`
	MaintOrdOpWorkDurationUnit     *string `json:"MaintOrdOpWorkDurationUnit"`
	MaintOrdOpConstraintStrtDteTme *string `json:"MaintOrdOpConstraintStrtDteTme"`
	ConstraintDateForBscStartDate  *string `json:"ConstraintDateForBscStartDate"`
	ConstraintTimeForBscStartTime  *string `json:"ConstraintTimeForBscStartTime"`
	MaintOrdOpCstrtFinishDteTme    *string `json:"MaintOrdOpCstrtFinishDteTme"`
	ConstraintDateForBscFinishDate *string `json:"ConstraintDateForBscFinishDate"`
	ConstraintTimeForBscFinishTime *string `json:"ConstraintTimeForBscFinishTime"`
	MaintOrdOperationExecutionRate *string `json:"MaintOrdOperationExecutionRate"`
	Equipment                      *string `json:"Equipment"`
	FunctionalLocation             *string `json:"FunctionalLocation"`
	MaintenanceActivityType        *string `json:"MaintenanceActivityType"`
	BusinessArea                   *string `json:"BusinessArea"`
	ProfitCenter                   *string `json:"ProfitCenter"`
	FunctionalArea                 *string `json:"FunctionalArea"`
	MaintControllingObjectClass    *string `json:"MaintControllingObjectClass"`
	WrkCtrIntCapRqmtsDistr         *string `json:"WrkCtrIntCapRqmtsDistr"`
	MaintOrdOperationOverheadCode  *string `json:"MaintOrdOperationOverheadCode"`
	MaintOrderOperationQuantity    *string `json:"MaintOrderOperationQuantity"`
	MaintOrdOperationQuantityUnit  *string `json:"MaintOrdOperationQuantityUnit"`
	Assembly                       *string `json:"Assembly"`
	MaintOperationExecStageCode    *string `json:"MaintOperationExecStageCode"`
	WBSElement                     *string `json:"WBSElement"`
	IsMarkedForDeletion            *string `json:"IsMarkedForDeletion tinyint"`
	MaintOrderOperationInternalID  *string `json:"MaintOrderOperationInternalID"`
	MaintenanceObjectListItem      *int    `json:"MaintenanceObjectListItem "`
	PurchaseRequisition            *string `json:"PurchaseRequisition"`
	PurchaseRequisitionItem        *string `json:"PurchaseRequisitionItem"`
	OpErlstSchedldExecStrtDte      *string `json:"OpErlstSchedldExecStrtDte"`
	OpErlstSchedldExecStrtTme      *string `json:"OpErlstSchedldExecStrtTme"`
	OpErlstSchedldExecEndDte       *string `json:"OpErlstSchedldExecEndDte"`
	OpErlstSchedldExecEndTme       *string `json:"OpErlstSchedldExecEndTme"`
	OpLtstSchedldExecStrtDte       *string `json:"OpLtstSchedldExecStrtDte"`
	OpLtstSchedldExecStrtTme       *string `json:"OpLtstSchedldExecStrtTme"`
	OpLtstSchedldExecEndDte        *string `json:"OpLtstSchedldExecEndDte"`
	OpLtstSchedldExecEndTme        *string `json:"OpLtstSchedldExecEndTme"`
	OpActualExecutionStartDate     *string `json:"OpActualExecutionStartDate"`
	OpActualExecutionStartTime     *string `json:"OpActualExecutionStartTime"`
	OpActualExecutionEndDate       *string `json:"OpActualExecutionEndDate"`
	OpActualExecutionEndTime       *string `json:"OpActualExecutionEndTime"`
	ForecastWorkQuantity           *string `json:"ForecastWorkQuantity"`
	ActualWorkQuantity             *string `json:"ActualWorkQuantity"`
	MaintOrdOpProcessPhaseCode     *string `json:"MaintOrdOpProcessPhaseCode"`
	MaintOrdOpProcessSubPhaseCode  *string `json:"MaintOrdOpProcessSubPhaseCode"`
	SystemStatusText               *string `json:"SystemStatusText"`
	UserStatusText                 *string `json:"UserStatusText"`
}
