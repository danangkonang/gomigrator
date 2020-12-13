package seed

import (
	"fmt"
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Kecamatan7to9() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO kecamatan (kecamatan,kabupaten_id,provinsi_id)VALUES
	('ABIANSEMAL', 101, 7),
	('KUTA', 101, 7),
	('KUTA SELATAN', 101, 7),
	('KUTA UTARA', 101, 7),
	('MENGWI', 101, 7),
	('PETANG', 101, 7),
	('BANGLI', 102, 7),
	('KINTAMANI', 102, 7),
	('SUSUT', 102, 7),
	('TEMBUKU', 102, 7),
	('BANJAR', 103, 7),
	('BULELENG', 103, 7),
	('BUSUNGBIU', 103, 7),
	('GEROKGAK', 103, 7),
	('KUBUTAMBAHAN', 103, 7),
	('SAWAN', 103, 7),
	('SERIRIT', 103, 7),
	('SUKASADA', 103, 7),
	('TEJAKULA', 103, 7),
	('DENPASAR BARAT', 104, 7),
	('DENPASAR SELATAN', 104, 7),
	('DENPASAR TIMUR', 104, 7),
	('DENPASAR UTARA', 104, 7),
	('BELAH BATUH (BLAHBATUH)', 105, 7),
	('GIANYAR', 105, 7),
	('PAYANGAN', 105, 7),
	('SUKAWATI', 105, 7),
	('TAMPAK SIRING', 105, 7),
	('TEGALLALANG', 105, 7),
	('UBUD', 105, 7),
	('JEMBRANA', 106, 7),
	('MELAYA', 106, 7),
	('MENDOYO', 106, 7),
	('NEGARA', 106, 7),
	('PEKUTATAN', 106, 7),
	('ABANG', 107, 7),
	('BEBANDEM', 107, 7),
	('KARANG ASEM', 107, 7),
	('MANGGIS', 107, 7),
	('RENDANG', 107, 7),
	('SELAT', 107, 7),
	('SIDEMEN', 107, 7),
	('BANJARANGKAN', 108, 7),
	('DAWAN', 108, 7),
	('KLUNGKUNG', 108, 7),
	('NUSAPENIDA', 108, 7),
	('BATURITI', 109, 7),
	('KERAMBITAN', 109, 7),
	('MARGA', 109, 7),
	('PENEBEL', 109, 7),
	('PUPUAN', 109, 7),
	('SELEMADEG', 109, 7),
	('SELEMADEG / SALAMADEG TIMUR', 109, 7),
	('SELEMADEG / SALEMADEG BARAT', 109, 7),
	('TABANAN', 109, 7),
	('ARONGAN LAMBALEK', 110, 8),
	('BUBON', 110, 8),
	('JOHAN PAHLAWAN', 110, 8),
	('KAWAY XVI', 110, 8),
	('MEUREUBO', 110, 8),
	('PANTE CEUREUMEN (PANTAI CEUREMEN)', 110, 8),
	('PANTON REU', 110, 8),
	('SAMATIGA', 110, 8),
	('SUNGAI MAS', 110, 8),
	('WOYLA', 110, 8),
	('WOYLA BARAT', 110, 8),
	('WOYLA TIMUR', 110, 8),
	('BABAH ROT', 111, 8),
	('BLANG PIDIE', 111, 8),
	('KUALA BATEE', 111, 8),
	('LEMBAH SABIL', 111, 8),
	('MANGGENG', 111, 8),
	('SETIA', 111, 8),
	('SUSOH', 111, 8),
	('TANGAN-TANGAN', 111, 8),
	('BAITUSSALAM', 112, 8),
	('BLANK BINTANG', 112, 8),
	('DARUL IMARAH', 112, 8),
	('DARUL KAMAL', 112, 8),
	('DARUSSALAM', 112, 8),
	('INDRAPURI', 112, 8),
	('INGIN JAYA', 112, 8),
	('KOTA COT GLIE (KUTA COT GLIE)', 112, 8),
	('KOTA JANTHO', 112, 8),
	('KOTA MALAKA (KUTA MALAKA)', 112, 8),
	('KRUENG BARONA JAYA', 112, 8),
	('KUTA BARO', 112, 8),
	('LEMBAH SEULAWAH', 112, 8),
	('LEUPUNG', 112, 8),
	('LHOKNGA (LHO\\NGA)', 112, 8),
	('LHOONG', 112, 8),
	('MANTASIEK (MONTASIK)', 112, 8),
	('MESJID RAYA', 112, 8),
	('PEUKAN BADA', 112, 8),
	('PULO ACEH', 112, 8),
	('SEULIMEUM', 112, 8),
	('SUKA MAKMUR', 112, 8),
	('DARUL HIKMAH', 113, 8),
	('INDRA JAYA', 113, 8),
	('JAYA', 113, 8),
	('KEUDE PANGA', 113, 8),
	('KRUENG SABEE', 113, 8),
	('PASIE RAYA', 113, 8),
	('SAMPOINIET', 113, 8),
	('SETIA BAKTI', 113, 8),
	('TEUNOM', 113, 8),
	('BAKONGAN', 114, 8),
	('BAKONGAN TIMUR', 114, 8),
	('KLUET SELATAN', 114, 8),
	('KLUET TENGAH', 114, 8),
	('KLUET TIMUR', 114, 8),
	('KLUET UTARA', 114, 8),
	('KOTA BAHAGIA', 114, 8),
	('LABUHAN HAJI BARAT', 114, 8),
	('LABUHAN HAJI TIMUR', 114, 8),
	('MEUKEK', 114, 8),
	('PASIE RAJA', 114, 8),
	('SAMA DUA', 114, 8),
	('TAPAK TUAN', 114, 8),
	('TRUMON', 114, 8),
	('TRUMON TENGAH', 114, 8),
	('TRUMON TIMUR', 114, 8),
	('DANAU PARIS', 115, 8),
	('GUNUNG MERIAH (MARIAH)', 115, 8),
	('KOTA BAHARU', 115, 8),
	('KUALA BARU', 115, 8),
	('PULAU BANYAK', 115, 8),
	('PULAU BANYAK BARAT', 115, 8),
	('SINGKIL UTARA', 115, 8),
	('SINGKOHOR', 115, 8),
	('SURO MAKMUR', 115, 8),
	('BANDA MULIA', 116, 8),
	('BANDAR PUSAKA', 116, 8),
	('BENDAHARA', 116, 8),
	('KARANG BARU', 116, 8),
	('KEJURUAN MUDA', 116, 8),
	('KOTA KUALA SIMPANG', 116, 8),
	('MANYAK PAYED', 116, 8),
	('RANTAU', 116, 8),
	('SEKERAK', 116, 8),
	('SERUWAY', 116, 8),
	('TAMIANG HULU', 116, 8),
	('TENGGULUN', 116, 8),
	('ATU LINTANG', 117, 8),
	('BEBESEN', 117, 8),
	('BIES', 117, 8),
	('BINTANG', 117, 8),
	('CELALA', 117, 8),
	('JAGONG JEGET', 117, 8),
	('KEBAYAKAN', 117, 8),
	('KETOL', 117, 8),
	('KUTE PANANG', 117, 8),
	('LINGE', 117, 8),
	('LUT TAWAR', 117, 8),
	('PEGASING', 117, 8),
	('RUSIP ANTARA', 117, 8),
	('SILIH NARA', 117, 8),
	('BABUL MAKMUR', 118, 8),
	('BABUL RAHMAH', 118, 8),
	('BABUSSALAM', 118, 8),
	('BADAR', 118, 8),
	('BAMBEL', 118, 8),
	('BUKIT TUSAM', 118, 8),
	('DARUL HASANAH', 118, 8),
	('DELENG POKHISEN', 118, 8),
	('KETAMBE', 118, 8),
	('LAWE ALAS', 118, 8),
	('LAWE BULAN', 118, 8),
	('LAWE SIGALA-GALA', 118, 8),
	('LAWE SUMUR', 118, 8),
	('LEUSER', 118, 8),
	('SEMADAM', 118, 8),
	('TANAH ALAS', 118, 8),
	('BANDA ALAM', 119, 8),
	('BIREM BAYEUN', 119, 8),
	('DARUL AMAN', 119, 8),
	('DARUL FALAH', 119, 8),
	('DARUL IKSAN (IHSAN)', 119, 8),
	('IDI RAYEUK', 119, 8),
	('IDI TIMUR', 119, 8),
	('IDI TUNONG', 119, 8),
	('INDRA MAKMUR', 119, 8),
	('JULOK', 119, 8),
	('MADAT', 119, 8),
	('NURUSSALAM', 119, 8),
	('PANTE BIDARI (BEUDARI)', 119, 8),
	('PEUDAWA', 119, 8),
	('PEUNARON', 119, 8),
	('PEUREULAK', 119, 8),
	('PEUREULAK BARAT', 119, 8),
	('PEUREULAK TIMUR', 119, 8),
	('RANTAU SELAMAT', 119, 8),
	('RANTO PEUREULAK', 119, 8),
	('SERBA JADI', 119, 8),
	('SIMPANG JERNIH', 119, 8),
	('SIMPANG ULIM', 119, 8),
	('BAKTIYA', 120, 8),
	('BAKTIYA BARAT', 120, 8),
	('BANDA BARO', 120, 8),
	('COT GIREK', 120, 8),
	('DEWANTARA', 120, 8),
	('GEUREDONG PASE', 120, 8),
	('KUTA MAKMUR', 120, 8),
	('LANGKAHAN', 120, 8),
	('LAPANG', 120, 8),
	('LHOKSUKON', 120, 8),
	('MATANGKULI', 120, 8),
	('MEURAH MULIA', 120, 8),
	('MUARA BATU', 120, 8),
	('NIBONG', 120, 8),
	('NISAM', 120, 8),
	('NISAM ANTARA', 120, 8),
	('PAYA BAKONG', 120, 8),
	('PIRAK TIMUR', 120, 8),
	('SAMUDERA', 120, 8),
	('SAWANG', 120, 8),
	('SEUNUDDON (SEUNUDON)', 120, 8),
	('SIMPANG KRAMAT (KERAMAT)', 120, 8),
	('SYAMTALIRA ARON', 120, 8),
	('SYAMTALIRA BAYU', 120, 8),
	('TANAH JAMBO AYE', 120, 8),
	('TANAH LUAS', 120, 8),
	('TANAH PASIR', 120, 8),
	('BAITURRAHMAN', 121, 8),
	('BANDA RAYA', 121, 8),
	('JAYA BARU', 121, 8),
	('KUTA ALAM', 121, 8),
	('KUTA RAJA', 121, 8),
	('LUENG BATA', 121, 8),
	('MEURAXA', 121, 8),
	('SYIAH KUALA', 121, 8),
	('ULEE KARENG', 121, 8),
	('BENER KELIPAH', 122, 8),
	('BUKIT', 122, 8),
	('GAJAH PUTIH', 122, 8),
	('MESIDAH', 122, 8),
	('PERMATA', 122, 8),
	('PINTU RIME GAYO', 122, 8),
	('SYIAH UTAMA', 122, 8),
	('TIMANG GAJAH', 122, 8),
	('WIH PESAM', 122, 8),
	('GANDA PURA', 123, 8),
	('JANGKA', 123, 8),
	('JEUMPA', 123, 8),
	('JEUNIEB', 123, 8),
	('JULI', 123, 8),
	('KOTA JUANG', 123, 8),
	('KUALA', 123, 8),
	('KUTA BLANG', 123, 8),
	('MAKMUR', 123, 8),
	('PANDRAH', 123, 8),
	('PEUDADA', 123, 8),
	('PEULIMBANG (PLIMBANG)', 123, 8),
	('PEUSANGAN', 123, 8),
	('PEUSANGAN SELATAN', 123, 8),
	('PEUSANGAN SIBLAH KRUENG', 123, 8),
	('SAMALANGA', 123, 8),
	('SIMPANG MAMPLAM', 123, 8),
	('BLANG JERANGO', 124, 8),
	('BLANG KEJEREN', 124, 8),
	('BLANG PEGAYON', 124, 8),
	('DABUN GELANG (DEBUN GELANG)', 124, 8),
	('KUTA PANJANG', 124, 8),
	('PANTAN CUACA', 124, 8),
	('PINING (PINDING)', 124, 8),
	('PUTRI BETUNG', 124, 8),
	('RIKIT GAIB', 124, 8),
	('TERANGUN (TERANGON)', 124, 8),
	('TERIPE/TRIPE JAYA', 124, 8),
	('LANGSA BARAT', 125, 8),
	('LANGSA BARO', 125, 8),
	('LANGSA KOTA', 125, 8),
	('LANGSA LAMA', 125, 8),
	('LANGSA TIMUR', 125, 8),
	('BANDA SAKTI', 126, 8),
	('BLANG MANGAT', 126, 8),
	('MUARA DUA', 126, 8),
	('MUARA SATU', 126, 8),
	('BEUTONG', 127, 8),
	('BEUTONG ATEUH BANGGALANG', 127, 8),
	('DARUL MAKMUR', 127, 8),
	('KUALA PESISIR', 127, 8),
	('SEUNAGAN', 127, 8),
	('SEUNAGAN TIMUR', 127, 8),
	('SUKA MAKMUE', 127, 8),
	('TADU RAYA', 127, 8),
	('TRIPA MAKMUR', 127, 8),
	('BATEE', 128, 8),
	('DELIMA', 128, 8),
	('GEUMPANG', 128, 8),
	('GLUMPANG BARO', 128, 8),
	('GLUMPANG TIGA (GEULUMPANG TIGA)', 128, 8),
	('GRONG GRONG', 128, 8),
	('INDRAJAYA', 128, 8),
	('KEMBANG TANJONG (TANJUNG)', 128, 8),
	('KEUMALA', 128, 8),
	('KOTA SIGLI', 128, 8),
	('MANE', 128, 8),
	('MILA', 128, 8),
	('MUARA TIGA', 128, 8),
	('MUTIARA', 128, 8),
	('MUTIARA TIMUR', 128, 8),
	('PADANG TIJI', 128, 8),
	('PEUKAN BARO', 128, 8),
	('PIDIE', 128, 8),
	('SAKTI', 128, 8),
	('SIMPANG TIGA', 128, 8),
	('TANGSE', 128, 8),
	('TIRO/TRUSEB', 128, 8),
	('TITEUE', 128, 8),
	('BANDAR BARU', 129, 8),
	('BANDAR DUA', 129, 8),
	('JANGKA BUYA', 129, 8),
	('MEURAH DUA', 129, 8),
	('MEUREUDU', 129, 8),
	('PANTERAJA', 129, 8),
	('TRIENGGADENG', 129, 8),
	('ULIM', 129, 8),
	('SUKAJAYA', 130, 8),
	('ALAPAN (ALAFAN)', 131, 8),
	('SALANG', 131, 8),
	('SIMEULEU BARAT', 131, 8),
	('SIMEULEU TENGAH', 131, 8),
	('SIMEULEU TIMUR', 131, 8),
	('SIMEULUE CUT', 131, 8),
	('TELUK DALAM', 131, 8),
	('TEUPAH BARAT', 131, 8),
	('TEUPAH SELATAN', 131, 8),
	('TEUPAH TENGAH', 131, 8),
	('LONGKIB', 132, 8),
	('PENANGGALAN', 132, 8),
	('RUNDENG', 132, 8),
	('SIMPANG KIRI', 132, 8),
	('SULTAN DAULAT', 132, 8),
	('AEK KUASAN', 133, 9),
	('AEK LEDONG', 133, 9),
	('AEK SONGSONGAN', 133, 9),
	('AIR BATU', 133, 9),
	('AIR JOMAN', 133, 9),
	('BANDAR PASIR MANDOGE', 133, 9),
	('BANDAR PULAU', 133, 9),
	('BUNTU PANE', 133, 9),
	('KISARAN BARAT KOTA', 133, 9),
	('KISARAN TIMUR KOTA', 133, 9),
	('PULAU RAKYAT', 133, 9),
	('PULO BANDRING', 133, 9),
	('RAHUNING', 133, 9),
	('RAWANG PANCA ARGA', 133, 9),
	('SEI DADAP', 133, 9),
	('SEI KEPAYANG', 133, 9),
	('SEI KEPAYANG BARAT', 133, 9),
	('SEI KEPAYANG TIMUR', 133, 9),
	('SETIA JANJI', 133, 9),
	('SILAU LAUT', 133, 9),
	('TANJUNG BALAI', 133, 9),
	('TINGGI RAJA', 133, 9),
	('AIR PUTIH', 134, 9),
	('LIMAPULUH', 134, 9),
	('MEDANG DERAS', 134, 9),
	('SEI BALAI', 134, 9),
	('SEI SUKA', 134, 9),
	('TANJUNG TIRAM', 134, 9),
	('BINJAI BARAT', 135, 9),
	('BINJAI KOTA', 135, 9),
	('BINJAI SELATAN', 135, 9),
	('BINJAI TIMUR', 135, 9),
	('BINJAI UTARA', 135, 9),
	('BERAMPU (BRAMPU)', 136, 9),
	('GUNUNG SITEMBER', 136, 9),
	('LAE PARIRA', 136, 9),
	('PARBULUAN', 136, 9),
	('PEGAGAN HILIR', 136, 9),
	('SIDIKALANG', 136, 9),
	('SIEMPAT NEMPU', 136, 9),
	('SIEMPAT NEMPU HILIR', 136, 9),
	('SIEMPAT NEMPU HULU', 136, 9),
	('SILAHI SABUNGAN', 136, 9),
	('SILIMA PUNGGA-PUNGGA', 136, 9),
	('SITINJO', 136, 9),
	('SUMBUL', 136, 9),
	('TANAH PINEM', 136, 9),
	('TIGA LINGGA', 136, 9),
	('BATANG KUIS', 137, 9),
	('BERINGIN', 137, 9),
	('BIRU-BIRU', 137, 9),
	('DELI TUA', 137, 9),
	('GUNUNG MERIAH', 137, 9),
	('HAMPARAN PERAK', 137, 9),
	('KUTALIMBARU', 137, 9),
	('LABUHAN DELI', 137, 9),
	('LUBUK PAKAM', 137, 9),
	('NAMO RAMBE', 137, 9),
	('PAGAR MERBAU', 137, 9),
	('PANCUR BATU', 137, 9),
	('PANTAI LABU', 137, 9),
	('PATUMBAK', 137, 9),
	('PERCUT SEI TUAN', 137, 9),
	('SIBOLANGIT', 137, 9),
	('SINEMBAH TANJUNG MUDA HILIR', 137, 9),
	('SINEMBAH TANJUNG MUDA HULU', 137, 9),
	('SUNGGAL', 137, 9),
	('TANJUNG MORAWA', 137, 9),
	('GUNUNGSITOLI', 138, 9),
	('GUNUNGSITOLI ALO\\OA', 138, 9),
	('GUNUNGSITOLI BARAT', 138, 9),
	('GUNUNGSITOLI IDANOI', 138, 9),
	('GUNUNGSITOLI SELATAN', 138, 9),
	('GUNUNGSITOLI UTARA', 138, 9),
	('BAKTI RAJA', 139, 9),
	('DOLOK SANGGUL', 139, 9),
	('LINTONG NIHUTA', 139, 9),
	('ONAN GANJANG', 139, 9),
	('PAKKAT', 139, 9),
	('PARANGINAN', 139, 9),
	('PARLILITAN', 139, 9),
	('POLLUNG', 139, 9),
	('SIJAMA POLANG', 139, 9),
	('TARA BINTANG', 139, 9),
	('BARUS JAHE', 140, 9),
	('BRASTAGI (BERASTAGI)', 140, 9),
	('DOLAT RAYAT', 140, 9),
	('JUHAR', 140, 9),
	('KABANJAHE', 140, 9),
	('KUTA BULUH', 140, 9),
	('LAUBALENG', 140, 9),
	('MARDINDING', 140, 9),
	('MERDEKA', 140, 9),
	('MEREK', 140, 9),
	('MUNTE', 140, 9),
	('NAMA TERAN', 140, 9),
	('TIGA BINANGA', 140, 9),
	('TIGA PANAH', 140, 9),
	('TIGANDERKET', 140, 9),
	('BILAH BARAT', 141, 9),
	('BILAH HILIR', 141, 9),
	('BILAH HULU', 141, 9),
	('PANAI HILIR', 141, 9),
	('PANAI HULU', 141, 9),
	('PANAI TENGAH', 141, 9),
	('PANGKATAN', 141, 9),
	('RANTAU SELATAN', 141, 9),
	('RANTAU UTARA', 141, 9),
	('KAMPUNG RAKYAT', 142, 9),
	('KOTA PINANG', 142, 9),
	('SEI/SUNGAI KANAN', 142, 9),
	('SILANGKITANG', 142, 9),
	('TORGAMBA', 142, 9),
	('AEK KUO', 143, 9),
	('AEK NATAS', 143, 9),
	('KUALA LEDONG (KUALUH LEIDONG)', 143, 9),
	('KUALUH HILIR', 143, 9),
	('KUALUH HULU', 143, 9),
	('KUALUH SELATAN', 143, 9),
	('MARBAU', 143, 9),
	('NA IX-X', 143, 9),
	('BABALAN', 144, 9),
	('BAHOROK', 144, 9),
	('BATANG SERANGAN', 144, 9),
	('BESITANG', 144, 9),
	('BINJAI', 144, 9),
	('BRANDAN BARAT', 144, 9),
	('HINAI', 144, 9),
	('KUTAMBARU', 144, 9),
	('PADANG TUALANG', 144, 9),
	('PANGKALAN SUSU', 144, 9),
	('PEMATANG JAYA', 144, 9),
	('SALAPIAN', 144, 9),
	('SAWIT SEBERANG', 144, 9),
	('SECANGGANG', 144, 9),
	('SEI BINGE (BINGAI)', 144, 9),
	('SEI LEPAN', 144, 9),
	('SELESAI', 144, 9),
	('SIRAPIT (SERAPIT)', 144, 9),
	('STABAT', 144, 9),
	('TANJUNGPURA', 144, 9),
	('WAMPU', 144, 9),
	('BATAHAN', 145, 9),
	('BATANG NATAL', 145, 9),
	('BUKIT MALINTANG', 145, 9),
	('HUTA BARGOT', 145, 9),
	('KOTANOPAN', 145, 9),
	('LANGGA BAYU (LINGGA BAYU)', 145, 9),
	('LEMBAH SORIK MERAPI', 145, 9),
	('MUARA BATANG GADIS', 145, 9),
	('MUARA SIPONGI', 145, 9),
	('NAGA JUANG', 145, 9),
	('NATAL', 145, 9),
	('PAKANTAN', 145, 9),
	('PANYABUNGAN BARAT', 145, 9),
	('PANYABUNGAN KOTA', 145, 9),
	('PANYABUNGAN SELATAN', 145, 9),
	('PANYABUNGAN TIMUR', 145, 9),
	('PANYABUNGAN UTARA', 145, 9),
	('PUNCAK SORIK MARAPI/MERAPI', 145, 9),
	('RANTO BAEK/BAIK', 145, 9),
	('SIABU', 145, 9),
	('SINUNUKAN', 145, 9),
	('TAMBANGAN', 145, 9),
	('ULU PUNGKUT', 145, 9),
	('MEDAN AMPLAS', 146, 9),
	('MEDAN AREA', 146, 9),
	('MEDAN BARAT', 146, 9),
	('MEDAN BARU', 146, 9),
	('MEDAN BELAWAN KOTA', 146, 9),
	('MEDAN DELI', 146, 9),
	('MEDAN DENAI', 146, 9),
	('MEDAN HELVETIA', 146, 9),
	('MEDAN JOHOR', 146, 9),
	('MEDAN KOTA', 146, 9),
	('MEDAN LABUHAN', 146, 9),
	('MEDAN MAIMUN', 146, 9),
	('MEDAN MARELAN', 146, 9),
	('MEDAN PERJUANGAN', 146, 9),
	('MEDAN PETISAH', 146, 9),
	('MEDAN POLONIA', 146, 9),
	('MEDAN SELAYANG', 146, 9),
	('MEDAN SUNGGAL', 146, 9),
	('MEDAN TEMBUNG', 146, 9),
	('MEDAN TIMUR', 146, 9),
	('MEDAN TUNTUNGAN', 146, 9),
	('BAWOLATO', 147, 9),
	('BOTOMUZOI', 147, 9),
	('GIDO', 147, 9),
	('HILI SERANGKAI (HILISARANGGU)', 147, 9),
	('HILIDUHO', 147, 9),
	('IDANO GAWO', 147, 9),
	('MA\\U', 147, 9),
	('SOGAE ADU (SOGAEADU)', 147, 9),
	('SOMOLO-MOLO (SAMOLO)', 147, 9),
	('ULUGAWO', 147, 9),
	('LAHOMI (GAHORI)', 148, 9),
	('LOLOFITU MOI', 148, 9),
	('MANDREHE', 148, 9),
	('MANDREHE BARAT', 148, 9),
	('MANDREHE UTARA', 148, 9),
	('MORO\\O', 148, 9),
	('SIROMBU', 148, 9),
	('ULU MORO\\O (ULU NARWO)', 148, 9),
	('AMANDRAYA', 149, 9),
	('ARAMO', 149, 9),
	('BORONADU', 149, 9),
	('FANAYAMA', 149, 9),
	('GOMO', 149, 9),
	('HIBALA', 149, 9),
	('HILIMEGAI', 149, 9),
	('HILISALAWA\\AHE (HILISALAWAAHE)', 149, 9),
	('HURUNA', 149, 9),
	('LAHUSA', 149, 9),
	('LOLOMATUA', 149, 9),
	('LOLOWAU', 149, 9),
	('MANIAMOLO', 149, 9),
	('MAZINO', 149, 9),
	('MAZO', 149, 9),
	('O\\O\\U (OOU)', 149, 9),
	('ONOHAZUMBA', 149, 9),
	('PULAU-PULAU BATU', 149, 9),
	('PULAU-PULAU BATU BARAT', 149, 9),
	('PULAU-PULAU BATU TIMUR', 149, 9),
	('PULAU-PULAU BATU UTARA', 149, 9),
	('SIDUA\\ORI', 149, 9),
	('SIMUK', 149, 9),
	('SOMAMBAWA', 149, 9),
	('SUSUA', 149, 9),
	('TANAH MASA', 149, 9),
	('TOMA', 149, 9),
	('ULUNOYO', 149, 9),
	('ULUSUSUA', 149, 9),
	('UMBUNASI', 149, 9),
	('AFULU', 150, 9),
	('ALASA', 150, 9),
	('ALASA TALUMUZOI', 150, 9),
	('LAHEWA', 150, 9),
	('LAHEWA TIMUR', 150, 9),
	('LOTU', 150, 9),
	('NAMOHALU ESIWA', 150, 9),
	('SAWO', 150, 9),
	('SITOLU ORI', 150, 9),
	('TUGALA OYO', 150, 9),
	('TUHEMBERUA', 150, 9),
	('AEK NABARA BARUMUN', 151, 9),
	('BARUMUN', 151, 9),
	('BARUMUN SELATAN', 151, 9),
	('BARUMUN TENGAH', 151, 9),
	('BATANG LUBU SUTAM', 151, 9),
	('HURISTAK', 151, 9),
	('HUTA RAJA TINGGI', 151, 9),
	('LUBUK BARUMUN', 151, 9),
	('SIHAPAS BARUMUN', 151, 9),
	('SOSA', 151, 9),
	('SOSOPAN', 151, 9),
	('ULU BARUMUN', 151, 9),
	('BATANG ONANG', 152, 9),
	('DOLOK', 152, 9),
	('DOLOK SIGOMPULON', 152, 9),
	('HALONGONAN', 152, 9),
	('HULU SIHAPAS', 152, 9),
	('PADANG BOLAK', 152, 9),
	('PADANG BOLAK JULU', 152, 9),
	('PORTIBI', 152, 9),
	('SIMANGAMBAT', 152, 9),
	('PADANG SIDEMPUAN ANGKOLA JULU', 153, 9),
	('PADANG SIDEMPUAN BATUNADUA', 153, 9),
	('PADANG SIDEMPUAN HUTAIMBARU', 153, 9),
	('PADANG SIDEMPUAN SELATAN', 153, 9),
	('PADANG SIDEMPUAN TENGGARA', 153, 9),
	('PADANG SIDEMPUAN UTARA (PADANGSIDIMPUAN)', 153, 9),
	('KERAJAAN', 154, 9),
	('PAGINDAR', 154, 9),
	('PERGETTENG GETTENG SENGKUT', 154, 9),
	('SALAK', 154, 9),
	('SIEMPAT RUBE', 154, 9),
	('SITELLU TALI URANG JEHE', 154, 9),
	('SITELLU TALI URANG JULU', 154, 9),
	('TINADA', 154, 9),
	('SIANTAR BARAT', 155, 9),
	('SIANTAR MARIHAT', 155, 9),
	('SIANTAR MARIMBUN', 155, 9),
	('SIANTAR MARTOBA', 155, 9),
	('SIANTAR SELATAN', 155, 9),
	('SIANTAR SITALASARI', 155, 9),
	('SIANTAR TIMUR', 155, 9),
	('SIANTAR UTARA', 155, 9),
	('HARIAN', 156, 9),
	('NAINGGOLAN', 156, 9),
	('ONAN RUNGGU', 156, 9),
	('PALIPI', 156, 9),
	('PANGURURAN', 156, 9),
	('RONGGUR NIHUTA', 156, 9),
	('SIANJUR MULA-MULA', 156, 9),
	('SIMANINDO', 156, 9),
	('SITIO-TIO', 156, 9),
	('BANDAR KHALIFAH', 157, 9),
	('BINTANG BAYU', 157, 9),
	('DOLOK MASIHUL', 157, 9),
	('DOLOK MERAWAN', 157, 9),
	('KOTARIH', 157, 9),
	('PEGAJAHAN', 157, 9),
	('PERBAUNGAN', 157, 9),
	('SEI BAMBAN', 157, 9),
	('SEI RAMPAH', 157, 9),
	('SILINDA', 157, 9),
	('SIPISPIS', 157, 9),
	('TANJUNG BERINGIN', 157, 9),
	('TEBING SYAHBANDAR', 157, 9),
	('TELUK MENGKUDU', 157, 9),
	('SIBOLGA KOTA', 158, 9),
	('SIBOLGA SAMBAS', 158, 9),
	('SIBOLGA SELATAN', 158, 9),
	('SIBOLGA UTARA', 158, 9),
	('BANDAR HULUAN', 159, 9),
	('BANDAR MASILAM', 159, 9),
	('BOSAR MALIGAS', 159, 9),
	('DOLOK BATU NANGGAR', 159, 9),
	('DOLOK PANRIBUAN', 159, 9),
	('DOLOK PARDAMEAN', 159, 9),
	('DOLOK SILAU', 159, 9),
	('GIRSANG SIPANGAN BOLON', 159, 9),
	('GUNUNG MALELA', 159, 9),
	('GUNUNG MALIGAS', 159, 9),
	('HARANGGAOL HORISON', 159, 9),
	('HATONDUHAN', 159, 9),
	('HUTA BAYU RAJA', 159, 9),
	('JAWA MARAJA BAH JAMBI', 159, 9),
	('JORLANG HATARAN', 159, 9),
	('PANEI', 159, 9),
	('PANOMBEIAN PANEI', 159, 9),
	('PEMATANG BANDAR', 159, 9),
	('PEMATANG SIDAMANIK', 159, 9),
	('PEMATANG SILIMA HUTA', 159, 9),
	('PURBA', 159, 9),
	('RAYA', 159, 9),
	('RAYA KAHEAN', 159, 9),
	('SIANTAR', 159, 9),
	('SIDAMANIK', 159, 9),
	('SILIMAKUTA', 159, 9),
	('SILOU KAHEAN', 159, 9),
	('TANAH JAWA', 159, 9),
	('TAPIAN DOLOK', 159, 9),
	('UJUNG PADANG', 159, 9),
	('DATUK BANDAR', 160, 9),
	('DATUK BANDAR TIMUR', 160, 9),
	('SEI TUALANG RASO', 160, 9),
	('TANJUNG BALAI SELATAN', 160, 9),
	('TANJUNG BALAI UTARA', 160, 9),
	('TELUK NIBUNG', 160, 9),
	('AEK BILAH', 161, 9),
	('ANGKOLA BARAT (PADANG SIDEMPUAN)', 161, 9),
	('ANGKOLA SANGKUNUR', 161, 9),
	('ANGKOLA SELATAN (SIAIS)', 161, 9),
	('ANGKOLA TIMUR (PADANG SIDEMPUAN)', 161, 9),
	('ARSE', 161, 9),
	('BATANG ANGKOLA', 161, 9),
	('BATANG TORU', 161, 9),
	('MARANCAR', 161, 9),
	('MUARA BATANG TORU', 161, 9),
	('SAIPAR DOLOK HOLE', 161, 9),
	('SAYUR MATINGGI', 161, 9),
	('SIPIROK', 161, 9),
	('TANO TOMBANGAN ANGKOLA', 161, 9),
	('ANDAM DEWI', 162, 9),
	('BADIRI', 162, 9),
	('BARUS', 162, 9),
	('BARUS UTARA', 162, 9),
	('KOLANG', 162, 9),
	('LUMUT', 162, 9),
	('MANDUAMAS', 162, 9),
	('PANDAN', 162, 9),
	('PASARIBU TOBING', 162, 9),
	('PINANGSORI', 162, 9),
	('SARUDIK', 162, 9),
	('SIBABANGUN', 162, 9),
	('SIRANDORUNG', 162, 9),
	('SITAHUIS', 162, 9),
	('SORKAM', 162, 9),
	('SORKAM BARAT', 162, 9),
	('SOSOR GADONG', 162, 9),
	('SUKA BANGUN', 162, 9),
	('TAPIAN NAULI', 162, 9),
	('TUKKA', 162, 9),
	('ADIAN KOTING', 163, 9),
	('GAROGA', 163, 9),
	('MUARA', 163, 9),
	('PAGARAN', 163, 9),
	('PAHAE JAE', 163, 9),
	('PAHAE JULU', 163, 9),
	('PANGARIBUAN', 163, 9),
	('PARMONANGAN', 163, 9),
	('PURBATUA', 163, 9),
	('SIATAS BARITA', 163, 9),
	('SIBORONG-BORONG', 163, 9),
	('SIMANGUMBAN', 163, 9),
	('SIPAHUTAR', 163, 9),
	('SIPOHOLON', 163, 9),
	('TARUTUNG', 163, 9),
	('BAJENIS', 164, 9),
	('PADANG HILIR', 164, 9),
	('PADANG HULU', 164, 9),
	('TEBING TINGGI KOTA', 164, 9),
	('AJIBATA', 165, 9),
	('BALIGE', 165, 9),
	('BONATUA LUNASI', 165, 9),
	('BOR-BOR', 165, 9),
	('HABINSARAN', 165, 9),
	('LAGUBOTI', 165, 9),
	('LUMBAN JULU', 165, 9),
	('NASSAU', 165, 9),
	('PARMAKSIAN', 165, 9),
	('PINTU POHAN MERANTI', 165, 9),
	('PORSEA', 165, 9),
	('SIANTAR NARUMONDA', 165, 9),
	('SIGUMPAR', 165, 9),
	('SILAEN', 165, 9),
	('TAMPAHAN', 165, 9),
	('ULUAN', 165, 9)
	`)
	if err != nil {
		log.Fatal(err)
		fmt.Print("kecamatan not 7-9 insert")
	}
}