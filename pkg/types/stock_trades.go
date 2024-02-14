// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaign_finance.avsc
 *     click_stream.avsc
 *     click_stream_codes.avsc
 *     click_stream_users.avsc
 *     credit_cards.avsc
 *     csv_product.avsc
 *     csv_user.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurance_customer_activity.avsc
 *     insurance_customers.avsc
 *     insurance_offers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizza_orders.avsc
 *     pizza_orders_cancelled.avsc
 *     pizza_orders_completed.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siem_logs.avsc
 *     stock_trades.avsc
 *     stores.avsc
 *     syslog_logs.avsc
 *     transactions.avsc
 *     user.avsc
 *     users.avsc
 *     users_array_map.avsc
 */
package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

// Defines a hypothetical stock trade using some known test stock symbols.
type StockTrades struct {
	// A simulated trade side (buy or sell or short)
	Side string `json:"side"`
	// A simulated random quantity of the trade
	Quantity int32 `json:"quantity"`
	// Simulated stock symbols
	Symbol string `json:"symbol"`
	// A simulated random trade price in pennies
	Price int32 `json:"price"`
	// Simulated accounts assigned to the trade
	Account string `json:"account"`
	// The simulated user who executed the trade
	Userid string `json:"userid"`
}

const StockTradesAvroCRC64Fingerprint = "\xc1ƴ\x96C\xb9\xff\xce"

func NewStockTrades() StockTrades {
	r := StockTrades{}
	return r
}

func DeserializeStockTrades(r io.Reader) (StockTrades, error) {
	t := NewStockTrades()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeStockTradesFromSchema(r io.Reader, schema string) (StockTrades, error) {
	t := NewStockTrades()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeStockTrades(r StockTrades, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Side, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Quantity, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Symbol, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Price, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Account, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Userid, w)
	if err != nil {
		return err
	}
	return err
}

func (r StockTrades) Serialize(w io.Writer) error {
	return writeStockTrades(r, w)
}

func (r StockTrades) Schema() string {
	return "{\"doc\":\"Defines a hypothetical stock trade using some known test stock symbols.\",\"fields\":[{\"doc\":\"A simulated trade side (buy or sell or short)\",\"name\":\"side\",\"type\":{\"arg.properties\":{\"options\":[\"BUY\",\"SELL\"]},\"type\":\"string\"}},{\"doc\":\"A simulated random quantity of the trade\",\"name\":\"quantity\",\"type\":{\"arg.properties\":{\"range\":{\"max\":5000,\"min\":1}},\"type\":\"int\"}},{\"doc\":\"Simulated stock symbols\",\"name\":\"symbol\",\"type\":{\"arg.properties\":{\"options\":[\"ZBZX\",\"ZJZZT\",\"ZTEST\",\"ZVV\",\"ZVZZT\",\"ZWZZT\",\"ZXZZT\"]},\"type\":\"string\"}},{\"doc\":\"A simulated random trade price in pennies\",\"name\":\"price\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1000,\"min\":5}},\"type\":\"int\"}},{\"doc\":\"Simulated accounts assigned to the trade\",\"name\":\"account\",\"type\":{\"arg.properties\":{\"options\":[\"ABC123\",\"LMN456\",\"XYZ789\"]},\"type\":\"string\"}},{\"doc\":\"The simulated user who executed the trade\",\"name\":\"userid\",\"type\":{\"arg.properties\":{\"regex\":\"User_[1-9]\"},\"type\":\"string\"}}],\"name\":\"ksql.StockTrades\",\"type\":\"record\"}"
}

func (r StockTrades) SchemaName() string {
	return "ksql.StockTrades"
}

func (_ StockTrades) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ StockTrades) SetInt(v int32)       { panic("Unsupported operation") }
func (_ StockTrades) SetLong(v int64)      { panic("Unsupported operation") }
func (_ StockTrades) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ StockTrades) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ StockTrades) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ StockTrades) SetString(v string)   { panic("Unsupported operation") }
func (_ StockTrades) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *StockTrades) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Side}

		return w

	case 1:
		w := types.Int{Target: &r.Quantity}

		return w

	case 2:
		w := types.String{Target: &r.Symbol}

		return w

	case 3:
		w := types.Int{Target: &r.Price}

		return w

	case 4:
		w := types.String{Target: &r.Account}

		return w

	case 5:
		w := types.String{Target: &r.Userid}

		return w

	}
	panic("Unknown field index")
}

func (r *StockTrades) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *StockTrades) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ StockTrades) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ StockTrades) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ StockTrades) HintSize(int)                     { panic("Unsupported operation") }
func (_ StockTrades) Finalize()                        {}

func (_ StockTrades) AvroCRC64Fingerprint() []byte {
	return []byte(StockTradesAvroCRC64Fingerprint)
}

func (r StockTrades) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["side"], err = json.Marshal(r.Side)
	if err != nil {
		return nil, err
	}
	output["quantity"], err = json.Marshal(r.Quantity)
	if err != nil {
		return nil, err
	}
	output["symbol"], err = json.Marshal(r.Symbol)
	if err != nil {
		return nil, err
	}
	output["price"], err = json.Marshal(r.Price)
	if err != nil {
		return nil, err
	}
	output["account"], err = json.Marshal(r.Account)
	if err != nil {
		return nil, err
	}
	output["userid"], err = json.Marshal(r.Userid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *StockTrades) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["side"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Side); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for side")
	}
	val = func() json.RawMessage {
		if v, ok := fields["quantity"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Quantity); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for quantity")
	}
	val = func() json.RawMessage {
		if v, ok := fields["symbol"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Symbol); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for symbol")
	}
	val = func() json.RawMessage {
		if v, ok := fields["price"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Price); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for price")
	}
	val = func() json.RawMessage {
		if v, ok := fields["account"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Account); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for account")
	}
	val = func() json.RawMessage {
		if v, ok := fields["userid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Userid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for userid")
	}
	return nil
}
