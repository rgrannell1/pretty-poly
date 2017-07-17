
package pretty_poly




import "encoding/json"





type Log struct {
	user_message string
	level        string
}

func (log *Log) String( ) string {

	val, err := json.Marshal(log)

	if err != nil {
		panic(err)
	}

	return string(val)

}
