package faker

var (
	currenciesCountries = []string{"AFGHANISTAN", "ALBANIA", "ALGERIA", "AMERICAN SAMOA", "ANDORRA", "ANGOLA", "ANGUILLA", "ANTARCTICA", "ANTIGUA AND BARBUDA", "ARGENTINA", "ARMENIA", "ARUBA", "AUSTRALIA", "AUSTRIA", "AZERBAIJAN", "BAHAMAS (THE)", "BAHRAIN", "BANGLADESH", "BARBADOS", "BELARUS", "BELGIUM", "BELIZE", "BENIN", "BERMUDA", "BHUTAN", "BHUTAN", "BOLIVIA (PLURINATIONAL STATE OF)", "BOLIVIA (PLURINATIONAL STATE OF)", "BONAIRE, SINT EUSTATIUS AND SABA", "BOSNIA AND HERZEGOVINA", "BOTSWANA", "BOUVET ISLAND", "BRAZIL", "BRITISH INDIAN OCEAN TERRITORY (THE)", "BRUNEI DARUSSALAM", "BULGARIA", "BURKINA FASO", "BURUNDI", "CABO VERDE", "CAMBODIA", "CAMEROON", "CANADA", "CAYMAN ISLANDS (THE)", "CENTRAL AFRICAN REPUBLIC (THE)", "CHAD", "CHILE", "CHILE", "CHINA", "CHRISTMAS ISLAND", "COCOS (KEELING) ISLANDS (THE)", "COLOMBIA", "COLOMBIA", "COMOROS (THE)", "CONGO (THE DEMOCRATIC REPUBLIC OF THE)", "CONGO (THE)", "COOK ISLANDS (THE)", "COSTA RICA", "CROATIA", "CUBA", "CUBA", "CURAÇAO", "CYPRUS", "CZECH REPUBLIC (THE)", "CÔTE D'IVOIRE", "DENMARK", "DJIBOUTI", "DOMINICA", "DOMINICAN REPUBLIC (THE)", "ECUADOR", "EGYPT", "EL SALVADOR", "EL SALVADOR", "EQUATORIAL GUINEA", "ERITREA", "ESTONIA", "ETHIOPIA", "EUROPEAN UNION", "FALKLAND ISLANDS (THE) [MALVINAS]", "FAROE ISLANDS (THE)", "FIJI", "FINLAND", "FRANCE", "FRENCH GUIANA", "FRENCH POLYNESIA", "FRENCH SOUTHERN TERRITORIES (THE)", "GABON", "GAMBIA (THE)", "GEORGIA", "GERMANY", "GHANA", "GIBRALTAR", "GREECE", "GREENLAND", "GRENADA", "GUADELOUPE", "GUAM", "GUATEMALA", "GUERNSEY", "GUINEA", "GUINEA-BISSAU", "GUYANA", "HAITI", "HAITI", "HEARD ISLAND AND McDONALD ISLANDS", "HOLY SEE (THE)", "HONDURAS", "HONG KONG", "HUNGARY", "ICELAND", "INDIA", "INDONESIA", "INTERNATIONAL MONETARY FUND (IMF) ", "IRAN (ISLAMIC REPUBLIC OF)", "IRAQ", "IRELAND", "ISLE OF MAN", "ISRAEL", "ITALY", "JAMAICA", "JAPAN", "JERSEY", "JORDAN", "KAZAKHSTAN", "KENYA", "KIRIBATI", "KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)", "KOREA (THE REPUBLIC OF)", "KUWAIT", "KYRGYZSTAN", "LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)", "LATVIA", "LEBANON", "LESOTHO", "LESOTHO", "LIBERIA", "LIBYA", "LIECHTENSTEIN", "LITHUANIA", "LUXEMBOURG", "MACAO", "MADAGASCAR", "MALAWI", "MALAYSIA", "MALDIVES", "MALI", "MALTA", "MARSHALL ISLANDS (THE)", "MARTINIQUE", "MAURITANIA", "MAURITIUS", "MAYOTTE", "MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP", "MEXICO", "MEXICO", "MICRONESIA (FEDERATED STATES OF)", "MOLDOVA (THE REPUBLIC OF)", "MONACO", "MONGOLIA", "MONTENEGRO", "MONTSERRAT", "MOROCCO", "MOZAMBIQUE", "MYANMAR", "NAMIBIA", "NAMIBIA", "NAURU", "NEPAL", "NETHERLANDS (THE)", "NEW CALEDONIA", "NEW ZEALAND", "NICARAGUA", "NIGER (THE)", "NIGERIA", "NIUE", "NORFOLK ISLAND", "NORTHERN MARIANA ISLANDS (THE)", "NORWAY", "OMAN", "PAKISTAN", "PALAU", "PALESTINE, STATE OF", "PANAMA", "PANAMA", "PAPUA NEW GUINEA", "PARAGUAY", "PERU", "PHILIPPINES (THE)", "PITCAIRN", "POLAND", "PORTUGAL", "PUERTO RICO", "QATAR", "REPUBLIC OF NORTH MACEDONIA", "ROMANIA", "RUSSIAN FEDERATION (THE)", "RWANDA", "RÉUNION", "SAINT BARTHÉLEMY", "SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA", "SAINT KITTS AND NEVIS", "SAINT LUCIA", "SAINT MARTIN (FRENCH PART)", "SAINT PIERRE AND MIQUELON", "SAINT VINCENT AND THE GRENADINES", "SAMOA", "SAN MARINO", "SAO TOME AND PRINCIPE", "SAUDI ARABIA", "SENEGAL", "SERBIA", "SEYCHELLES", "SIERRA LEONE", "SINGAPORE", "SINT MAARTEN (DUTCH PART)", "SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS \"SUCRE\"", "SLOVAKIA", "SLOVENIA", "SOLOMON ISLANDS", "SOMALIA", "SOUTH AFRICA", "SOUTH GEORGIA AND THE SOUTH SANDWICH ISLANDS", "SOUTH SUDAN", "SPAIN", "SRI LANKA", "SUDAN (THE)", "SURINAME", "SVALBARD AND JAN MAYEN", "SWAZILAND", "SWEDEN", "SWITZERLAND", "SWITZERLAND", "SWITZERLAND", "SYRIAN ARAB REPUBLIC", "TAIWAN (PROVINCE OF CHINA)", "TAJIKISTAN", "TANZANIA, UNITED REPUBLIC OF", "THAILAND", "TIMOR-LESTE", "TOGO", "TOKELAU", "TONGA", "TRINIDAD AND TOBAGO", "TUNISIA", "TURKEY", "TURKMENISTAN", "TURKS AND CAICOS ISLANDS (THE)", "TUVALU", "UGANDA", "UKRAINE", "UNITED ARAB EMIRATES (THE)", "UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)", "UNITED STATES MINOR OUTLYING ISLANDS (THE)", "UNITED STATES OF AMERICA (THE)", "UNITED STATES OF AMERICA (THE)", "URUGUAY", "URUGUAY", "UZBEKISTAN", "VANUATU", "VENEZUELA (BOLIVARIAN REPUBLIC OF)", "VENEZUELA (BOLIVARIAN REPUBLIC OF)", "VIET NAM", "VIRGIN ISLANDS (BRITISH)", "VIRGIN ISLANDS (U.S.)", "WALLIS AND FUTUNA", "WESTERN SAHARA", "YEMEN", "ZAMBIA", "ZIMBABWE", "ÅLAND ISLANDS"}

	currencies = []string{
		"Afghani", "Lek", "Algerian Dinar", "US Dollar", "Euro", "Kwanza", "East Caribbean Dollar", "No universal currency", "East Caribbean Dollar", "Argentine Peso", "Armenian Dram", "Aruban Florin", "Australian Dollar", "Euro", "Azerbaijanian Manat", "Bahamian Dollar", "Bahraini Dinar", "Taka", "Barbados Dollar", "Belarussian Ruble", "Euro", "Belize Dollar", "CFA Franc BCEAO", "Bermudian Dollar", "Ngultrum", "Indian Rupee", "Boliviano", "Mvdol", "US Dollar", "Convertible Mark", "Pula", "Norwegian Krone", "Brazilian Real", "US Dollar", "Brunei Dollar", "Bulgarian Lev", "CFA Franc BCEAO", "Burundi Franc", "Cabo Verde Escudo", "Riel", "CFA Franc BEAC", "Canadian Dollar", "Cayman Islands Dollar", "CFA Franc BEAC", "CFA Franc BEAC", "Unidad de Fomento", "Chilean Peso", "Yuan Renminbi", "Australian Dollar", "Australian Dollar", "Colombian Peso", "Unidad de Valor Real", "Comoro Franc", "Congolese Franc", "CFA Franc BEAC", "New Zealand Dollar", "Costa Rican Colon", "Kuna", "Peso Convertible", "Cuban Peso", "Netherlands Antillean Guilder", "Euro", "Czech Koruna", "CFA Franc BCEAO", "Danish Krone", "Djibouti Franc", "East Caribbean Dollar", "Dominican Peso", "US Dollar", "Egyptian Pound", "El Salvador Colon", "US Dollar", "CFA Franc BEAC", "Nakfa", "Euro", "Ethiopian Birr", "Euro", "Falkland Islands Pound", "Danish Krone", "Fiji Dollar", "Euro", "Euro", "Euro", "CFP Franc", "Euro", "CFA Franc BEAC", "Dalasi", "Lari", "Euro", "Ghana Cedi", "Gibraltar Pound", "Euro", "Danish Krone", "East Caribbean Dollar", "Euro", "US Dollar", "Quetzal", "Pound Sterling", "Guinea Franc", "CFA Franc BCEAO", "Guyana Dollar", "Gourde", "US Dollar", "Australian Dollar", "Euro", "Lempira", "Hong Kong Dollar", "Forint", "Iceland Krona", "Indian Rupee", "Rupiah", "SDR (Special Drawing Right)", "Iranian Rial", "Iraqi Dinar", "Euro", "Pound Sterling", "New Israeli Sheqel", "Euro", "Jamaican Dollar", "Yen", "Pound Sterling", "Jordanian Dinar", "Tenge", "Kenyan Shilling", "Australian Dollar", "North Korean Won", "Won", "Kuwaiti Dinar", "Som", "Kip", "Euro", "Lebanese Pound", "Loti", "Rand", "Liberian Dollar", "Libyan Dinar", "Swiss Franc", "Euro", "Euro", "Pataca", "Malagasy Ariary", "Kwacha", "Malaysian Ringgit", "Rufiyaa", "CFA Franc BCEAO", "Euro", "US Dollar", "Euro", "Ouguiya", "Mauritius Rupee", "Euro", "ADB Unit of Account", "Mexican Peso", "Mexican Unidad de Inversion (UDI)", "US Dollar", "Moldovan Leu", "Euro", "Tugrik", "Euro", "East Caribbean Dollar", "Moroccan Dirham", "Mozambique Metical", "Kyat", "Namibia Dollar", "Rand", "Australian Dollar", "Nepalese Rupee", "Euro", "CFP Franc", "New Zealand Dollar", "Cordoba Oro", "CFA Franc BCEAO", "Naira", "New Zealand Dollar", "Australian Dollar", "US Dollar", "Norwegian Krone", "Rial Omani", "Pakistan Rupee", "US Dollar", "No universal currency", "Balboa", "US Dollar", "Kina", "Guarani", "Nuevo Sol", "Philippine Peso", "New Zealand Dollar", "Zloty", "Euro", "US Dollar", "Qatari Rial", "Denar", "Romanian Leu", "Russian Ruble", "Rwanda Franc", "Euro", "Euro", "Saint Helena Pound", "East Caribbean Dollar", "East Caribbean Dollar", "Euro", "Euro", "East Caribbean Dollar", "Tala", "Euro", "Dobra", "Saudi Riyal", "CFA Franc BCEAO", "Serbian Dinar", "Seychelles Rupee", "Leone", "Singapore Dollar", "Netherlands Antillean Guilder", "Sucre", "Euro", "Euro", "Solomon Islands Dollar", "Somali Shilling", "Rand", "No universal currency", "South Sudanese Pound", "Euro", "Sri Lanka Rupee", "Sudanese Pound", "Surinam Dollar", "Norwegian Krone", "Lilangeni", "Swedish Krona", "WIR Euro", "Swiss Franc", "WIR Franc", "Syrian Pound", "New Taiwan Dollar", "Somoni", "Tanzanian Shilling", "Baht", "US Dollar", "CFA Franc BCEAO", "New Zealand Dollar", "Pa’anga", "Trinidad and Tobago Dollar", "Tunisian Dinar", "Turkish Lira", "Turkmenistan New Manat", "US Dollar", "Australian Dollar", "Uganda Shilling", "Hryvnia", "UAE Dirham", "Pound Sterling", "US Dollar", "US Dollar", "US Dollar (Next day)", "Uruguay Peso en Unidades Indexadas (URUIURUI)", "Peso Uruguayo", "Uzbekistan Sum", "Vatu", "Bolivar", "Bolivar", "Dong", "US Dollar", "US Dollar", "CFP Franc", "Moroccan Dirham", "Yemeni Rial", "Zambian Kwacha", "Zimbabwe Dollar", "Euro"}

	currenciesCodes = []string{
		"AFN", "ALL", "DZD", "USD", "EUR", "AOA", "XCD", "", "XCD", "ARS", "AMD", "AWG", "AUD", "EUR", "AZN", "BSD", "BHD", "BDT", "BBD", "BYN", "EUR", "BZD", "XOF", "BMD", "BTN", "INR", "BOB", "BOV", "USD", "BAM", "BWP", "NOK", "BRL", "USD", "BND", "BGN", "XOF", "BIF", "CVE", "KHR", "XAF", "CAD", "KYD", "XAF", "XAF", "CLF", "CLP", "CNY", "AUD", "AUD", "COP", "COU", "KMF", "CDF", "XAF", "NZD", "CRC", "HRK", "CUC", "CUP", "ANG", "EUR", "CZK", "XOF", "DKK", "DJF", "XCD", "DOP", "USD", "EGP", "SVC", "USD", "XAF", "ERN", "EUR", "ETB", "EUR", "FKP", "DKK", "FJD", "EUR", "EUR", "EUR", "XPF", "EUR", "XAF", "GMD", "GEL", "EUR", "GHS", "GIP", "EUR", "DKK", "XCD", "EUR", "USD", "GTQ", "GBP", "GNF", "XOF", "GYD", "HTG", "USD", "AUD", "EUR", "HNL", "HKD", "HUF", "ISK", "INR", "IDR", "XDR", "IRR", "IQD", "EUR", "GBP", "ILS", "EUR", "JMD", "JPY", "GBP", "JOD", "KZT", "KES", "AUD", "KPW", "KRW", "KWD", "KGS", "LAK", "EUR", "LBP", "LSL", "ZAR", "LRD", "LYD", "CHF", "EUR", "EUR", "MOP", "MGA", "MWK", "MYR", "MVR", "XOF", "EUR", "USD", "EUR", "MRU", "MUR", "EUR", "XUA", "MXN", "MXV", "USD", "MDL", "EUR", "MNT", "EUR", "XCD", "MAD", "MZN", "MMK", "NAD", "ZAR", "AUD", "NPR", "EUR", "XPF", "NZD", "NIO", "XOF", "NGN", "NZD", "AUD", "USD", "NOK", "OMR", "PKR", "USD", "", "PAB", "USD", "PGK", "PYG", "PEN", "PHP", "NZD", "PLN", "EUR", "USD", "QAR", "MKD", "RON", "RUB", "RWF", "EUR", "EUR", "SHP", "XCD", "XCD", "EUR", "EUR", "XCD", "WST", "EUR", "STN", "SAR", "XOF", "RSD", "SCR", "SLL", "SGD", "ANG", "XSU", "EUR", "EUR", "SBD", "SOS", "ZAR", "", "SSP", "EUR", "LKR", "SDG", "SRD", "NOK", "SZL", "SEK", "CHE", "CHF", "CHW", "SYP", "TWD", "TJS", "TZS", "THB", "USD", "XOF", "NZD", "TOP", "TTD", "TND", "TRY", "TMT", "USD", "AUD", "UGX", "UAH", "AED", "GBP", "USD", "USD", "USN", "UYI", "UYU", "UZS", "VUV", "VEF", "VED", "VND", "USD", "USD", "XPF", "MAD", "YER", "ZMW", "ZWL", "EUR"}

	currenciesNumbers = []int{971, 8, 12, 840, 978, 973, 951, 0, 951, 32, 51, 533, 36, 978, 944, 44, 48, 50, 52, 933, 978, 84, 952, 60, 64, 356, 68, 984, 840, 977, 72, 578, 986, 840, 96, 975, 952, 108, 132, 116, 950, 124, 136, 950, 950, 990, 152, 156, 36, 36, 170, 970, 174, 976, 950, 554, 188, 191, 931, 192, 532, 978, 203, 952, 208, 262, 951, 214, 840, 818, 222, 840, 950, 232, 978, 230, 978, 238, 208, 242, 978, 978, 978, 953, 978, 950, 270, 981, 978, 936, 292, 978, 208, 951, 978, 840, 320, 826, 324, 952, 328, 332, 840, 36, 978, 340, 344, 348, 352, 356, 360, 960, 364, 368, 978, 826, 376, 978, 388, 392, 826, 400, 398, 404, 36, 408, 410, 414, 417, 418, 978, 422, 426, 710, 430, 434, 756, 978, 978, 446, 969, 454, 458, 462, 952, 978, 840, 978, 929, 480, 978, 965, 484, 979, 840, 498, 978, 496, 978, 951, 504, 943, 104, 516, 710, 36, 524, 978, 953, 554, 558, 952, 566, 554, 36, 840, 578, 512, 586, 840, 0, 590, 840, 598, 600, 604, 608, 554, 985, 978, 840, 634, 807, 946, 643, 646, 978, 978, 654, 951, 951, 978, 978, 951, 882, 978, 930, 682, 952, 941, 690, 694, 702, 532, 994, 978, 978, 90, 706, 710, 0, 728, 978, 144, 938, 968, 578, 748, 752, 947, 756, 948, 760, 901, 972, 834, 764, 840, 952, 554, 776, 780, 788, 949, 934, 840, 36, 800, 980, 784, 826, 840, 840, 997, 940, 858, 860, 548, 937, 926, 704, 840, 840, 953, 504, 886, 967, 932, 978}
)

// Currency is a faker struct for generating currency data
type Currency struct {
	Faker *Faker
}

// Currency returns a random currency name
func (c Currency) Currency() string {
	return c.Faker.RandomStringElement(currencies)
}

// Code returns a random currency code
func (c Currency) Code() string {
	return c.Faker.RandomStringElement(currenciesCodes)
}

// Code returns a random currency number
func (c Currency) Number() int {
	return c.Faker.RandomIntElement(currenciesNumbers)
}

// Country returns a random currency country
func (c Currency) Country() string {
	return c.Faker.RandomStringElement(currenciesCountries)
}

// CurrencyAndCode returns a random currency name and currency code
func (c Currency) CurrencyAndCode() (string, string) {
	index := c.Faker.IntBetween(0, len(currencies)-1)
	return currencies[index], currenciesCodes[index]
}
