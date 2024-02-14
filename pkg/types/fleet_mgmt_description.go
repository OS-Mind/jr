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

type FleetMgmtDescription struct {
	Vehicle_id int32 `json:"vehicle_id"`

	Driver_name string `json:"driver_name"`

	License_plate string `json:"license_plate"`
}

const FleetMgmtDescriptionAvroCRC64Fingerprint = "\x03\xfe;\x14\xaco\xfa\x88"

func NewFleetMgmtDescription() FleetMgmtDescription {
	r := FleetMgmtDescription{}
	return r
}

func DeserializeFleetMgmtDescription(r io.Reader) (FleetMgmtDescription, error) {
	t := NewFleetMgmtDescription()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFleetMgmtDescriptionFromSchema(r io.Reader, schema string) (FleetMgmtDescription, error) {
	t := NewFleetMgmtDescription()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFleetMgmtDescription(r FleetMgmtDescription, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Vehicle_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Driver_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.License_plate, w)
	if err != nil {
		return err
	}
	return err
}

func (r FleetMgmtDescription) Serialize(w io.Writer) error {
	return writeFleetMgmtDescription(r, w)
}

func (r FleetMgmtDescription) Schema() string {
	return "{\"fields\":[{\"name\":\"vehicle_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":9999,\"min\":1000}},\"type\":\"int\"}},{\"name\":\"driver_name\",\"type\":{\"arg.properties\":{\"options\":[\"Frieda Baldi\",\"Cherrita Gallaccio\",\"Matt Cleugh\",\"Dulciana Murfill\",\"Germayne Streetley\",\"Brenna Woolfall\",\"Gerhardt Tenbrug\",\"Hayley Tuma\",\"Winny Cadigan\",\"Bonnibelle Macek\",\"Lionel Byneth\",\"Trev Roper\",\"Lena MacFadzean\",\"Benton Allcorn\",\"Avis Moyler\",\"Marchall Rochewell\",\"Adele Bohl\",\"Barnett Mcall\",\"Frieda Pirrone\",\"Pattin Eringey\",\"Kalila Fewings\",\"Giacobo Beuscher\",\"Rozalin Hair\",\"Egon Beagan\",\"Owen Strotton\",\"Fernando Rosensaft\",\"Carleton Gwyther\",\"Kata Coll\",\"Rossie Hobben\",\"Stephanie Gookey\",\"Robyn Milazzo\",\"Tilda O'Lunney\",\"Nolan Kidney\",\"Jori Ottiwill\",\"Benito Graveson\",\"Zechariah Wrate\",\"Chelsae Napton\",\"Jeremy Heffernon\",\"Derk McAviy\",\"Constantin Mears\",\"Fitz Ballin\",\"Essy Bettles\",\"Gene Klemt\",\"Nikolai Arnopp\",\"Gustave Westhofer\",\"Simona Mayhow\",\"Cort Bainbridge\",\"Sibyl Vockins\",\"Andriette Gaze\",\"Shaughn De Simoni\",\"Nathaniel Hallowell\",\"Charley Dudill\",\"Cirstoforo Joblin\",\"Hyacinthia Kinastan\",\"Dur Lasselle\",\"Gay Chadburn\",\"Livvie Hawyes\",\"Aldrich MacVay\",\"Riva Rossant\",\"Johanna Reichartz\",\"Trent Gantlett\",\"Aryn Haskell\",\"Byrann Barock\",\"Gerda Cleugher\",\"Sonnie Guildford\",\"Vergil Borge\",\"Lurline Rocco\",\"Geoff Eddy\",\"Zea Leighton\",\"Leif Baden\",\"Quint Bidgod\",\"Talbot Cashell\",\"Sheridan Foulsham\",\"Camile Shrimplin\",\"Marcel Nayshe\",\"Lea Murrish\",\"Lucais Midson\",\"Zeb Rylatt\",\"Nertie Zuker\",\"Babara Henderson\",\"Electra Ridgley\",\"Jere Standingford\",\"Cyril Yellowlea\",\"Isadora Peegrem\",\"Caria Smewings\",\"Karena Kauffman\",\"Haywood Snowball\",\"Winslow Starcks\",\"Alis Ponton\",\"Marietta Lezemere\",\"Emilee Broadbridge\",\"Faye Beaument\",\"Shannah Beatson\",\"West Doy\",\"Chryste Wren\",\"Trumann Labba\",\"Anatollo Beckwith\",\"Konstanze Dunsford\",\"Raychel Roset\",\"Heindrick Ravenscroft\"]},\"type\":\"string\"}},{\"name\":\"license_plate\",\"type\":{\"arg.properties\":{\"regex\":\"[A-Z]{2}\\\\d{6}[A-Z]\"},\"type\":\"string\"}}],\"name\":\"fleet_mgmt.FleetMgmtDescription\",\"type\":\"record\"}"
}

func (r FleetMgmtDescription) SchemaName() string {
	return "fleet_mgmt.FleetMgmtDescription"
}

func (_ FleetMgmtDescription) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FleetMgmtDescription) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FleetMgmtDescription) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FleetMgmtDescription) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FleetMgmtDescription) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FleetMgmtDescription) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FleetMgmtDescription) SetString(v string)   { panic("Unsupported operation") }
func (_ FleetMgmtDescription) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FleetMgmtDescription) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Vehicle_id}

		return w

	case 1:
		w := types.String{Target: &r.Driver_name}

		return w

	case 2:
		w := types.String{Target: &r.License_plate}

		return w

	}
	panic("Unknown field index")
}

func (r *FleetMgmtDescription) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *FleetMgmtDescription) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ FleetMgmtDescription) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FleetMgmtDescription) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FleetMgmtDescription) HintSize(int)                     { panic("Unsupported operation") }
func (_ FleetMgmtDescription) Finalize()                        {}

func (_ FleetMgmtDescription) AvroCRC64Fingerprint() []byte {
	return []byte(FleetMgmtDescriptionAvroCRC64Fingerprint)
}

func (r FleetMgmtDescription) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["vehicle_id"], err = json.Marshal(r.Vehicle_id)
	if err != nil {
		return nil, err
	}
	output["driver_name"], err = json.Marshal(r.Driver_name)
	if err != nil {
		return nil, err
	}
	output["license_plate"], err = json.Marshal(r.License_plate)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *FleetMgmtDescription) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["vehicle_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Vehicle_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for vehicle_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["driver_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Driver_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for driver_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["license_plate"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.License_plate); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for license_plate")
	}
	return nil
}
