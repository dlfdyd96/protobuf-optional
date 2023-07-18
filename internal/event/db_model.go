package event

type LocaleText struct {
	EN string  `pg:"EN",json:"EN"`
	KO *string `pg:"KO",json:"KO"`
	VI *string `pg:"VI",json:"VI"`
	TH *string `pg:"TH",json:"TH"`
	ID *string `pg:"ID",json:"ID"`
	ZH *string `pg:"ZH",json:"ZH"`
	TL *string `pg:"TL",json:"TL"`
	RU *string `pg:"RU",json:"RU"`
	KM *string `pg:"KM",json:"KM"`
	UZ *string `pg:"UZ",json:"UZ"`
	HI *string `pg:"HI",json:"HI"`
	NE *string `pg:"NE",json:"NE"`
	SI *string `pg:"SI",json:"SI"`
	BN *string `pg:"BN",json:"BN"`
	MY *string `pg:"MY",json:"MY"`
	MS *string `pg:"MS",json:"MS"`
	UR *string `pg:"UR",json:"UR"`
}

type Event struct {
	tableName struct{} `pg:"event,alias:t"`

	ID           int32       `pg:"id,pk"`
	EventDetail1 *LocaleText `pg:"event_detail_1,composite:locale_text"`
}
