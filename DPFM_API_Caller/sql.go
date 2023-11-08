package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-maintenance-order-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-maintenance-order-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) HeaderRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE header.MaintenanceOrder = %s ", input.Header.MaintenanceOrder),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	header.MaintenanceOrder
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_maintenance_order_header_data as header 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemsRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ObjectListItem {
	where := fmt.Sprintf("WHERE item.MaintenanceOrder IS NOT NULL\nAND header.MaintenanceOrder = %d", input.Header.MaintenanceOrder)
	rows, err := c.db.Query(
		`SELECT 
			item.MaintenanceOrder, item.ObjectListItem
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_maintenance_order_item_data as item
		INNER JOIN DataPlatformMastersAndTransactionsMysqlKube.data_platform_maintenance_order_header_data as header
		ON header.MaintenanceOrder = item.MaintenanceOrder ` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
