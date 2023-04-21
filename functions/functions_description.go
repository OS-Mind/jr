//Copyright © 2022 Ugo Landini <ugo.landini@gmail.com>
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in
//all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//THE SOFTWARE.

package functions

func Description(name string) (FunctionDescription, bool) {
	v, ok := funcDesc[name]
	return v, ok
}

func DescriptionMap() map[string]FunctionDescription {
	return funcDesc
}

type FunctionDescription struct {
	Name        string
	Category    string
	Description string
	Parameters  string
	Localizable bool
	Return      string
	Example     string
	Output      string
}

var funcDesc = map[string]FunctionDescription{

	"upper": {
		Name:        "upper",
		Category:    "text utilities",
		Description: "converts a string to uppercase",
		Parameters:  "text string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{upper \"hello\"}}'",
		Output:      "HELLO",
	},
	"lower": {
		Name:        "lower",
		Category:    "text utilities",
		Description: "converts a string to lowercase",
		Parameters:  "text string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{lower \"HELLO\"}}'",
		Output:      "hello",
	},
	"title": {
		Name:        "title",
		Category:    "text utilities",
		Description: "converts a string to title case",
		Parameters:  "text string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{title \"hello world\"}}'",
		Output:      "Hello World",
	},
	"trim": {
		Name:        "trim",
		Category:    "text utilities",
		Description: "trims whitespace from a string",
		Parameters:  "text string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{trim \" hello world \"}}'",
		Output:      "hello world",
	},
	"squeeze": {
		Name:        "squeeze",
		Category:    "text utilities",
		Description: "removes all spaces from a string",
		Parameters:  "text string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{squeeze \" hello   world \"}}'",
		Output:      "helloworld",
	},
	"substr": {
		Name:        "substr",
		Category:    "text utilities",
		Description: "returns a substring of a string",
		Parameters:  "from int, to int, text string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{substr 0 5 \"hello world\"}}'",
		Output:      "hello",
	},
	"first": {
		Name:        "first",
		Category:    "text utilities",
		Description: "returns the first character of a string",
		Parameters:  "text string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{first \"hello world\"}}'",
		Output:      "h",
	},
	"firstword": {
		Name:        "firstword",
		Category:    "text utilities",
		Description: "returns the first word of a text",
		Parameters:  "text string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{firstword \"hello world\"}}'",
		Output:      "hello",
	},
	"join": {
		Name:        "join",
		Category:    "text utilities",
		Description: "joins a list of strings with a separator",
		Parameters:  "strings []string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{join \"hello,\" \"world\"}}'",
		Output:      "hello,world",
	},
	"repeat": {
		Name:        "repeat",
		Category:    "text utilities",
		Description: "repeats a string a number of times",
		Parameters:  "text string, number int",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{repeat \"hello\" 3}}'",
		Output:      "hellohellohello",
	},
	"squeezechars": {
		Name:        "squeezechars",
		Category:    "text utilities",
		Description: "removes all instances of a character from a string",
		Parameters:  "set string, chars string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{squeezechars \"hello world\" \"l\"}}'",
		Output:      "heo word",
	},
	"replaceall": {
		Name:        "replaceall",
		Category:    "text utilities",
		Description: "replaces all instances of a string with another string",
		Parameters:  "set string, original string, replaced string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{replaceall \"hello world\" \"hello\" \"goodbye\"}}'",
		Output:      "goodbye world",
	},
	"trimchars": {
		Name:        "trimchars",
		Category:    "text utilities",
		Description: "trims all characters in the given set from the beginning and end of a string",
		Parameters:  "set string, chars string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{trimchars \"hello world\" \"hld\"}}'",
		Output:      "ello wor",
	},
	"atoi": {
		Name:        "atoi",
		Category:    "text utilities",
		Description: "converts a string to an integer",
		Parameters:  "string string",
		Localizable: false,
		Return:      "int",
		Example:     "jr run --template '{{atoi \"123\"}}'",
		Output:      "123",
	},
	"split": {
		Name:        "split",
		Category:    "text utilities",
		Description: "splits a string into a list of strings",
		Parameters:  "text string, separator string",
		Localizable: false,
		Return:      "[]string",
		Example:     "jr run --template '{{split \"hello,world\" \",\"}}'",
		Output:      "[hello world]",
	},
	"markov": {
		Name:        "markov",
		Category:    "text utilities",
		Description: "generates a markov chain from a string",
		Parameters:  "chain int, text string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{markov 10 \"hello world\"}}'",
		Output:      "hello world",
	},
	"lorem": {
		Name:        "lorem",
		Category:    "text utilities",
		Description: "generates a Lorem ipsum string",
		Parameters:  "words int",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{lorem 10}}'",
		Output:      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.Fusce elit magna.",
	},
	"sentence": {
		Name:        "sentence",
		Category:    "text utilities",
		Description: "generates a random Sentence of n words",
		Parameters:  "words int",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{sentence 15}}'",
		Output:      "Alice was not going to happen next. First, she tried to curtsey as she fell",
	},
	"sentence_prefix": {
		Name:        "sentence_prefix",
		Category:    "text utilities",
		Description: "generates a random Sentence of n words, with a prefix length",
		Parameters:  "prefix int, length int",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{sentence_prefix 3 15}}'",
		Output:      "Alice was beginning to get very tired of sitting by her sister on the bank.",
	},
	"regex": {
		Name:        "regex",
		Category:    "text utilities",
		Description: "returns a random string matching the Regex",
		Parameters:  "regex string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{regex \"[a-z]{5}\"}}'",
		Output:      "xxlbh",
	},
	"random": {
		Name:        "random",
		Category:    "text utilities",
		Description: "returns a random string from a list of strings",
		Parameters:  "list []string",
		Localizable: false,
		Return:      "string",
		Example:     "",
		Output:      "hello",
	},
	"randoms": {
		Name:        "randoms",
		Category:    "text utilities",
		Parameters:  "list string",
		Localizable: false,
		Description: "returns a random strings from a | separated list string",
		Return:      "string",
		Example:     "jr run --template '{{randoms \"ALPHA|BETA|GAMMA|DELTA\"}}'",
		Output:      "BETA",
	},
	"random_string": {
		Name:        "random_string",
		Category:    "text utilities",
		Description: "returns a random string long between min and max characters",
		Parameters:  "min int, max int, vocabulary string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{random_string 3 10}}'",
		Output:      "YBCEjxmn",
	},
	"random_string_vocabulary": {
		Name:        "random_string_vocabulary",
		Category:    "text utilities",
		Description: "returns a random string long between min and max characters using a vocabulary",
		Parameters:  "min int, max int, vocabulary string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{random_string_vocabulary 3 10 \"hello world\"}}'",
		Output:      "llolh",
	},
	"from": {
		Name:        "from",
		Category:    "text utilities",
		Description: "returns a random string from a list of strings in a file. Files are in '$HOME/.jr/data/locale'",
		Parameters:  "set string",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{from \"actor\"}}'",
		Output:      "John Travolta",
	},
	"from_at": {
		Name:        "from_at",
		Category:    "text utilities",
		Description: "returns a string at a given position in a list of strings in a file. Files are in '$HOME/.jr/data/locale'",
		Parameters:  "index int",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{from_at \"actor\" 3}}'",
		Output:      "James Dean",
	},
	"from_shuffle": {
		Name:        "from_shuffle",
		Category:    "text utilities",
		Description: "returns a shuffled list of strings in a file. Files are in '$HOME/.jr/data/locale'",
		Parameters:  "set string",
		Localizable: true,
		Return:      "[]string",
		Example:     "jr run --template '{{from_shuffle \"state_short\"}}'",
		Output:      "[ND IL MO WA NC SD MS PA AZ HI DE SC WI WV TN AL MA IA NH NV OH VA WY MT MN NM LA OK IN CA OR VT MD NY RI UT AK NE AR CO FL ID KY TX ME GA NJ MI KS CT]",
	},
	"from_n": {
		Name:        "from_n",
		Category:    "text utilities",
		Description: "return a subset of elements in a list of string in a file. Files are in '$HOME/.jr/data/locale'",
		Parameters:  "set string, number int",
		Localizable: true,
		Return:      "[]string",
		Example:     "jr run --template '{{from_n \"State\" 5}}'",
		Output:      "[West Virginia Idaho Maryland New Hampshire Wyoming]",
	},
	"add": {
		Name:        "add",
		Category:    "math",
		Parameters:  "first int, second int",
		Localizable: false,
		Description: "adds two numbers",
		Return:      "int",
		Example:     "jr run --template '{{add 1 2}}'",
		Output:      "3",
	},
	"sub": {
		Name:        "sub",
		Category:    "math",
		Description: "subtracts two numbers",
		Parameters:  "first int, second int",
		Localizable: false,
		Return:      "int",
		Example:     "jr run --template '{{sub 1 2}}'",
		Output:      "-1",
	},
	"div": {
		Name:        "div",
		Category:    "math",
		Description: "divides two numbers",
		Parameters:  "first int, second int",
		Localizable: false,
		Return:      "int",
		Example:     "jr run --template '{{div 10 2}}'",
		Output:      "5",
	},
	"mul": {
		Name:        "mul",
		Category:    "math",
		Description: "multiplies two numbers",
		Parameters:  "first int, second int",
		Localizable: false,
		Return:      "int",
		Example:     "jr run --template '{{mul 10 2}}'",
		Output:      "20",
	},
	"mod": {
		Name:        "mod",
		Category:    "math",
		Description: "returns the remainder of two numbers",
		Parameters:  "first int, second int",
		Localizable: false,
		Return:      "int",
		Example:     "jr run --template '{{mod 10 2}}'",
		Output:      "0",
	},
	"max": {
		Name:        "max",
		Category:    "math",
		Description: "returns the maximum of two numbers",
		Parameters:  "first int, second int",
		Localizable: false,
		Return:      "int",
		Example:     "jr run --template '{{max 10 2}}'",
		Output:      "10",
	},
	"min": {
		Name:        "min",
		Category:    "math",
		Description: "returns the minimum of two numbers",
		Parameters:  "min int, max int",
		Localizable: false,
		Return:      "int",
		Example:     "jr run --template '{{min 10 2}}'",
		Output:      "2",
	},
	"integer": {
		Name:        "integer",
		Category:    "math",
		Description: "returns a random integer between min and max",
		Parameters:  "min int, max int",
		Localizable: false,
		Return:      "int",
		Example:     "jr run --template '{{integer 10 20}}'",
		Output:      "13",
	},
	"integer64": {
		Name:        "integer64",
		Category:    "math",
		Description: "returns a random int64 between min and max",
		Parameters:  "min int64, max int64",
		Localizable: false,
		Return:      "int64",
		Example:     "jr run --template '{{integer64 10 20}}'",
		Output:      "13",
	},
	"floating": {
		Name:        "floating",
		Category:    "math",
		Description: "returns a random float64 between min and max",
		Parameters:  "min float64, max float64",
		Localizable: false,
		Return:      "float64",
		Example:     "jr run --template '{{floating 10 20}}'",
		Output:      "13.123",
	},
	"ip": {
		Name:        "ip",
		Category:    "network",
		Description: "returns a random Ip Address matching the given cidr",
		Parameters:  "cidr string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{ip \"10.2.0.0/16\"}}'",
		Output:      "10.2.55.217",
	},
	"ipv6": {
		Name:        "ipv6",
		Category:    "network",
		Description: "returns a random Ipv6 Address",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{ipv6}}'",
		Output:      "2001:db8:85a3:8d3:1319:8a2e:370:7348",
	},
	"mac": {
		Name:        "mac",
		Category:    "network",
		Description: "returns a random mac Address",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{mac}}'",
		Output:      "7e:8e:75:a5:0a:85",
	},
	"http_method": {
		Name:        "http_method",
		Category:    "network",
		Description: "returns a random http method",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{http_method}}'",
		Output:      "GET",
	},
	"ip_known_protocol": {
		Name:        "ip_known_protocol",
		Category:    "network",
		Description: "returns a random known protocol",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{ip_known_protocol}}'",
		Output:      "tcp",
	},
	"ip_known_port": {
		Name:        "ip_known_port",
		Category:    "network",
		Description: "returns a random known port number",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{ip_known_port}}'",
		Output:      "80",
	},
	"password": {
		Name:        "password",
		Category:    "security",
		Description: "returns a random Password of given length, memorable, and with prefix and suffix",
		Parameters:  "length int, memorable bool, prefix string, suffix string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{password 10 true \"!\" \"?\"}}'",
		Output:      "!itAPYgivIH?",
	},
	"useragent": {
		Name:        "useragent",
		Category:    "security",
		Description: "returns a random user agent",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{useragent}}'",
		Output:      "Mozilla/5.0 (Android 10) AppleWebKit/592.64 (KHTML, like Gecko) Chrome Mobile/7.3.5.7 Mobile Safari/9.10",
	},
	"unix_time_stamp": {
		Name:        "unix_time_stamp",
		Category:    "time",
		Description: "returns a random unix timestamp not older than the given number of days",
		Parameters:  "days int",
		Localizable: false,
		Return:      "int64",
		Example:     "jr run --template '{{unix_time_stamp 10}}'",
		Output:      "1679703304",
	},
	"name": {
		Name:        "name",
		Category:    "people",
		Description: "returns a random Name",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{name}}'",
		Output:      "Lisa",
	},
	"name_m": {
		Name:        "name_m",
		Category:    "people",
		Description: "returns a random male Name",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{name_m}}'",
		Output:      "John",
	},
	"name_f": {
		Name:        "name_f",
		Category:    "people",
		Description: "returns a random female Name",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{name_f}}'",
		Output:      "Lisa",
	},
	"middle_name": {
		Name:        "middle_name",
		Category:    "people",
		Description: "returns a random middle Name",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{middle_name}}'",
		Output:      "J",
	},
	"surname": {
		Name:        "surname",
		Category:    "people",
		Description: "returns a random surname",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{surname}}'",
		Output:      "Wright",
	},
	"username": {
		Name:        "username",
		Category:    "people",
		Description: "returns a random Username using Name, Surname and a length",
		Parameters:  "name string, surname string, length int",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{username \"barack\" \"obama\" 12 }}'",
		Output:      "barackobama75",
	},
	"address": {
		Name:        "address",
		Category:    "people",
		Description: "returns a random address",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{address}}'",
		Output:      "1234 Main Street",
	},
	"country": {
		Name:        "country",
		Category:    "people",
		Description: "returns a random ISO 3166 country code",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{country}}'",
		Output:      "IT",
	},
	"capital": {
		Name:        "capital",
		Category:    "people",
		Description: "returns a random capital",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{capital}}'",
		Output:      "Phoenix",
	},
	"capital_at": {
		Name:        "capital_at",
		Category:    "people",
		Description: "returns Capital at given index",
		Parameters:  "index int",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{capital_at 4}}'",
		Output:      "Sacramento",
	},
	"state": {
		Name:        "state",
		Category:    "people",
		Description: "returns a random state",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{state}}'",
		Output:      "New York",
	},
	"state_at": {
		Name:        "state_at",
		Category:    "people",
		Description: "returns state at given index",
		Parameters:  "index int",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{state_at 3}}'",
		Output:      "Arkansas",
	},
	"state_short": {
		Name:        "state_short",
		Category:    "people",
		Description: "returns a random short State",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{state_short}}'",
		Output:      "KY",
	},
	"state_short_at": {
		Name:        "state_short_at",
		Category:    "people",
		Description: "returns short State at given index",
		Parameters:  "index int",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{state_short_at 3}}'",
		Output:      "AR",
	},
	"zip": {
		Name:        "zip",
		Category:    "people",
		Description: "returns a random zip code",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{zip}}'",
		Output:      "21401",
	},
	"zip_at": {
		Name:        "zip_at",
		Category:    "people",
		Description: "returns Zip code at given index",
		Parameters:  "index int",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{zip_at 3}}'",
		Output:      "72201",
	},
	"company": {
		Name:        "company",
		Category:    "people",
		Description: "returns a random company name",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{company}}'",
		Output:      "Umbrella Corporation",
	},
	"mail_provider": {
		Name:        "mail_provider",
		Category:    "people",
		Description: "returns a random mail provider",
		Parameters:  "",
		Localizable: true,
		Return:      "string",
		Example:     "jr run --template '{{mail_provider}}'",
		Output:      "gmail.com",
	},
	"key": {
		Name:        "key",
		Category:    "utilities",
		Description: "returns a random key string using a prefix and a length",
		Parameters:  "prefix string, length int",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{key \"KEY\" 20}}'",
		Output:      "KEY4",
	},
	"uuid": {
		Name:        "uuid",
		Category:    "utilities",
		Description: "returns a random uuid",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{uuid}}'",
		Output:      "a6da3ed0-5fcb-4bb8-a6aa-654120a1e6e3",
	},
	"bool": {
		Name:        "bool",
		Category:    "utilities",
		Description: "returns a random boolean",
		Parameters:  "",
		Localizable: false,
		Return:      "bool",
		Example:     "jr run --template '{{bool}}'",
		Output:      "true",
	},
	"yesorno": {
		Name:        "yesorno",
		Category:    "utilities",
		Description: "returns a random yes or no",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{yesorno}}'",
		Output:      "yes",
	},
	"array": {
		Name:        "array",
		Category:    "utilities",
		Description: "returns an empty array of given size",
		Parameters:  "size int",
		Localizable: false,
		Return:      "array",
		Example:     "jr run --template '{{array 5}}'",
		Output:      "[0,0,0,0,0]",
	},
	"counter": {
		Name:        "counter",
		Category:    "utilities",
		Description: "returns a named counter, starting at n incrementing by i",
		Parameters:  "name string, start int, step int",
		Localizable: false,
		Return:      "int",
		Example:     "jr run -n 5 --template '{{counter \"mycounter\" 0 1}}'",
		Output:      "\n0\n1\n2\n3\n4",
	},
	"cusip": {
		Name:        "cusip",
		Category:    "finance",
		Description: "returns a valid 9 characters cusip code",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{cusip}}'",
		Output:      "DWNFYN9W2",
	},
	"sedol": {
		Name:        "sedol",
		Category:    "finance",
		Description: "returns a valid 7 characters sedol code",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{sedol}}'",
		Output:      "",
	},
	"valor": {
		Name:        "valor",
		Category:    "finance",
		Description: "returns a valid 5-9 digits valor code",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{valor}}'",
		Output:      "",
	},
	"isin": {
		Name:        "isin",
		Category:    "finance",
		Description: "returns a valid 12 characters isin code",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{isin}}'",
		Output:      "",
	},
	"wkn": {
		Name:        "wkn",
		Category:    "finance",
		Description: "returns a valid 6 characters wkn code",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{wkn}}'",
		Output:      "",
	},
	"card": {
		Name:        "card",
		Category:    "finance",
		Description: "returns a random credit card number",
		Parameters:  "issuer string",
		Localizable: false,
		Return:      "issuer string",
		Example:     "jr run --template '{{card \"amex\"}}'",
		Output:      "376794009305701",
	},
	"cf": {
		Name:        "cf",
		Category:    "people",
		Description: "returns an Italian codice fiscale",
		Parameters:  "name string, surname string, gender string, birthdate string, birthcity string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{cf \"Mario\" \"Rossi\" \"M\" \"1970-01-30\" \"Roma\"}}'",
		Output:      "RSSMRA70A30H501W",
	},
	"ssn": {
		Name:        "ssn",
		Category:    "people",
		Description: "returns a random ssn (Social Security Number)",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{ssn}}'",
		Output:      "834-76-1234",
	},
	"account": {
		Name:        "account",
		Category:    "finance",
		Description: "Account returns a random account number of given length",
		Parameters:  "length int",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{account 10}}'",
		Output:      "6192117146",
	},
	"amount": {
		Name:        "amount",
		Category:    "finance",
		Description: "Amount returns an amount of money between min and max, and given currency",
		Parameters:  "min float32, max float32, currency string",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{account 10 1000 \"$\"}}'",
		Output:      "$7409.66",
	},
	"swift": {
		Name:        "swift",
		Category:    "finance",
		Description: "Swift returns a swift/bic code",
		Parameters:  "",
		Localizable: false,
		Return:      "string",
		Example:     "jr run --template '{{swift}}'",
		Output:      "KZMTMP84448",
	},
}
