// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import "database/sql"

// PostgreSQLCreateTablesReferences - создаёт таблицы для схемы references
func PostgreSQLCreateTablesReferences(dbc *sql.DB) {

	// Рецепты и список покупок

	var CreateStatements = NamedCreateStatements{
		NamedCreateStatement{
			TableName: "files",
			CreateStatement: `CREATE TABLE "references".files
			(
				id bigserial NOT NULL,
				filename character varying(255) COLLATE pg_catalog."default",
				filesize bigint,
				filetype character varying(50) COLLATE pg_catalog."default",
				file_id character varying(50) COLLATE pg_catalog."default",
				preview_id character varying(50) COLLATE pg_catalog."default",
				CONSTRAINT "Files_pkey" PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".files
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "countries",
			CreateStatement: `CREATE TABLE "references".countries
			(
				id bigserial NOT NULL,
				name character varying(60) COLLATE pg_catalog."default",
				full_name character varying(60) COLLATE pg_catalog."default",
				english character varying(100) COLLATE pg_catalog."default",
				alpha_2 character varying(2) COLLATE pg_catalog."default",
				alpha_3 character varying(3) COLLATE pg_catalog."default",
				iso character varying(3) COLLATE pg_catalog."default",
				location character varying(20) COLLATE pg_catalog."default",
				location_precise character varying(30) COLLATE pg_catalog."default",
				CONSTRAINT countries_pkey PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".countries
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "countries fill",
			CreateStatement: `INSERT INTO
			"references".countries(name, full_name, english, alpha_2, alpha_3, iso, location, location_precise) 
			VALUES
			('Абхазия', 'Республика Абхазия', 'Abkhazia', 'AB', 'ABH', '895', 'Азия', 'Закавказье'),
			('Австралия', '', 'Australia', 'AU', 'AUS', '036', 'Океания', 'Австралия и Новая Зеландия'),
			('Австрия', 'Австрийская Республика', 'Austria', 'AT', 'AUT', '040', 'Европа', 'Западная Европа'),
			('Азербайджан', 'Республика Азербайджан', 'Azerbaijan', 'AZ', 'AZE', '031', 'Азия', 'Западная Азия'),
			('Албания', 'Республика Албания', 'Albania', 'AL', 'ALB', '008', 'Европа', 'Южная Европа'),
			('Алжир', 'Алжирская Народная Демократическая Республика', 'Algeria', 'DZ', 'DZA', '012', 'Африка', 'Северная Африка'),
			('Американское Самоа', '', 'American Samoa', 'AS', 'ASM', '016', 'Океания', 'Полинезия'),
			('Ангилья', '', 'Anguilla', 'AI', 'AIA', '660', 'Америка', 'Карибский бассейн'),
			('Ангола', 'Республика Ангола', 'Angola', 'AO', 'AGO', '024', 'Африка', 'Центральная Африка'),
			('Андорра', 'Княжество Андорра', 'Andorra', 'AD', 'AND', '020', 'Европа', 'Южная Европа'),
			('Антарктида', '', 'Antarctica', 'AQ', 'ATA', '010', 'Антарктика', ' '),
			('Антигуа и Барбуда', '', 'Antigua and Barbuda', 'AG', 'ATG', '028', 'Америка', 'Карибский бассейн'),
			('Аргентина', 'Аргентинская Республика', 'Argentina', 'AR', 'ARG', '032', 'Америка', 'Южная Америка'),
			('Армения', 'Республика Армения', 'Armenia', 'AM', 'ARM', '051', 'Азия', 'Западная Азия'),
			('Аруба', '', 'Aruba', 'AW', 'ABW', '533', 'Америка', 'Карибский бассейн'),
			('Афганистан', 'Переходное Исламское Государство Афганистан', 'Afghanistan', 'AF', 'AFG', '004', 'Азия', 'Южная часть Центральной Азии'),
			('Багамы', 'Содружество Багамы', 'Bahamas', 'BS', 'BHS', '044', 'Америка', 'Карибский бассейн'),
			('Бангладеш', 'Народная Республика Бангладеш', 'Bangladesh', 'BD', 'BGD', '050', 'Азия', 'Южная часть Центральной Азии'),
			('Барбадос', '', 'Barbados', 'BB', 'BRB', '052', 'Америка', 'Карибский бассейн'),
			('Бахрейн', 'Королевство Бахрейн', 'Bahrain', 'BH', 'BHR', '048', 'Азия', 'Западная Азия'),
			('Беларусь', 'Республика Беларусь', 'Belarus', 'BY', 'BLR', '112', 'Европа', 'Восточная Европа'),
			('Белиз', '', 'Belize', 'BZ', 'BLZ', '084', 'Америка', 'Карибский бассейн'),
			('Бельгия', 'Королевство Бельгии', 'Belgium', 'BE', 'BEL', '056', 'Европа', 'Западная Европа'),
			('Бенин', 'Республика Бенин', 'Benin', 'BJ', 'BEN', '204', 'Африка', 'Западная Африка'),
			('Бермуды', '', 'Bermuda', 'BM', 'BMU', '060', 'Америка', 'Северная Америка'),
			('Болгария', 'Республика Болгария', 'Bulgaria', 'BG', 'BGR', '100', 'Европа', 'Восточная Европа'),
			('Боливия, Многонациональное Государство', 'Многонациональное Государство Боливия', 'Bolivia, plurinational state of', 'BO', 'BOL', '068', 'Америка', 'Южная Америка'),
			('Бонайре, Саба и Синт-Эстатиус', '', 'Bonaire, Sint Eustatius and Saba', 'BQ', 'BES', '535', 'Америка', 'Карибский бассейн'),
			('Босния и Герцеговина', '', 'Bosnia and Herzegovina', 'BA', 'BIH', '070', 'Европа', 'Южная Европа'),
			('Ботсвана', 'Республика Ботсвана', 'Botswana', 'BW', 'BWA', '072', 'Африка', 'Южная часть Африки'),
			('Бразилия', 'Федеративная Республика Бразилия', 'Brazil', 'BR', 'BRA', '076', 'Америка', 'Южная Америка'),
			('Британская территория в Индийском океане', '', 'British Indian Ocean Territory', 'IO', 'IOT', '086', 'Океания', 'Индийский океан'),
			('Бруней-Даруссалам', '', 'Brunei Darussalam', 'BN', 'BRN', '096', 'Азия', 'Юго-Восточная Азия'),
			('Буркина-Фасо', '', 'Burkina Faso', 'BF', 'BFA', '854', 'Африка', 'Западная Африка'),
			('Бурунди', 'Республика Бурунди', 'Burundi', 'BI', 'BDI', '108', 'Африка', 'Восточная Африка'),
			('Бутан', 'Королевство Бутан', 'Bhutan', 'BT', 'BTN', '064', 'Азия', 'Южная часть Центральной Азии'),
			('Вануату', 'Республика Вануату', 'Vanuatu', 'VU', 'VUT', '548', 'Океания', 'Меланезия'),
			('Венгрия', 'Венгерская Республика', 'Hungary', 'HU', 'HUN', '348', 'Европа', 'Восточная Европа'),
			('Венесуэла Боливарианская Республика', 'Боливарийская Республика Венесуэла', 'Venezuela', 'VE', 'VEN', '862', 'Америка', 'Южная Америка'),
			('Виргинские острова, Британские', 'Британские Виргинские острова', 'Virgin Islands, British', 'VG', 'VGB', '092', 'Америка', 'Карибский бассейн'),
			('Виргинские острова, США', 'Виргинские острова Соединенных Штатов', 'Virgin Islands, U.S.', 'VI', 'VIR', '850', 'Америка', 'Карибский бассейн'),
			('Вьетнам', 'Социалистическая Республика Вьетнам', 'Vietnam', 'VN', 'VNM', '704', 'Азия', 'Юго-Восточная Азия'),
			('Габон', 'Габонская Республика', 'Gabon', 'GA', 'GAB', '266', 'Африка', 'Центральная Африка'),
			('Гаити', 'Республика Гаити', 'Haiti', 'HT', 'HTI', '332', 'Америка', 'Карибский бассейн'),
			('Гайана', 'Республика Гайана', 'Guyana', 'GY', 'GUY', '328', 'Америка', 'Южная Америка'),
			('Гамбия', 'Республика Гамбия', 'Gambia', 'GM', 'GMB', '270', 'Африка', 'Западная Африка'),
			('Гана', 'Республика Гана', 'Ghana', 'GH', 'GHA', '288', 'Африка', 'Западная Африка'),
			('Гваделупа', '', 'Guadeloupe', 'GP', 'GLP', '312', 'Америка', 'Карибский бассейн'),
			('Гватемала', 'Республика Гватемала', 'Guatemala', 'GT', 'GTM', '320', 'Америка', 'Центральная Америка'),
			('Гвинея', 'Гвинейская Республика', 'Guinea', 'GN', 'GIN', '324', 'Африка', 'Западная Африка'),
			('Гвинея-Бисау', 'Республика Гвинея-Бисау', 'Guinea-Bissau', 'GW', 'GNB', '624', 'Африка', 'Западная Африка'),
			('Германия', 'Федеративная Республика Германия', 'Germany', 'DE', 'DEU', '276', 'Европа', 'Западная Европа'),
			('Гернси', '', 'Guernsey', 'GG', 'GGY', '831', 'Европа', 'Северная Европа'),
			('Гибралтар', '', 'Gibraltar', 'GI', 'GIB', '292', 'Европа', 'Южная Европа'),
			('Гондурас', 'Республика Гондурас', 'Honduras', 'HN', 'HND', '340', 'Америка', 'Центральная Америка'),
			('Гонконг', 'Специальный  административный  регион Китая Гонконг', 'Hong Kong', 'HK', 'HKG', '344', 'Азия', 'Восточная Азия'),
			('Гренада', '', 'Grenada', 'GD', 'GRD', '308', 'Америка', 'Карибский бассейн'),
			('Гренландия', '', 'Greenland', 'GL', 'GRL', '304', 'Америка', 'Северная Америка'),
			('Греция', 'Греческая Республика', 'Greece', 'GR', 'GRC', '300', 'Европа', 'Южная Европа'),
			('Грузия', '', 'Georgia', 'GE', 'GEO', '268', 'Азия', 'Западная Азия'),
			('Гуам', '', 'Guam', 'GU', 'GUM', '316', 'Океания', 'Микронезия'),
			('Дания', 'Королевство Дания', 'Denmark', 'DK', 'DNK', '208', 'Европа', 'Северная Европа'),
			('Джерси', '', 'Jersey', 'JE', 'JEY', '832', 'Европа', 'Северная Европа'),
			('Джибути', 'Республика Джибути', 'Djibouti', 'DJ', 'DJI', '262', 'Африка', 'Восточная Африка'),
			('Доминика', 'Содружество Доминики', 'Dominica', 'DM', 'DMA', '212', 'Америка', 'Карибский бассейн'),
			('Доминиканская Республика', '', 'Dominican Republic', 'DO', 'DOM', '214', 'Америка', 'Карибский бассейн'),
			('Египет', 'Арабская Республика Египет', 'Egypt', 'EG', 'EGY', '818', 'Африка', 'Северная Африка'),
			('Замбия', 'Республика Замбия', 'Zambia', 'ZM', 'ZMB', '894', 'Африка', 'Восточная Африка'),
			('Западная Сахара', '', 'Western Sahara', 'EH', 'ESH', '732', 'Африка', 'Северная Африка'),
			('Зимбабве', 'Республика Зимбабве', 'Zimbabwe', 'ZW', 'ZWE', '716', 'Африка', 'Восточная Африка'),
			('Израиль', 'Государство Израиль', 'Israel', 'IL', 'ISR', '376', 'Азия', 'Западная Азия'),
			('Индия', 'Республика Индия', 'India', 'IN', 'IND', '356', 'Азия', 'Южная часть Центральной Азии'),
			('Индонезия', 'Республика Индонезия', 'Indonesia', 'ID', 'IDN', '360', 'Азия', 'Юго-Восточная Азия'),
			('Иордания', 'Иорданское Хашимитское Королевство', 'Jordan', 'JO', 'JOR', '400', 'Азия', 'Западная Азия'),
			('Ирак', 'Республика Ирак', 'Iraq', 'IQ', 'IRQ', '368', 'Азия', 'Западная Азия'),
			('Иран, Исламская Республика', 'Исламская Республика Иран', 'Iran, Islamic Republic of', 'IR', 'IRN', '364', 'Азия', 'Южная часть Центральной Азии'),
			('Ирландия', '', 'Ireland', 'IE', 'IRL', '372', 'Европа', 'Северная Европа'),
			('Исландия', 'Республика Исландия', 'Iceland', 'IS', 'ISL', '352', 'Европа', 'Северная Европа'),
			('Испания', 'Королевство Испания', 'Spain', 'ES', 'ESP', '724', 'Европа', 'Южная Европа'),
			('Италия', 'Итальянская Республика', 'Italy', 'IT', 'ITA', '380', 'Европа', 'Южная Европа'),
			('Йемен', 'Йеменская Республика', 'Yemen', 'YE', 'YEM', '887', 'Азия', 'Западная Азия'),
			('Кабо-Верде', 'Республика Кабо-Верде', 'Cape Verde', 'CV', 'CPV', '132', 'Африка', 'Западная Африка'),
			('Казахстан', 'Республика Казахстан', 'Kazakhstan', 'KZ', 'KAZ', '398', 'Азия', 'Южная часть Центральной Азии'),
			('Камбоджа', 'Королевство Камбоджа', 'Cambodia', 'KH', 'KHM', '116', 'Азия', 'Юго-Восточная Азия'),
			('Камерун', 'Республика Камерун', 'Cameroon', 'CM', 'CMR', '120', 'Африка', 'Центральная Африка'),
			('Канада', '', 'Canada', 'CA', 'CAN', '124', 'Америка', 'Северная Америка'),
			('Катар', 'Государство Катар', 'Qatar', 'QA', 'QAT', '634', 'Азия', 'Западная Азия'),
			('Кения', 'Республика Кения', 'Kenya', 'KE', 'KEN', '404', 'Африка', 'Восточная Африка'),
			('Кипр', 'Республика Кипр', 'Cyprus', 'CY', 'CYP', '196', 'Азия', 'Западная Азия'),
			('Киргизия', 'Киргизская Республика', 'Kyrgyzstan', 'KG', 'KGZ', '417', 'Азия', 'Южная часть Центральной Азии'),
			('Кирибати', 'Республика Кирибати', 'Kiribati', 'KI', 'KIR', '296', 'Океания', 'Микронезия'),
			('Китай', 'Китайская Народная Республика', 'China', 'CN', 'CHN', '156', 'Азия', 'Восточная Азия'),
			('Кокосовые (Килинг) острова', '', 'Cocos (Keeling) Islands', 'CC', 'CCK', '166', 'Океания', 'Индийский океан'),
			('Колумбия', 'Республика Колумбия', 'Colombia', 'CO', 'COL', '170', 'Америка', 'Южная Америка'),
			('Коморы', 'Союз Коморы', 'Comoros', 'KM', 'COM', '174', 'Африка', 'Восточная Африка'),
			('Конго', 'Республика Конго', 'Congo', 'CG', 'COG', '178', 'Африка', 'Центральная Африка'),
			('Конго, Демократическая Республика', 'Демократическая Республика Конго', 'Congo, Democratic Republic of the', 'CD', 'COD', '180', 'Африка', 'Центральная Африка'),
			('Корея, Народно-Демократическая Республика', 'Корейская Народно-Демократическая Республика', 'Korea, Democratic Peoples republic of', 'KP', 'PRK', '408', 'Азия', 'Восточная Азия'),
			('Корея, Республика', 'Республика Корея', 'Korea, Republic of', 'KR', 'KOR', '410', 'Азия', 'Восточная Азия'),
			('Коста-Рика', 'Республика Коста-Рика', 'Costa Rica', 'CR', 'CRI', '188', 'Америка', 'Центральная Америка'),
			('Кот дИвуар', 'Республика Кот дИвуар', 'Cote dIvoire', 'CI', 'CIV', '384', 'Африка', 'Западная Африка'),
			('Куба', 'Республика Куба', 'Cuba', 'CU', 'CUB', '192', 'Америка', 'Карибский бассейн'),
			('Кувейт', 'Государство Кувейт', 'Kuwait', 'KW', 'KWT', '414', 'Азия', 'Западная Азия'),
			('Кюрасао', '', 'Cura?ao', 'CW', 'CUW', '531', 'Америка', 'Карибский бассейн'),
			('Лаос', 'Лаосская Народно-Демократическая Республика', 'Lao Peoples Democratic Republic', 'LA', 'LAO', '418', 'Азия', 'Юго-Восточная Азия'),
			('Латвия', 'Латвийская Республика', 'Latvia', 'LV', 'LVA', '428', 'Европа', 'Северная Европа'),
			('Лесото', 'Королевство Лесото', 'Lesotho', 'LS', 'LSO', '426', 'Африка', 'Южная часть Африки'),
			('Ливан', 'Ливанская Республика', 'Lebanon', 'LB', 'LBN', '422', 'Азия', 'Западная Азия'),
			('Ливийская Арабская Джамахирия', 'Социалистическая Народная Ливийская Арабская Джамахирия', 'Libyan Arab Jamahiriya', 'LY', 'LBY', '434', 'Африка', 'Северная Африка'),
			('Либерия', 'Республика Либерия', 'Liberia', 'LR', 'LBR', '430', 'Африка', 'Западная Африка'),
			('Лихтенштейн', 'Княжество Лихтенштейн', 'Liechtenstein', 'LI', 'LIE', '438', 'Европа', 'Западная Европа'),
			('Литва', 'Литовская Республика', 'Lithuania', 'LT', 'LTU', '440', 'Европа', 'Северная Европа'),
			('Люксембург', 'Великое Герцогство Люксембург', 'Luxembourg', 'LU', 'LUX', '442', 'Европа', 'Западная Европа'),
			('Маврикий', 'Республика Маврикий', 'Mauritius', 'MU', 'MUS', '480', 'Африка', 'Восточная Африка'),
			('Мавритания', 'Исламская Республика Мавритания', 'Mauritania', 'MR', 'MRT', '478', 'Африка', 'Западная Африка'),
			('Мадагаскар', 'Республика Мадагаскар', 'Madagascar', 'MG', 'MDG', '450', 'Африка', 'Восточная Африка'),
			('Майотта', '', 'Mayotte', 'YT', 'MYT', '175', 'Африка', 'Южная часть Африки'),
			('Макао', 'Специальный административный регион Китая Макао', 'Macao', 'MO', 'MAC', '446', 'Азия', 'Восточная Азия'),
			('Малави', 'Республика Малави', 'Malawi', 'MW', 'MWI', '454', 'Африка', 'Восточная Африка'),
			('Малайзия', '', 'Malaysia', 'MY', 'MYS', '458', 'Азия', 'Юго-Восточная Азия'),
			('Мали', 'Республика Мали', 'Mali', 'ML', 'MLI', '466', 'Африка', 'Западная Африка'),
			('Малые Тихоокеанские отдаленные острова Соединенных Штатов', '', 'United States Minor Outlying Islands', 'UM', 'UMI', '581', 'Океания', 'Индийский океан'),
			('Мальдивы', 'Мальдивская Республика', 'Maldives', 'MV', 'MDV', '462', 'Азия', 'Южная часть Центральной Азии'),
			('Мальта', 'Республика Мальта', 'Malta', 'MT', 'MLT', '470', 'Европа', 'Южная Европа'),
			('Марокко', 'Королевство Марокко', 'Morocco', 'MA', 'MAR', '504', 'Африка', 'Северная Африка'),
			('Мартиника', '', 'Martinique', 'MQ', 'MTQ', '474', 'Америка', 'Карибский бассейн'),
			('Маршалловы острова', 'Республика Маршалловы острова', 'Marshall Islands', 'MH', 'MHL', '584', 'Океания', 'Микронезия'),
			('Мексика', 'Мексиканские Соединенные Штаты', 'Mexico', 'MX', 'MEX', '484', 'Америка', 'Центральная Америка'),
			('Микронезия, Федеративные Штаты', 'Федеративные штаты Микронезии', 'Micronesia, Federated States of', 'FM', 'FSM', '583', 'Океания', 'Микронезия'),
			('Мозамбик', 'Республика Мозамбик', 'Mozambique', 'MZ', 'MOZ', '508', 'Африка', 'Восточная Африка'),
			('Молдова, Республика', 'Республика Молдова', 'Moldova', 'MD', 'MDA', '498', 'Европа', 'Восточная Европа'),
			('Монако', 'Княжество Монако', 'Monaco', 'MC', 'MCO', '492', 'Европа', 'Западная Европа'),
			('Монголия', '', 'Mongolia', 'MN', 'MNG', '496', 'Азия', 'Восточная Азия'),
			('Монтсеррат', '', 'Montserrat', 'MS', 'MSR', '500', 'Америка', 'Карибский бассейн'),
			('Мьянма', 'Союз Мьянма', 'Burma', 'MM', 'MMR', '104', 'Азия', 'Юго-Восточная Азия'),
			('Намибия', 'Республика Намибия', 'Namibia', 'NA', 'NAM', '516', 'Африка', 'Южная часть Африки'),
			('Науру', 'Республика Науру', 'Nauru', 'NR', 'NRU', '520', 'Океания', 'Микронезия'),
			('Непал', 'Королевство Непал', 'Nepal', 'NP', 'NPL', '524', 'Азия', 'Южная часть Центральной Азии'),
			('Нигер', 'Республика Нигер', 'Niger', 'NE', 'NER', '562', 'Африка', 'Западная Африка'),
			('Нигерия', 'Федеративная Республика Нигерия', 'Nigeria', 'NG', 'NGA', '566', 'Африка', 'Западная Африка'),
			('Нидерланды', 'Королевство Нидерландов', 'Netherlands', 'NL', 'NLD', '528', 'Европа', 'Западная Европа'),
			('Никарагуа', 'Республика Никарагуа', 'Nicaragua', 'NI', 'NIC', '558', 'Америка', 'Центральная Америка'),
			('Ниуэ', 'Республика Ниуэ', 'Niue', 'NU', 'NIU', '570', 'Океания', 'Полинезия'),
			('Новая Зеландия', '', 'New Zealand', 'NZ', 'NZL', '554', 'Океания', 'Австралия и Новая Зеландия'),
			('Новая Каледония', '', 'New Caledonia', 'NC', 'NCL', '540', 'Океания', 'Меланезия'),
			('Норвегия', 'Королевство Норвегия', 'Norway', 'NO', 'NOR', '578', 'Европа', 'Северная Европа'),
			('Объединенные Арабские Эмираты', '', 'United Arab Emirates', 'AE', 'ARE', '784', 'Азия', 'Западная Азия'),
			('Оман', 'Султанат Оман', 'Oman', 'OM', 'OMN', '512', 'Азия', 'Западная Азия'),
			('Остров Буве', '', 'Bouvet Island', 'BV', 'BVT', '074', '', 'Южный океан'),
			('Остров Мэн', '', 'Isle of Man', 'IM', 'IMN', '833', 'Европа', 'Северная Европа'),
			('Остров Норфолк', '', 'Norfolk Island', 'NF', 'NFK', '574', 'Океания', 'Австралия и Новая Зеландия'),
			('Остров Рождества', '', 'Christmas Island', 'CX', 'CXR', '162', 'Азия', 'Индийский океан'),
			('Остров Херд и острова Макдональд', '', 'Heard Island and McDonald Islands', 'HM', 'HMD', '334', '', 'Индийский океан'),
			('Острова Кайман', '', 'Cayman Islands', 'KY', 'CYM', '136', 'Америка', 'Карибский бассейн'),
			('Острова Кука', '', 'Cook Islands', 'CK', 'COK', '184', 'Океания', 'Полинезия'),
			('Острова Теркс и Кайкос', '', 'Turks and Caicos Islands', 'TC', 'TCA', '796', 'Америка', 'Карибский бассейн'),
			('Пакистан', 'Исламская Республика Пакистан', 'Pakistan', 'PK', 'PAK', '586', 'Азия', 'Южная часть Центральной Азии'),
			('Палау', 'Республика Палау', 'Palau', 'PW', 'PLW', '585', 'Океания', 'Микронезия'),
			('Палестинская территория, оккупированная', 'Оккупированная Палестинская территория', 'Palestinian Territory, Occupied', 'PS', 'PSE', '275', 'Азия', 'Западная Азия'),
			('Панама', 'Республика Панама', 'Panama', 'PA', 'PAN', '591', 'Америка', 'Центральная Америка'),
			('Папский Престол (Государство &mdash; город Ватикан)', '', 'Holy See (Vatican City State)', 'VA', 'VAT', '336', 'Европа', 'Южная Европа'),
			('Папуа-Новая Гвинея', '', 'Papua New Guinea', 'PG', 'PNG', '598', 'Океания', 'Меланезия'),
			('Парагвай', 'Республика Парагвай', 'Paraguay', 'PY', 'PRY', '600', 'Америка', 'Южная Америка'),
			('Перу', 'Республика Перу', 'Peru', 'PE', 'PER', '604', 'Америка', 'Южная Америка'),
			('Питкерн', '', 'Pitcairn', 'PN', 'PCN', '612', 'Океания', 'Полинезия'),
			('Польша', 'Республика Польша', 'Poland', 'PL', 'POL', '616', 'Европа', 'Восточная Европа'),
			('Португалия', 'Португальская Республика', 'Portugal', 'PT', 'PRT', '620', 'Европа', 'Южная Европа'),
			('Пуэрто-Рико', '', 'Puerto Rico', 'PR', 'PRI', '630', 'Америка', 'Карибский бассейн'),
			('Республика Македония', '', 'Macedonia, The Former Yugoslav Republic Of', 'MK', 'MKD', '807', 'Европа', 'Южная Европа'),
			('Реюньон', '', 'Reunion', 'RE', 'REU', '638', 'Африка', 'Восточная Африка'),
			('Россия', 'Российская Федерация', 'Russian Federation', 'RU', 'RUS', '643', 'Европа', 'Восточная Европа'),
			('Руанда', 'Руандийская Республика', 'Rwanda', 'RW', 'RWA', '646', 'Африка', 'Восточная Африка'),
			('Румыния', '', 'Romania', 'RO', 'ROU', '642', 'Европа', 'Восточная Европа'),
			('Самоа', 'Независимое Государство Самоа', 'Samoa', 'WS', 'WSM', '882', 'Океания', 'Полинезия'),
			('Сан-Марино', 'Республика Сан-Марино', 'San Marino', 'SM', 'SMR', '674', 'Европа', 'Южная Европа'),
			('Сан-Томе и Принсипи', 'Демократическая Республика Сан-Томе и Принсипи', 'Sao Tome and Principe', 'ST', 'STP', '678', 'Африка', 'Центральная Африка'),
			('Саудовская Аравия', 'Королевство Саудовская Аравия', 'Saudi Arabia', 'SA', 'SAU', '682', 'Азия', 'Западная Азия'),
			('Святая Елена, Остров вознесения, Тристан-да-Кунья', '', 'Saint Helena, Ascension And Tristan Da Cunha', 'SH', 'SHN', '654', 'Африка', 'Западная Африка'),
			('Северные Марианские острова', 'Содружество Северных Марианских островов', 'Northern Mariana Islands', 'MP', 'MNP', '580', 'Океания', 'Микронезия'),
			('Сен-Бартельми', '', 'Saint Barth?lemy', 'BL', 'BLM', '652', 'Америка', 'Карибский бассейн'),
			('Сен-Мартен', '', 'Saint Martin (French Part)', 'MF', 'MAF', '663', 'Америка', 'Карибский бассейн'),
			('Сенегал', 'Республика Сенегал', 'Senegal', 'SN', 'SEN', '686', 'Африка', 'Западная Африка'),
			('Сент-Винсент и Гренадины', '', 'Saint Vincent and the Grenadines', 'VC', 'VCT', '670', 'Америка', 'Карибский бассейн'),
			('Сент-Китс и Невис', '', 'Saint Kitts and Nevis', 'KN', 'KNA', '659', 'Америка', 'Карибский бассейн'),
			('Сент-Люсия', '', 'Saint Lucia', 'LC', 'LCA', '662', 'Америка', 'Карибский бассейн'),
			('Сент-Пьер и Микелон', '', 'Saint Pierre and Miquelon', 'PM', 'SPM', '666', 'Америка', 'Северная Америка'),
			('Сербия', 'Республика Сербия', 'Serbia', 'RS', 'SRB', '688', 'Европа', 'Южная Европа'),
			('Сейшелы', 'Республика Сейшелы', 'Seychelles', 'SC', 'SYC', '690', 'Африка', 'Восточная Африка'),
			('Сингапур', 'Республика Сингапур', 'Singapore', 'SG', 'SGP', '702', 'Азия', 'Юго-Восточная Азия'),
			('Синт-Мартен', '', 'Sint Maarten', 'SX', 'SXM', '534', 'Америка', 'Карибский бассейн'),
			('Сирийская Арабская Республика', '', 'Syrian Arab Republic', 'SY', 'SYR', '760', 'Азия', 'Западная Азия'),
			('Словакия', 'Словацкая Республика', 'Slovakia', 'SK', 'SVK', '703', 'Европа', 'Восточная Европа'),
			('Словения', 'Республика Словения', 'Slovenia', 'SI', 'SVN', '705', 'Европа', 'Южная Европа'),
			('Соединенное Королевство', 'Соединенное Королевство Великобритании и Северной Ирландии', 'United Kingdom', 'GB', 'GBR', '826', 'Европа', 'Северная Европа'),
			('Соединенные Штаты', 'Соединенные Штаты Америки', 'United States', 'US', 'USA', '840', 'Америка', 'Северная Америка'),
			('Соломоновы острова', '', 'Solomon Islands', 'SB', 'SLB', '090', 'Океания', 'Меланезия'),
			('Сомали', 'Сомалийская Республика', 'Somalia', 'SO', 'SOM', '706', 'Африка', 'Восточная Африка'),
			('Судан', 'Республика Судан', 'Sudan', 'SD', 'SDN', '729', 'Африка', 'Северная Африка'),
			('Суринам', 'Республика Суринам', 'Suriname', 'SR', 'SUR', '740', 'Америка', 'Южная Америка'),
			('Сьерра-Леоне', 'Республика Сьерра-Леоне', 'Sierra Leone', 'SL', 'SLE', '694', 'Африка', 'Западная Африка'),
			('Таджикистан', 'Республика Таджикистан', 'Tajikistan', 'TJ', 'TJK', '762', 'Азия', 'Южная часть Центральной Азии'),
			('Таиланд', 'Королевство Таиланд', 'Thailand', 'TH', 'THA', '764', 'Азия', 'Юго-Восточная Азия'),
			('Тайвань (Китай)', '', 'Taiwan, Province of China', 'TW', 'TWN', '158', 'Азия', 'Восточная Азия'),
			('Танзания, Объединенная Республика', 'Объединенная Республика Танзания', 'Tanzania, United Republic Of', 'TZ', 'TZA', '834', 'Африка', 'Восточная Африка'),
			('Тимор-Лесте', 'Демократическая Республика Тимор-Лесте', 'Timor-Leste', 'TL', 'TLS', '626', 'Азия', 'Юго-Восточная Азия'),
			('Того', 'Тоголезская Республика', 'Togo', 'TG', 'TGO', '768', 'Африка', 'Западная Африка'),
			('Токелау', '', 'Tokelau', 'TK', 'TKL', '772', 'Океания', 'Полинезия'),
			('Тонга', 'Королевство Тонга', 'Tonga', 'TO', 'TON', '776', 'Океания', 'Полинезия'),
			('Тринидад и Тобаго', 'Республика Тринидад и Тобаго', 'Trinidad and Tobago', 'TT', 'TTO', '780', 'Америка', 'Карибский бассейн'),
			('Тувалу', '', 'Tuvalu', 'TV', 'TUV', '798', 'Океания', 'Полинезия'),
			('Тунис', 'Тунисская Республика', 'Tunisia', 'TN', 'TUN', '788', 'Африка', 'Северная Африка'),
			('Туркмения', 'Туркменистан', 'Turkmenistan', 'TM', 'TKM', '795', 'Азия', 'Южная часть Центральной Азии'),
			('Турция', 'Турецкая Республика', 'Turkey', 'TR', 'TUR', '792', 'Азия', 'Западная Азия'),
			('Уганда', 'Республика Уганда', 'Uganda', 'UG', 'UGA', '800', 'Африка', 'Восточная Африка'),
			('Узбекистан', 'Республика Узбекистан', 'Uzbekistan', 'UZ', 'UZB', '860', 'Азия', 'Южная часть Центральной Азии'),
			('Украина', '', 'Ukraine', 'UA', 'UKR', '804', 'Европа', 'Восточная Европа'),
			('Уоллис и Футуна', '', 'Wallis and Futuna', 'WF', 'WLF', '876', 'Океания', 'Полинезия'),
			('Уругвай', 'Восточная Республика Уругвай', 'Uruguay', 'UY', 'URY', '858', 'Америка', 'Южная Америка'),
			('Фарерские острова', '', 'Faroe Islands', 'FO', 'FRO', '234', 'Европа', 'Северная Европа'),
			('Фиджи', 'Республика островов Фиджи', 'Fiji', 'FJ', 'FJI', '242', 'Океания', 'Меланезия'),
			('Филиппины', 'Республика Филиппины', 'Philippines', 'PH', 'PHL', '608', 'Азия', 'Юго-Восточная Азия'),
			('Финляндия', 'Финляндская Республика', 'Finland', 'FI', 'FIN', '246', 'Европа', 'Северная Европа'),
			('Фолклендские острова (Мальвинские)', '', 'Falkland Islands (Malvinas)', 'FK', 'FLK', '238', 'Америка', 'Южная Америка'),
			('Франция', 'Французская Республика', 'France', 'FR', 'FRA', '250', 'Европа', 'Западная Европа'),
			('Французская Гвиана', '', 'French Guiana', 'GF', 'GUF', '254', 'Америка', 'Южная Америка'),
			('Французская Полинезия', '', 'French Polynesia', 'PF', 'PYF', '258', 'Океания', 'Полинезия'),
			('Французские Южные территории', '', 'French Southern Territories', 'TF', 'ATF', '260', '', 'Индийский океан'),
			('Хорватия', 'Республика Хорватия', 'Croatia', 'HR', 'HRV', '191', 'Европа', 'Южная Европа'),
			('Центрально-Африканская Республика', '', 'Central African Republic', 'CF', 'CAF', '140', 'Африка', 'Центральная Африка'),
			('Чад', 'Республика Чад', 'Chad', 'TD', 'TCD', '148', 'Африка', 'Центральная Африка'),
			('Черногория', 'Республика Черногория', 'Montenegro', 'ME', 'MNE', '499', 'Европа', 'Южная Европа'),
			('Чешская Республика', '', 'Czech Republic', 'CZ', 'CZE', '203', 'Европа', 'Восточная Европа'),
			('Чили', 'Республика Чили', 'Chile', 'CL', 'CHL', '152', 'Америка', 'Южная Америка'),
			('Швейцария', 'Швейцарская Конфедерация', 'Switzerland', 'CH', 'CHE', '756', 'Европа', 'Западная Европа'),
			('Швеция', 'Королевство Швеция', 'Sweden', 'SE', 'SWE', '752', 'Европа', 'Северная Европа'),
			('Шпицберген и Ян Майен', '', 'Svalbard and Jan Mayen', 'SJ', 'SJM', '744', 'Европа', 'Северная Европа'),
			('Шри-Ланка', 'Демократическая Социалистическая Республика Шри-Ланка', 'Sri Lanka', 'LK', 'LKA', '144', 'Азия', 'Южная часть Центральной Азии'),
			('Эквадор', 'Республика Эквадор', 'Ecuador', 'EC', 'ECU', '218', 'Америка', 'Южная Америка'),
			('Экваториальная Гвинея', 'Республика Экваториальная Гвинея', 'Equatorial Guinea', 'GQ', 'GNQ', '226', 'Африка', 'Центральная Африка'),
			('Эландские острова', '', '?land Islands', 'AX', 'ALA', '248', 'Европа', 'Северная Европа'),
			('Эль-Сальвадор', 'Республика Эль-Сальвадор', 'El Salvador', 'SV', 'SLV', '222', 'Америка', 'Центральная Америка'),
			('Эритрея', '', 'Eritrea', 'ER', 'ERI', '232', 'Африка', 'Восточная Африка'),
			('Эсватини', 'Королевство Эсватини', 'Eswatini', 'SZ', 'SWZ', '748', 'Африка', 'Южная часть Африки'),
			('Эстония', 'Эстонская Республика', 'Estonia', 'EE', 'EST', '233', 'Европа', 'Северная Европа'),
			('Эфиопия', 'Федеративная Демократическая Республика Эфиопия', 'Ethiopia', 'ET', 'ETH', '231', 'Африка', 'Восточная Африка'),
			('Южная Африка', 'Южно-Африканская Республика', 'South Africa', 'ZA', 'ZAF', '710', 'Африка', 'Южная часть Африки'),
			('Южная Джорджия и Южные Сандвичевы острова', '', 'South Georgia and the South Sandwich Islands', 'GS', 'SGS', '239', '', 'Южный океан'),
			('Южная Осетия', 'Республика Южная Осетия', 'South Ossetia', 'OS', 'OST', '896', 'Азия', 'Закавказье'),
			('Южный Судан', '', 'South Sudan', 'SS', 'SSD', '728', 'Африка', 'Северная Африка'),
			('Ямайка', '', 'Jamaica', 'JM', 'JAM', '388', 'Америка', 'Карибский бассейн'),
			('Япония', '', 'Japan', 'JP', 'JPN', '392', 'Азия', 'Восточная Азия'); `,
		},
		NamedCreateStatement{
			TableName: "addresses",
			CreateStatement: `CREATE TABLE "references".addresses
			(
				id bigserial NOT NULL,
				index character varying(6) COLLATE pg_catalog."default",
				country_id bigint,
				city character varying(100) COLLATE pg_catalog."default",
				district character varying(100) COLLATE pg_catalog."default",
				street character varying(100) COLLATE pg_catalog."default",
				name character varying(200) COLLATE pg_catalog."default",
				CONSTRAINT addresses_pkey PRIMARY KEY (id),
				CONSTRAINT addresses_country_id_fkey FOREIGN KEY (country_id)
					REFERENCES "references".countries (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".addresses
				OWNER to postgres;
			
			CREATE INDEX fki_addresses_country_id_fkey
				ON "references".addresses USING btree
				(country_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "authors",
			CreateStatement: `CREATE TABLE "references".authors
			(
				id bigserial NOT NULL,
				first_name character varying(200) COLLATE pg_catalog."default",
				middle_name character varying(200) COLLATE pg_catalog."default",
				last_name character varying(200) COLLATE pg_catalog."default",
				bio text COLLATE pg_catalog."default",
				file_id bigint,
				country_id bigint,
				city character varying(100) COLLATE pg_catalog."default",
				eng_name character varying(200) COLLATE pg_catalog."default",
				user_id uuid NOT NULL,
				CONSTRAINT authors_pkey PRIMARY KEY (id),
				CONSTRAINT authors_country_id_fkey FOREIGN KEY (country_id)
					REFERENCES "references".countries (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT authors_file_id_fkey FOREIGN KEY (file_id)
					REFERENCES "references".files (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT,
				CONSTRAINT authors_user_id_fkey FOREIGN KEY (user_id)
					REFERENCES secret.users (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".authors
				OWNER to postgres;
			
			CREATE INDEX fki_authors_country_id_fkey
				ON "references".authors USING btree
				(country_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_authors_file_id_fkey
				ON "references".authors USING btree
				(file_id ASC NULLS LAST)
				TABLESPACE pg_default;
			
			CREATE INDEX fki_authors_user_id_fkey
				ON "references".authors USING btree
				(user_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
		NamedCreateStatement{
			TableName: "artwork_types",
			CreateStatement: `CREATE TABLE "references".artwork_types
			(
				id bigserial NOT NULL,
				name character varying(50) COLLATE pg_catalog."default",
				eng_name character varying(50) COLLATE pg_catalog."default",
				CONSTRAINT artwork_types_pkey PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".artwork_types
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "currencies",
			CreateStatement: `CREATE TABLE "references".currencies
			(
				id bigserial NOT NULL,
				rus_name character varying(50) COLLATE pg_catalog."default",
				eng_name character varying(50) COLLATE pg_catalog."default",
				iso_lat_3 character varying(3) COLLATE pg_catalog."default",
				iso_dig character varying(3) COLLATE pg_catalog."default",
				symbol character varying(1) COLLATE pg_catalog."default",
				CONSTRAINT currencies_pkey PRIMARY KEY (id)
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".currencies
				OWNER to postgres;`,
		},
		NamedCreateStatement{
			TableName: "currencies fill",
			CreateStatement: `INSERT INTO 
			"references".currencies(rus_name, eng_name, iso_lat_3, iso_dig, symbol)
			VALUES
			('Russian Ruble', 'РОССИЙСКИЙ РУБЛЬ', 'RUB', '643', '₽'),
			('Euro', 'ЕВРО', 'EUR', '978', '€'),
			('US Dollar', 'ДОЛЛАР США', 'USD', '840', '$'),
			('Pound Sterling', 'ФУНТ СТЕРЛИНГОВ', 'GBP', '826', '£'),
			('Yen', 'ЙЕНА', 'JPY', '392', '¥'),
			('Yuan Renminbi', 'КИТАЙСКИЙ ЮАНЬ', 'CNY', '156', '¥');`,
		},
		NamedCreateStatement{
			TableName: "terms",
			CreateStatement: `CREATE TABLE "references".terms
			(
				id bigserial NOT NULL,
				delivery_time text COLLATE pg_catalog."default",
				returns text COLLATE pg_catalog."default",
				delivery_cost text COLLATE pg_catalog."default",
				name character varying(100) COLLATE pg_catalog."default",
				currency_id bigint NOT NULL,
				CONSTRAINT terms_pkey PRIMARY KEY (id),
				CONSTRAINT terms_currency_id_fkey FOREIGN KEY (currency_id)
					REFERENCES "references".currencies (id) MATCH FULL
					ON UPDATE RESTRICT
					ON DELETE RESTRICT
			)
			
			TABLESPACE pg_default;
			
			ALTER TABLE "references".terms
				OWNER to postgres;
			
			CREATE INDEX fki_terms_currency_id_fkey
				ON "references".terms USING btree
				(currency_id ASC NULLS LAST)
				TABLESPACE pg_default;`,
		},
	}

	for _, ncs := range CreateStatements {
		PostgreSQLExecuteCreateStatement(dbc, ncs)
	}

}
