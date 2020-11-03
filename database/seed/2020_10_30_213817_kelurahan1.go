package seed

import (
	"log"

	"github.com/danangkonang/migrasion-go-cli/config"
)

func Kelurahan1() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO kelurahan1 (kelurahan,kecamatan_id,kabupaten_id,provinsi_id)VALUES
		('BAGENDUNG', 1, 1, 1),
		('BENDUNGAN', 1, 1, 1),
		('CIWADUK', 1, 1, 1),
		('CIWEDUS', 1, 1, 1),
		('KETILENG', 1, 1, 1),
		('CITANGKIL', 2, 1, 1),
		('DERINGO', 2, 1, 1),
		('KEBONSARI', 2, 1, 1),
		('LEBAKDENOK', 2, 1, 1),
		('SAMANGRAYA', 2, 1, 1),
		('TAMANBARU', 2, 1, 1),
		('WARNASARI', 2, 1, 1),
		('BANJAR NEGARA', 3, 1, 1),
		('GUNUNGSUGIH', 3, 1, 1),
		('KEPUH', 3, 1, 1),
		('KUBANGSARI', 3, 1, 1),
		('RANDAKARI', 3, 1, 1),
		('TEGALRATU', 3, 1, 1),
		('GEREM', 4, 1, 1),
		('GEROGOL/GROGOL', 4, 1, 1),
		('KOTASARI', 4, 1, 1),
		('RAWA ARUM', 4, 1, 1),
		('LEBAK GEDE', 5, 1, 1),
		('MEKARSARI', 5, 1, 1),
		('SURALAYA', 5, 1, 1),
		('TAMANSARI', 5, 1, 1),
		('CIPAISAN', 6, 1, 1),
		('CISEUREUH', 6, 1, 1),
		('CITALANG', 6, 1, 1),
		('KEBONDALEM', 6, 1, 1),
		('KOTABUMI', 6, 1, 1),
		('MUNJULJAYA', 6, 1, 1),
		('NAGRIKALER', 6, 1, 1),
		('NAGRIKIDUL', 6, 1, 1),
		('NAGRITENGAH', 6, 1, 1),
		('PABEAN', 6, 1, 1),
		('PURWAKARTA', 6, 1, 1),
		('PURWAMEKAR', 6, 1, 1),
		('RAMANUJU', 6, 1, 1),
		('SINDANGKASIH', 6, 1, 1),
		('TEGAL BUNDER', 6, 1, 1),
		('TEGALMUNJUL', 6, 1, 1),
		('BAYAH BARAT', 7, 2, 1),
		('BAYAH TIMUR', 7, 2, 1),
		('CIDIKIT', 7, 2, 1),
		('CIMANCAK', 7, 2, 1),
		('CISUREN', 7, 2, 1),
		('DARMASARI', 7, 2, 1),
		('PAMUBULAN', 7, 2, 1),
		('PASIRGOMBONG', 7, 2, 1),
		('SAWARNA', 7, 2, 1),
		('SAWARNA TIMUR', 7, 2, 1),
		('SUWAKAN', 7, 2, 1),
		('BOJONGMANIK', 8, 2, 1),
		('CIMAYANG', 8, 2, 1),
		('HARJAWANA', 8, 2, 1),
		('KADURAHAYU', 8, 2, 1),
		('KEBONCAU', 8, 2, 1),
		('MEKAR RAHAYU', 8, 2, 1),
		('MEKARMANIK', 8, 2, 1),
		('PARAKANBEUSI', 8, 2, 1),
		('PASIRBITUNG', 8, 2, 1),
		('CIBUNGUR', 9, 2, 1),
		('CIGEMBLONG', 9, 2, 1),
		('CIKADONGDONG', 9, 2, 1),
		('CIKARET', 9, 2, 1),
		('CIKATE', 9, 2, 1),
		('MUGIJAYA', 9, 2, 1),
		('PEUCANGPARI', 9, 2, 1),
		('WANGUNJAYA', 9, 2, 1),
		('BARUNAI', 10, 2, 1),
		('CIHARA', 10, 2, 1),
		('CIPARAHU', 10, 2, 1),
		('CITEPUSEUN', 10, 2, 1),
		('KARANGKAMULYAN', 10, 2, 1),
		('LEBAK PEUNDEUY', 10, 2, 1),
		('MEKARSARI', 10, 2, 1),
		('PANYAUNGAN', 10, 2, 1),
		('PONDOKPANJANG', 10, 2, 1),
		('CIAPUS', 11, 2, 1),
		('CIBEUREUM', 11, 2, 1),
		('CIHUJAN', 11, 2, 1),
		('CIJAKU', 11, 2, 1),
		('CIKARATUAN', 11, 2, 1),
		('CIMEGA', 11, 2, 1),
		('CIPALABUH', 11, 2, 1),
		('KANDANGSAPI', 11, 2, 1),
		('KAPUNDUHAN', 11, 2, 1),
		('MEKARJAYA', 11, 2, 1),
		('SUKASENANG', 11, 2, 1),
		('ANGGALAN', 12, 2, 1),
		('CIGOONG SELATAN', 12, 2, 1),
		('CIGOONG UTARA', 12, 2, 1),
		('CIKULUR', 12, 2, 1),
		('CURUGPANJANG', 12, 2, 1),
		('MUARADUA', 12, 2, 1),
		('MUNCANGKOPONG', 12, 2, 1),
		('PARAGE', 12, 2, 1),
		('PASIRGINTUNG', 12, 2, 1),
		('SUKADAYA', 12, 2, 1),
		('SUKAHARJA', 12, 2, 1),
		('SUMURBANDUNG', 12, 2, 1),
		('TAMAN JAYA', 12, 2, 1),
		('BANJARSARI', 13, 2, 1),
		('CIKAREO', 13, 2, 1),
		('CILELES', 13, 2, 1),
		('CIPADANG', 13, 2, 1),
		('DAROYON', 13, 2, 1),
		('GUMURUH', 13, 2, 1),
		('KUJANGSARI', 13, 2, 1),
		('MARGAMULYA', 13, 2, 1),
		('MEKARJAYA', 13, 2, 1),
		('PARUNGKUJANG', 13, 2, 1),
		('PASINDANGAN', 13, 2, 1),
		('PRABUGANTUNGAN', 13, 2, 1),
		('CIBARENO', 14, 2, 1),
		('CIJENGKOL', 14, 2, 1),
		('CIKAMUNDING', 14, 2, 1),
		('CIKATOMAS', 14, 2, 1),
		('CILOGRANG', 14, 2, 1),
		('CIREUNDEU', 14, 2, 1),
		('GIRIMUKTI', 14, 2, 1),
		('GUNUNGBATU', 14, 2, 1),
		('LEBAKTIPAR', 14, 2, 1),
		('PASIRBUNGUR', 14, 2, 1),
		('CIMARGA', 15, 2, 1),
		('GIRIMUKTI', 15, 2, 1),
		('GUNUNG ANTEN', 15, 2, 1),
		('INTEN JAYA', 15, 2, 1),
		('JAYAMANIK', 15, 2, 1),
		('JAYASARI', 15, 2, 1),
		('KARYA JAYA', 15, 2, 1),
		('MARGA JAYA', 15, 2, 1),
		('MARGALUYU', 15, 2, 1),
		('MARGATIRTA', 15, 2, 1),
		('MEKAR JAYA', 15, 2, 1),
		('MEKARMULYA', 15, 2, 1),
		('SANGIANG JAYA', 15, 2, 1),
		('SANGKAN MANIK', 15, 2, 1),
		('SARAGENI', 15, 2, 1),
		('SUDAMANIK', 15, 2, 1),
		('TAMBAK', 15, 2, 1),
		('BADUR', 16, 2, 1),
		('CEMPAKA', 16, 2, 1),
		('CIBARANI', 16, 2, 1),
		('CIRINTEN', 16, 2, 1),
		('DATARCAE', 16, 2, 1),
		('KADUDAMAS', 16, 2, 1),
		('KARANGNUNGGAL', 16, 2, 1),
		('KAROYA', 16, 2, 1),
		('NANGERANG', 16, 2, 1),
		('PARAKANLIMA', 16, 2, 1),
		('CANDI', 17, 2, 1),
		('CIBURUY', 17, 2, 1),
		('CIDADAP', 17, 2, 1),
		('CILAYANG', 17, 2, 1),
		('CIPINING', 17, 2, 1),
		('CURUGBITUNG', 17, 2, 1),
		('GURADOG', 17, 2, 1),
		('LEBAKKASIH', 17, 2, 1),
		('MAYAK', 17, 2, 1),
		('SEKARWANGI', 17, 2, 1),
		('BOJONGKONENG', 18, 2, 1),
		('BULAKAN', 18, 2, 1),
		('CIAKAR', 18, 2, 1),
		('CICARINGIN', 18, 2, 1),
		('CIGINGGANG', 18, 2, 1),
		('CIMANYANGRAY', 18, 2, 1),
		('CISAMPANG', 18, 2, 1),
		('GUNUNG KENCANA', 18, 2, 1),
		('GUNUNGKENDENG', 18, 2, 1),
		('KERAMATJAYA/KRAMATJAYA', 18, 2, 1),
		('SUKANEGARA', 18, 2, 1),
		('TANJUNGSARI INDAH', 18, 2, 1),
		('AWEH', 19, 2, 1),
		('CIKATAPIS', 19, 2, 1),
		('CILANGKAP', 19, 2, 1),
		('KALANGANYAR', 19, 2, 1),
		('PASIRKUPA', 19, 2, 1),
		('SANGIANG TANJUNG', 19, 2, 1),
		('SUKAMEKARSARI', 19, 2, 1),
		('BANJAR IRIGASI', 20, 2, 1),
		('BANJARSARI', 20, 2, 1),
		('CILADAEUN', 20, 2, 1),
		('LEBAKGEDONG', 20, 2, 1),
		('LEBAKSANGKA', 20, 2, 1),
		('LEBAKSITU', 20, 2, 1),
		('BOJONG MENTENG', 21, 2, 1),
		('CIBUNGUR', 21, 2, 1),
		('CISIMEUT', 21, 2, 1),
		('CISIMEUT RAYA', 21, 2, 1),
		('JALUPANG MULYA', 21, 2, 1),
		('KANEKES', 21, 2, 1),
		('LEBAK PARAHIANG', 21, 2, 1),
		('LEUWIDAMAR', 21, 2, 1),
		('MARGAWANGI', 21, 2, 1),
		('NAYAGATI', 21, 2, 1),
		('SANGKANWANGI', 21, 2, 1),
		('WANTISARI', 21, 2, 1),
		('BOLANG', 22, 2, 1),
		('CILANGKAHAN', 22, 2, 1),
		('CIPEUNDEUY', 22, 2, 1),
		('KADUJAJAR', 22, 2, 1),
		('KERSARATU', 22, 2, 1),
		('MALINGPING SELATAN', 22, 2, 1),
		('MALINGPING UTARA', 22, 2, 1),
		('PAGELARAN', 22, 2, 1),
		('RAHONG', 22, 2, 1),
		('SANGHIANG', 22, 2, 1),
		('SENANGHATI', 22, 2, 1),
		('SUKAMANAH', 22, 2, 1),
		('SUKARAJA', 22, 2, 1),
		('SUMBER WARAS', 22, 2, 1),
		('CIKARANG', 23, 2, 1),
		('CIMINYAK', 23, 2, 1),
		('GIRIJAGABAYA', 23, 2, 1),
		('JAGARAKSA', 23, 2, 1),
		('KARANGCOMBONG', 23, 2, 1),
		('LEUWICOO', 23, 2, 1),
		('MEKARWANGI', 23, 2, 1),
		('MUNCANG', 23, 2, 1),
		('PASIRNANGKA', 23, 2, 1),
		('SINDANGWANGI', 23, 2, 1),
		('SUKANAGARA', 23, 2, 1),
		('TANJUNGWANGI', 23, 2, 1),
		('CIBARENGKOK', 24, 2, 1),
		('CIMANDIRI', 24, 2, 1),
		('GUNUNG GEDE', 24, 2, 1),
		('HEGARMANAH', 24, 2, 1),
		('JATAKE', 24, 2, 1),
		('MEKARJAYA', 24, 2, 1),
		('PANGGARANGAN', 24, 2, 1),
		('SINDANGRATU', 24, 2, 1),
		('SITURAGEN', 24, 2, 1),
		('SOGONG', 24, 2, 1),
		('SUKAJADI', 24, 2, 1),
		('CIJORO LEBAK', 25, 2, 1),
		('CIJORO PASIR', 25, 2, 1),
		('CIMANGEUNGTEUNG', 25, 2, 1),
		('CITERAS', 25, 2, 1),
		('JATIMULYA', 25, 2, 1),
		('KOLELET WETAN', 25, 2, 1),
		('MEKARSARI', 25, 2, 1),
		('MUARA CIUJUNG BARAT', 25, 2, 1),
		('MUARA CIUJUNG TIMUR', 25, 2, 1),
		('NAMENG', 25, 2, 1),
		('NARIMBANG MULIA', 25, 2, 1),
		('PABUARAN', 25, 2, 1),
		('PASIRTANJUNG', 25, 2, 1),
		('RANGKASBITUNG BARAT', 25, 2, 1),
		('RANGKASBITUNG TIMUR', 25, 2, 1),
		('SUKAMANAH', 25, 2, 1),
		('BANGUNMEKAR', 26, 2, 1),
		('CALUNGBUNGUR', 26, 2, 1),
		('CIUYAH', 26, 2, 1),
		('MARAYA', 26, 2, 1),
		('MARGALUYU', 26, 2, 1),
		('MEKARSARI', 26, 2, 1),
		('PAJA', 26, 2, 1),
		('PAJAGAN', 26, 2, 1),
		('PARUNGSARI', 26, 2, 1),
		('SAJIRA', 26, 2, 1),
		('SAJIRAMEKAR', 26, 2, 1),
		('SINDANGSARI', 26, 2, 1),
		('SUKAJAYA', 26, 2, 1),
		('SUKAMARGA', 26, 2, 1),
		('SUKARAME', 26, 2, 1),
		('BOJEN', 27, 2, 1),
		('CILEBANG', 27, 2, 1),
		('CIMANIS', 27, 2, 1),
		('CIMANIS', 27, 2, 1),
		('CIROMPANG', 27, 2, 1),
		('CITUJAH', 27, 2, 1),
		('HARIANG', 27, 2, 1),
		('KERTARAHARJA', 27, 2, 1),
		('KUTAMEKAR', 27, 2, 1),
		('MAJASARI', 27, 2, 1),
		('PANGKALAN', 27, 2, 1),
		('SINDANGLAYA', 27, 2, 1),
		('SOBANG', 27, 2, 1),
		('SOBANG', 27, 2, 1),
		('SUKAJAYA', 27, 2, 1),
		('SUKAMAJU', 27, 2, 1),
		('SUKARESMI', 27, 2, 1),
		('TELUKLADA', 27, 2, 1),
		('BEJOD', 28, 2, 1),
		('CIKEUSIK', 28, 2, 1),
		('CILANGKAP', 28, 2, 1),
		('CIPEDANG', 28, 2, 1),
		('CIPEUCANG', 28, 2, 1),
		('CISARAP', 28, 2, 1),
		('KARANGPAMINDANGAN', 28, 2, 1),
		('KETAPANG', 28, 2, 1),
		('MUARA', 28, 2, 1),
		('PARUNGPANJANG', 28, 2, 1),
		('PARUNGSARI', 28, 2, 1),
		('SUKATANI', 28, 2, 1),
		('WANASALAM', 28, 2, 1),
		('BANJARSARI', 29, 2, 1),
		('BAROS', 29, 2, 1),
		('CEMPAKA', 29, 2, 1),
		('CIBUAH', 29, 2, 1),
		('JAGABAYA', 29, 2, 1),
		('PADASUKA', 29, 2, 1),
		('PASIRTANGKIL', 29, 2, 1),
		('SELARAJA', 29, 2, 1),
		('SINDANGSARI', 29, 2, 1),
		('SUKARAJA', 29, 2, 1),
		('SUKARENDAH', 29, 2, 1),
		('WARUNGGUNUNG', 29, 2, 1),
		('CADASARI', 30, 3, 1),
		('CIINJUK', 30, 3, 1),
		('CIKENTRUNG', 30, 3, 1),
		('KADUELA', 30, 3, 1),
		('KADUENGANG', 30, 3, 1),
		('KAUNGCAANG', 30, 3, 1),
		('KORANJI', 30, 3, 1),
		('KURUNGDAHU', 30, 3, 1),
		('PASIRPEUTEUY', 30, 3, 1),
		('TANAGARA', 30, 3, 1),
		('TAPOS', 30, 3, 1),
		('BANJARMASIN', 31, 3, 1),
		('CARITA', 31, 3, 1),
		('CINOYONG', 31, 3, 1),
		('KAWOYANG', 31, 3, 1),
		('PEJAMBEN', 31, 3, 1),
		('SINDANGLAUT', 31, 3, 1),
		('SUKAJADI', 31, 3, 1),
		('SUKANAGARA', 31, 3, 1),
		('SUKARAME', 31, 3, 1),
		('TEMBONG', 31, 3, 1),
		('CIBALIUNG', 32, 3, 1),
		('CIBINGBIN', 32, 3, 1),
		('CIHANJUANG', 32, 3, 1),
		('CURUG', 32, 3, 1),
		('MAHENDRA', 32, 3, 1),
		('MENDUNG', 32, 3, 1),
		('SORONGAN', 32, 3, 1),
		('SUDIMANIK', 32, 3, 1),
		('SUKAJADI', 32, 3, 1),
		('BANYUASIH', 33, 3, 1),
		('CIGEULIS', 33, 3, 1),
		('CISEUREUHEUN', 33, 3, 1),
		('KARANGBOLONG', 33, 3, 1),
		('KARYABUANA', 33, 3, 1),
		('KATUMBIRI', 33, 3, 1),
		('SINARJAYA', 33, 3, 1),
		('TARUMANAGARA', 33, 3, 1),
		('WARINGINJAYA', 33, 3, 1),
		('BABAKANLOR', 34, 3, 1),
		('BANGKUYUNG', 34, 3, 1),
		('CENING', 34, 3, 1),
		('CIPICUNG', 34, 3, 1),
		('DAHU', 34, 3, 1),
		('KARYASARI', 34, 3, 1),
		('KARYAUTAMA', 34, 3, 1),
		('MEKARJAYA', 34, 3, 1),
		('PADAHAYU', 34, 3, 1),
		('TEGAL', 34, 3, 1),
		('CIKADONGDONG', 35, 3, 1),
		('CIKEUSIK', 35, 3, 1),
		('CIKIRUHWETAN', 35, 3, 1),
		('CURUGCIUNG', 35, 3, 1),
		('LEUWIBALANG', 35, 3, 1),
		('NANGGALA', 35, 3, 1),
		('PARUNGKOKOSAN', 35, 3, 1),
		('RANCASENENG', 35, 3, 1),
		('SUKAMULYA', 35, 3, 1),
		('SUKASENENG', 35, 3, 1),
		('SUKAWARIS', 35, 3, 1),
		('SUMURBATU', 35, 3, 1),
		('TANJUNGAN', 35, 3, 1),
		('UMBULAN', 35, 3, 1),
		('BATUBANTAR', 36, 3, 1),
		('CIMANUK', 36, 3, 1),
		('DALEMBALAR', 36, 3, 1),
		('GUNUNGCUPU', 36, 3, 1),
		('GUNUNGDATAR', 36, 3, 1),
		('KADUBUNGBANG', 36, 3, 1),
		('KADUDODOL', 36, 3, 1),
		('KADUMADANG', 36, 3, 1),
		('KUPAHANDAP', 36, 3, 1),
		('ROCEK', 36, 3, 1),
		('SEKONG', 36, 3, 1),
		('BATURANJANG', 37, 3, 1),
		('CIKADUEUN', 37, 3, 1),
		('CURUGBARANG', 37, 3, 1),
		('KADUGADUNG', 37, 3, 1),
		('KALANGGUNUNG', 37, 3, 1),
		('KONCANG', 37, 3, 1),
		('PALANYAR', 37, 3, 1),
		('PARUMASAN', 37, 3, 1),
		('PASIREURIH', 37, 3, 1),
		('PASIRMAE', 37, 3, 1),
		('CIBARANI', 38, 3, 1),
		('CIHERANGJAYA', 38, 3, 1),
		('CISEREH', 38, 3, 1),
		('KADURONYOK', 38, 3, 1),
		('KONDANGJAYA', 38, 3, 1),
		('KUBANGKONDANG', 38, 3, 1),
		('PALEMBANG', 38, 3, 1),
		('PASIREURIH', 38, 3, 1),
		('RAWASARI', 38, 3, 1),
		('BABADSARI', 39, 3, 1),
		('BANYURESMI', 39, 3, 1),
		('CITAMAN', 39, 3, 1),
		('JANAKA', 39, 3, 1),
		('JAYAMEKAR', 39, 3, 1),
		('JIPUT', 39, 3, 1),
		('PAMARAYAN', 39, 3, 1),
		('SALAPRAYA', 39, 3, 1),
		('SAMPANGBITUNG', 39, 3, 1),
		('SIKULAN', 39, 3, 1),
		('SUKACAI', 39, 3, 1),
		('SUKAMANAH', 39, 3, 1),
		('TENJOLAHANG', 39, 3, 1),
		('BANJARSARI', 40, 3, 1),
		('BAYUMUNDU', 40, 3, 1),
		('CAMPAKA', 40, 3, 1),
		('CIPUTRI', 40, 3, 1),
		('KADUGEMBLO', 40, 3, 1),
		('MANDALASARI', 40, 3, 1),
		('PALURAHAN', 40, 3, 1),
		('SANINTEN', 40, 3, 1),
		('SUKAMANAH', 40, 3, 1),
		('SUKASARI', 40, 3, 1),
		('CIGADUNG', 41, 3, 1),
		('JUHUT', 41, 3, 1),
		('KADUMERAK', 41, 3, 1),
		('PAGADUNGAN', 41, 3, 1),
		('AWILEGA', 42, 3, 1),
		('BANGKONOL', 42, 3, 1),
		('GERENDONG', 42, 3, 1),
		('KARANGSETRA', 42, 3, 1),
		('KORONCONG', 42, 3, 1),
		('PAKULURAN', 42, 3, 1),
		('PANIIS', 42, 3, 1),
		('PASIRJAKSA', 42, 3, 1),
		('PASIRKARANG', 42, 3, 1),
		('SETRAJAYA', 42, 3, 1),
		('SUKAJAYA', 42, 3, 1),
		('TEGALONGOK', 42, 3, 1),
		('CILAJA', 43, 3, 1),
		('KARATON', 43, 3, 1),
		('PAGERBATU', 43, 3, 1),
		('SARUNI', 43, 3, 1),
		('SUKARATU', 43, 3, 1),
		('CIKONENG', 44, 3, 1),
		('CIKUMBUEUN', 44, 3, 1),
		('CURUGLEMO', 44, 3, 1),
		('GIRIPAWANA', 44, 3, 1),
		('GUNUNGSARI', 44, 3, 1),
		('KURUNGKAMBING', 44, 3, 1),
		('MANDALASARI', 44, 3, 1),
		('MANDALAWANGI', 44, 3, 1),
		('NEMBOL', 44, 3, 1),
		('PANDAT', 44, 3, 1),
		('PANJANGJAYA', 44, 3, 1),
		('PARI', 44, 3, 1),
		('RAMEA', 44, 3, 1),
		('SINARJAYA', 44, 3, 1),
		('SIRNAGALIH', 44, 3, 1),
		('KADUBELANG', 45, 3, 1),
		('KADUJANGKUNG', 45, 3, 1),
		('MEDONG', 45, 3, 1),
		('MEKARJAYA', 45, 3, 1),
		('PAREANG', 45, 3, 1),
		('RANCABUGEL', 45, 3, 1),
		('SUKAMULYA', 45, 3, 1),
		('WIRASINGA', 45, 3, 1),
		('ALASWANGI', 46, 3, 1),
		('CIGANDENG', 46, 3, 1),
		('CILABANBULAN', 46, 3, 1),
		('KADUPAYUNG', 46, 3, 1),
		('KANANGA', 46, 3, 1),
		('MENES', 46, 3, 1),
		('MURUY', 46, 3, 1),
		('PURWARAJA', 46, 3, 1),
		('RAMAYA', 46, 3, 1),
		('SINDANGKARYA', 46, 3, 1),
		('SUKAMANAH', 46, 3, 1),
		('TEGALWANGI', 46, 3, 1),
		('CIBITUNG', 47, 3, 1),
		('CURUGLANGLANG', 47, 3, 1),
		('GUNUNGBATU', 47, 3, 1),
		('KOTADUKUH', 47, 3, 1),
		('LEBAK', 47, 3, 1),
		('MUNJUL', 47, 3, 1),
		('PANACARAN', 47, 3, 1),
		('PASANGGRAHAN', 47, 3, 1),
		('SUKASABA', 47, 3, 1),
		('BABAKAN KALANG ANYAR', 48, 3, 1),
		('KABAYAN', 48, 3, 1),
		('KADOMAS', 48, 3, 1),
		('PANDEGLANG', 48, 3, 1),
		('CITEUREUP', 49, 3, 1),
		('GOMBONG', 49, 3, 1),
		('MEKARJAYA', 49, 3, 1),
		('MEKARSARI', 49, 3, 1),
		('PANIMBANGJAYA', 49, 3, 1),
		('TANJUNGJAYA', 49, 3, 1),
		('BABAKANKEUSIK', 50, 3, 1),
		('CIAWI', 50, 3, 1),
		('CIMOYAN', 50, 3, 1),
		('IDAMAN', 50, 3, 1),
		('PASIRGADUNG', 50, 3, 1),
		('PATIA', 50, 3, 1),
		('RAHAYU', 50, 3, 1),
		('SIMPANGTIGA', 50, 3, 1),
		('SURIANEUN', 50, 3, 1),
		('TURUS', 50, 3, 1),
		('BUNGURCOPONG', 51, 3, 1),
		('CIHERANG', 51, 3, 1),
		('CILILITAN', 51, 3, 1),
		('GANGGAENG', 51, 3, 1),
		('KADUBERA', 51, 3, 1),
		('KADUPANDAK', 51, 3, 1),
		('KOLELET', 51, 3, 1),
		('PASIRPANJANG', 51, 3, 1),
		('PASIRSEDANG', 51, 3, 1),
		('CIANDUR', 52, 3, 1),
		('GIRIJAYA', 52, 3, 1),
		('KADUDAMPIT', 52, 3, 1),
		('LANGENSARI', 52, 3, 1),
		('MAJAU', 52, 3, 1),
		('MEDALSARI', 52, 3, 1),
		('MEKARWANGI', 52, 3, 1),
		('PARIGI', 52, 3, 1),
		('SAKETI', 52, 3, 1),
		('SINDANGHAYU', 52, 3, 1),
		('SODONG', 52, 3, 1),
		('SUKALANGU', 52, 3, 1),
		('TALAGASARI', 52, 3, 1),
		('WANAGIRI', 52, 3, 1),
		('BOJONGMANIK', 53, 3, 1),
		('CAMPAKAWARNA', 53, 3, 1),
		('CIODENG', 53, 3, 1),
		('KADUMALATI', 53, 3, 1),
		('PASIRDURUNG', 53, 3, 1),
		('PASIRLANCAR', 53, 3, 1),
		('PASIRLOA', 53, 3, 1),
		('PASIRTENJO', 53, 3, 1),
		('SINDANGRESMI', 53, 3, 1),
		('CIGORONDONG', 54, 3, 1),
		('KERTAJAYA', 54, 3, 1),
		('KERTAMUKTI', 54, 3, 1),
		('SUMBERJAYA', 54, 3, 1),
		('TAMANJAYA', 54, 3, 1),
		('TUNGGALJAYA', 54, 3, 1),
		('UJUNGJAYA', 54, 3, 1),
		('ANYAR', 55, 4, 1),
		('BANDULU', 55, 4, 1),
		('BANJARSARI', 55, 4, 1),
		('BUNIHARA', 55, 4, 1),
		('CIKONENG', 55, 4, 1),
		('GROGOL INDAH', 55, 4, 1),
		('KOSAMBI RONYOK', 55, 4, 1),
		('MEKARSARI', 55, 4, 1),
		('SINDANGKARYA', 55, 4, 1),
		('SINDANGMANDI', 55, 4, 1),
		('TAMBANG AYAM', 55, 4, 1),
		('TANJUNGMANIS', 55, 4, 1),
		('BOJONEGARA', 56, 4, 1),
		('KARANGKEPUH', 56, 4, 1),
		('KERTASANA', 56, 4, 1),
		('LAMBANGSARI', 56, 4, 1),
		('MANGKUNEGARA', 56, 4, 1),
		('MARGAGIRI', 56, 4, 1),
		('MEKAR JAYA', 56, 4, 1),
		('PAKUNCEN', 56, 4, 1),
		('PENGARENGAN', 56, 4, 1),
		('UKIRSARI', 56, 4, 1),
		('WANAKARTA', 56, 4, 1),
		('CARENANG', 57, 4, 1),
		('MANDAYA', 57, 4, 1),
		('MEKARSARI', 57, 4, 1),
		('PAMANUK', 57, 4, 1),
		('PANENJOAN', 57, 4, 1),
		('RAGASMESIGIT', 57, 4, 1),
		('TERAS', 57, 4, 1),
		('WALIKUKUN', 57, 4, 1),
		('BAKUNG', 58, 4, 1),
		('CIKANDE', 58, 4, 1),
		('CIKANDE PERMAI', 58, 4, 1),
		('GEMBOR UDIK', 58, 4, 1),
		('JULANG', 58, 4, 1),
		('KAMURANG', 58, 4, 1),
		('KOPER', 58, 4, 1),
		('LEUWILIMUS', 58, 4, 1),
		('NAMBO UDIK', 58, 4, 1),
		('PARIGI', 58, 4, 1),
		('SITUTERATE', 58, 4, 1),
		('SONGGOM JAYA', 58, 4, 1),
		('SUKATANI', 58, 4, 1),
		('BANTAR PANJANG', 59, 4, 1),
		('CIKEUSAL', 59, 4, 1),
		('CILAYANG', 59, 4, 1),
		('CILAYANG GUHA', 59, 4, 1),
		('CIMAUNG', 59, 4, 1),
		(600, 'DAHU', 59, 4, 1),
		(601, 'GANDAYASA', 59, 4, 1),
		(602, 'HARUNDANG', 59, 4, 1),
		(603, 'KATULISAN', 59, 4, 1),
		(604, 'MONGPOK', 59, 4, 1),
		(605, 'PANOSOGAN', 59, 4, 1),
		(606, 'PANYABRANGAN', 59, 4, 1),
		(607, 'SUKAMAJU', 59, 4, 1),
		(608, 'SUKAMENAK', 59, 4, 1),
		(609, 'SUKARAJA', 59, 4, 1),
		(610, 'SUKARAME', 59, 4, 1),
		(611, 'SUKARATU', 59, 4, 1),
		(612, 'BANTARWANGI', 60, 4, 1),
		(613, 'BANTARWARU', 60, 4, 1),
		(614, 'BAROS JAYA', 60, 4, 1),
		(615, 'BULAKAN', 60, 4, 1),
		(616, 'CIKOLELET', 60, 4, 1),
		(617, 'CINANGKA', 60, 4, 1),
		(618, 'KAMASAN', 60, 4, 1),
		(619, 'KARANG SURAGA', 60, 4, 1),
		(620, 'KUBANG BAROS', 60, 4, 1),
		(621, 'MEKARSARI', 60, 4, 1),
		(622, 'PASAURAN', 60, 4, 1),
		(623, 'RANCASANGGAL', 60, 4, 1),
		(624, 'SINDANGLAYA', 60, 4, 1),
		(625, 'UMBUL TANJUNG', 60, 4, 1),
		(626, 'CEMPLANG', 61, 4, 1),
		(627, 'CIAPUS', 61, 4, 1),
		(628, 'CIOMAS', 61, 4, 1),
		(629, 'CIOMAS RAHAYU', 61, 4, 1),
		(630, 'CISITU', 61, 4, 1),
		(631, 'CITAMAN', 61, 4, 1),
		(632, 'KOTA BATU', 61, 4, 1),
		(633, 'LALADON', 61, 4, 1),
		(634, 'LEBAK', 61, 4, 1),
		(635, 'MEKARJAYA', 61, 4, 1),
		(636, 'PADASUKA', 61, 4, 1),
		(637, 'PAGELARAN', 61, 4, 1),
		(638, 'PANYAUNGAN JAYA', 61, 4, 1),
		(639, 'PARAKAN', 61, 4, 1),
		(640, 'PONDOK KAHURU', 61, 4, 1),
		(641, 'SIKETUG', 61, 4, 1),
		(642, 'SUKABARES', 61, 4, 1),
		(643, 'SUKADANA', 61, 4, 1),
		(644, 'SUKAHARJA', 61, 4, 1),
		(645, 'SUKAMAKMUR', 61, 4, 1),
		(646, 'SUKARENA', 61, 4, 1),
		(647, 'UJUNGTEBU', 61, 4, 1),
		(648, 'BANJARAGUNG', 62, 4, 1),
		(649, 'BANJARSARI', 62, 4, 1),
		(650, 'CIPOCOK JAYA', 62, 4, 1),
		(651, 'DALUNG', 62, 4, 1),
		(652, 'GELAM', 62, 4, 1),
		(653, 'KARUNDANG', 62, 4, 1),
		(654, 'PANANCANGAN', 62, 4, 1),
		(655, 'TEMBONG', 62, 4, 1),
		(656, 'BEBERAN', 63, 4, 1),
		(657, 'BUMIJAYA', 63, 4, 1),
		(658, 'CIGELAM', 63, 4, 1),
		(659, 'CIRUAS', 63, 4, 1),
		(660, 'CITEREP', 63, 4, 1),
		(661, 'GOSARA', 63, 4, 1),
		(662, 'KADIKARAN', 63, 4, 1),
		(663, 'KASERANGAN', 63, 4, 1),
		(664, 'KEPANDEAN', 63, 4, 1),
		(665, 'PAMONG', 63, 4, 1),
		(666, 'PELAWAD', 63, 4, 1),
		(667, 'PENGGALANG', 63, 4, 1),
		(668, 'PULO', 63, 4, 1),
		(669, 'RANJENG', 63, 4, 1),
		(670, 'SINGAMERTA', 63, 4, 1),
		(671, 'BINONG', 64, 4, 1),
		(672, 'CILAKU', 64, 4, 1),
		(673, 'CIPETE', 64, 4, 1),
		(674, 'CUKANG GALIH', 64, 4, 1),
		(675, 'CURUG', 64, 4, 1),
		(676, 'CURUG KULON', 64, 4, 1),
		(677, 'CURUG WETAN', 64, 4, 1),
		(678, 'CURUGMANIS', 64, 4, 1),
		(679, 'KADU', 64, 4, 1),
		(680, 'KADU JAYA', 64, 4, 1),
		(681, 'KAMANISAN', 64, 4, 1),
		(682, 'PANCALAKSANA', 64, 4, 1),
		(683, 'SUKA BAKTI', 64, 4, 1),
		(684, 'SUKAJAYA', 64, 4, 1),
		(685, 'SUKALAKSANA', 64, 4, 1),
		(686, 'SUKAWANA', 64, 4, 1),
		(687, 'TINGGAR', 64, 4, 1),
		(688, 'BOJOT', 65, 4, 1),
		(689, 'CEMPLANG', 65, 4, 1),
		(690, 'JAWILAN', 65, 4, 1),
		(691, 'JUNTI', 65, 4, 1),
		(692, 'KAREO', 65, 4, 1),
		(693, 'MAJASARI', 65, 4, 1),
		(694, 'PAGINTUNGAN', 65, 4, 1),
		(695, 'PARAKAN', 65, 4, 1),
		(696, 'PASIRBUYUT', 65, 4, 1),
		(697, 'BANTEN', 66, 4, 1),
		(698, 'BENDUNG', 66, 4, 1),
		(699, 'KASEMEN', 66, 4, 1),
		('KASUNYATAN', 66, 4, 1),
		('KILASAH', 66, 4, 1),
		('MARGALUYU', 66, 4, 1),
		('MESJID PRIYAYI', 66, 4, 1),
		('SAWAH LUHUR', 66, 4, 1),
		('TERUMBU', 66, 4, 1),
		('WARUNG JAUD', 66, 4, 1),
		('BARENGKOK', 67, 4, 1),
		('CIAGEL', 67, 4, 1),
		('CIJERUK', 67, 4, 1),
		('KETOS', 67, 4, 1),
		('KIBIN', 67, 4, 1),
		('NAGARA', 67, 4, 1),
		('NAMBO ILIR', 67, 4, 1),
		('SUKAMAJU', 67, 4, 1),
		('TAMBAK', 67, 4, 1),
		('BABAKAN JAYA', 68, 4, 1),
		('CARENANG UDIK', 68, 4, 1),
		('CIDAHU', 68, 4, 1),
		('GABUS', 68, 4, 1),
		('GARUT', 68, 4, 1),
		('KOPO', 68, 4, 1),
		('MEKARBARU', 68, 4, 1),
		('NANGGUNG', 68, 4, 1),
		('NYOMPOK', 68, 4, 1),
		('RANCASUMUR', 68, 4, 1),
		('CISAIT', 69, 4, 1),
		('DUKUH', 69, 4, 1),
		('JERUKTIPIS', 69, 4, 1),
		('KENDAYAKAN', 69, 4, 1),
		('KRAGILAN', 69, 4, 1),
		('KRAMATJATI', 69, 4, 1),
		('PEMATANG', 69, 4, 1),
		('SENTUL', 69, 4, 1),
		('SILEBU', 69, 4, 1),
		('SUKAJADI', 69, 4, 1),
		('TEGALMAJA', 69, 4, 1),
		('UNDAR ANDIR', 69, 4, 1),
		('HARJATANI', 70, 4, 1),
		('KRAMATWATU', 70, 4, 1),
		('LEBAKWANA', 70, 4, 1),
		('MARGASANA', 70, 4, 1),
		('PAMENGKANG', 70, 4, 1),
		('PEGADINGAN', 70, 4, 1),
		('PEJATEN', 70, 4, 1),
		('PELAMUNAN', 70, 4, 1),
		('SERDANG', 70, 4, 1),
		('TELUK TERATE', 70, 4, 1),
		('TERATE', 70, 4, 1),
		('TONJONG', 70, 4, 1),
		('TOYOMERTO', 70, 4, 1),
		('WANAYASA', 70, 4, 1),
		('BOLANG', 71, 4, 1),
		('KAMARUTON', 71, 4, 1),
		('KEBONRATU', 71, 4, 1),
		('KENCANA HARAPAN', 71, 4, 1),
		('LEBAK KEPUH', 71, 4, 1),
		('LEBAKWANGI', 71, 4, 1),
		('PEGANDIKAN (PEGANTIKAN)', 71, 4, 1),
		('PURWADADI', 71, 4, 1),
		('TERASBENDUNG', 71, 4, 1),
		('TIREM', 71, 4, 1),
		('ANGSANA', 72, 4, 1),
		('BALE KENCANA', 72, 4, 1),
		('BALEKAMBANG', 72, 4, 1),
		('BATUKUDA', 72, 4, 1),
		('CIKEDUNG', 72, 4, 1),
		('CIWARNA', 72, 4, 1),
		('LABUHAN', 72, 4, 1),
		('MANCAK', 72, 4, 1),
		('PASIRWARU', 72, 4, 1),
		('SANGIANG', 72, 4, 1),
		('SIGEDONG', 72, 4, 1),
		('TALAGA', 72, 4, 1),
		('WARINGIN', 72, 4, 1),
		('WINONG', 72, 4, 1),
		('BARUGBUG', 73, 4, 1),
		('BATUKUWUNG', 73, 4, 1),
		('BUGEL', 73, 4, 1),
		('CIBOJONG', 73, 4, 1),
		('CIOMAS', 73, 4, 1),
		('CIPAYUNG', 73, 4, 1),
		('CISAAT', 73, 4, 1),
		('CITASUK', 73, 4, 1),
		('CURUG GOONG', 73, 4, 1),
		('KADU KEMPONG', 73, 4, 1),
		('KADUBEUREUM', 73, 4, 1),
		('KALUMPANG', 73, 4, 1),
		('KRAMATLABAN', 73, 4, 1),
		('PADARINCANG', 73, 4, 1),
		('BINONG', 74, 4, 1),
		('DAMPING', 74, 4, 1),
		('KAMPUNG BARU', 74, 4, 1),
		('KEBON CAU', 74, 4, 1),
		('PAMARAYAN', 74, 4, 1),
		('PASIR KEMBANG', 74, 4, 1),
		('PASIRLIMUS', 74, 4, 1),
		('PUDAR', 74, 4, 1),
		('SANGIANG', 74, 4, 1),
		('WIRANA', 74, 4, 1),
		('BOJONG NANGKA', 75, 4, 1),
		('CIRANGKONG', 75, 4, 1),
		('CIREUNDEU', 75, 4, 1),
		('KADUGENEP', 75, 4, 1),
		('KAMPUNG BARU', 75, 4, 1),
		('KUBANG JAYA', 75, 4, 1),
		('MEKARBARU', 75, 4, 1),
		('NAGARA PADANG', 75, 4, 1),
		('PADASUKA', 75, 4, 1),
		('PETIR', 75, 4, 1),
		('SANDING', 75, 4, 1),
		('SEUAT', 75, 4, 1),
		('SEUAT JAYA', 75, 4, 1),
		('SINDANGSARI', 75, 4, 1),
		('TAMBILUK', 75, 4, 1),
		('DOMAS', 76, 4, 1),
		('KALAPIAN', 76, 4, 1),
		('KESERANGAN', 76, 4, 1),
		('KUBANG PUJI', 76, 4, 1),
		('LINDUK', 76, 4, 1),
		('PONTANG', 76, 4, 1),
		('PULO KENCANA', 76, 4, 1),
		('SINGARAJAN', 76, 4, 1),
		('SUKAJAYA', 76, 4, 1),
		('SUKANEGARA', 76, 4, 1),
		('WANAYASA', 76, 4, 1),
		('ARGAWANA', 77, 4, 1),
		('BANYUWANGI', 77, 4, 1),
		('KEDUNG SOKA', 77, 4, 1),
		('MANGUNREJA', 77, 4, 1),
		('MARGASARI', 77, 4, 1),
		('PULO AMPEL', 77, 4, 1),
		('PULO PANJANG', 77, 4, 1),
		('SALIRA', 77, 4, 1),
		('SUMURANJA', 77, 4, 1),
		('CIMUNCANG', 78, 4, 1),
		('CIPARE', 78, 4, 1),
		('KAGUNGAN', 78, 4, 1),
		('KALIGANDU', 78, 4, 1),
		('KOTABARU', 78, 4, 1),
		('LONTARBARU', 78, 4, 1),
		('LOPANG', 78, 4, 1),
		('SERANG', 78, 4, 1),
		('SUKAWANA', 78, 4, 1),
		('SUMURPECUNG', 78, 4, 1),
		('TERONDOL', 78, 4, 1),
		('UNYUR', 78, 4, 1),
		('CILOWONG', 79, 4, 1),
		('DRANGONG', 79, 4, 1),
		('KALANG ANYAR', 79, 4, 1),
		('KURANJI', 79, 4, 1),
		('LIALANG', 79, 4, 1),
		('PANCUR', 79, 4, 1),
		('PANGGUNGJATI', 79, 4, 1),
		('SAYAR', 79, 4, 1),
		('SEPANG', 79, 4, 1),
		('TAKTAKAN', 79, 4, 1),
		('TAMANBARU', 79, 4, 1),
		('UMBUL TENGAH', 79, 4, 1),
		('BENDUNG', 80, 4, 1),
		('CERUKCUK', 80, 4, 1),
		('CIBODAS', 80, 4, 1),
		('LEMPUYANG', 80, 4, 1),
		('PEDALEMAN', 80, 4, 1),
		('SIREMEN', 80, 4, 1),
		('SUKAMANAH', 80, 4, 1),
		('TANARA', 80, 4, 1),
		('TENJO AYU', 80, 4, 1),
		('ALANG ALANG', 81, 4, 1),
		('KEBON', 81, 4, 1),
		('KEBUYUTAN', 81, 4, 1),
		('KEMANISAN', 81, 4, 1),
		('LABAN', 81, 4, 1),
		('LONTAR', 81, 4, 1),
		('PONTANG LEGON', 81, 4, 1),
		('PUSER', 81, 4, 1),
		('SAMPARWADI', 81, 4, 1),
		('SUJUNG', 81, 4, 1),
		('SUSUKAN', 81, 4, 1),
		('TENGKURAK', 81, 4, 1),
		('TIRTAYASA', 81, 4, 1),
		('WARGASARA', 81, 4, 1),
		('BOJONG CATANG', 82, 4, 1),
		('BOJONG MENTENG', 82, 4, 1),
		('BOJONG PANDAN', 82, 4, 1),
		('KAMUNING', 82, 4, 1),
		('MALANGGAH', 82, 4, 1),
		('PANCARECANG', 82, 4, 1),
		('PANUNGGULAN', 82, 4, 1),
		('SUKASARI', 82, 4, 1),
		('TUNJUNG TEJA', 82, 4, 1),
		('CIGOONG', 83, 4, 1),
		('KALODORAN / KALODRAN', 83, 4, 1),
		('KEPUREN', 83, 4, 1),
		('KIARA', 83, 4, 1),
		('LEBAKWANGI', 83, 4, 1),
		('NYAPAH', 83, 4, 1),
		('PABUARAN', 83, 4, 1),
		('PAGERAGUNG', 83, 4, 1),
		('PASULUHAN', 83, 4, 1),
		('PENGAMPELAN', 83, 4, 1),
		('PIPITAN', 83, 4, 1),
		('TEGALSARI', 83, 4, 1),
		('TERITIH', 83, 4, 1),
		('WALANTAKA', 83, 4, 1),
		('BINANGUN', 84, 4, 1),
		('COKOPSULANJANA', 84, 4, 1),
		('KEMUNING', 84, 4, 1),
		('MELATI', 84, 4, 1),
		('SAMBILAWANG', 84, 4, 1),
		('SAMPIR', 84, 4, 1),
		('SASAHAN', 84, 4, 1),
		('SUKABARES', 84, 4, 1),
		('SUKADALEM', 84, 4, 1),
		('TELAGA LUHUR', 84, 4, 1),
		('WARINGINKURUNG', 84, 4, 1),
		('BALARAJA', 85, 5, 1),
		('CANGKUDU', 85, 5, 1),
		('GEMBONG', 85, 5, 1),
		('SAGA', 85, 5, 1),
		('SENTUL', 85, 5, 1),
		('SENTUL JAYA', 85, 5, 1),
		('SUKA MURNI', 85, 5, 1),
		('TALAGASARI', 85, 5, 1),
		('TOBAT', 85, 5, 1),
		('BATU CEPER', 86, 5, 1),
		('BATU JAYA', 86, 5, 1),
		('BATU SARI', 86, 5, 1),
		('KEBON BESAR', 86, 5, 1),
		('PORIS GAGA', 86, 5, 1),
		('PORIS GAGA BARU', 86, 5, 1),
		('PORIS JAYA', 86, 5, 1),
		('BELENDUNG', 87, 5, 1),
		('BENDA', 87, 5, 1),
		('JURUMUDI', 87, 5, 1),
		('JURUMUDI BARU', 87, 5, 1),
		('PAJANG', 87, 5, 1),
		('CIBODAS', 88, 5, 1),
		('CIBODAS BARU', 88, 5, 1),
		('CIBODAS SARI', 88, 5, 1),
		('JATIUWUNG', 88, 5, 1),
		('PANUNGGANGAN BARAT', 88, 5, 1),
		('UWUNG JAYA', 88, 5, 1),
		('BITUNG JAYA', 89, 5, 1),
		('BOJONG', 89, 5, 1),
		('BUDI MULYA', 89, 5, 1),
		('BUNDER', 89, 5, 1),
		('CIBADAK', 89, 5, 1),
		('CIKUPA', 89, 5, 1),
		('DUKUH', 89, 5, 1),
		('PASIR GADUNG', 89, 5, 1),
		('PASIR JAYA', 89, 5, 1),
		('SUKA DAMAI', 89, 5, 1),
		('SUKA MULYA', 89, 5, 1),
		('SUKA NAGARA', 89, 5, 1),
		('TALAGA', 89, 5, 1),
		('TALAGA SARI', 89, 5, 1),
		('CIPONDOH', 90, 5, 1),
		('CIPONDOH INDAH', 90, 5, 1),
		('CIPONDOH MAKMUR', 90, 5, 1),
		('GONDRONG', 90, 5, 1),
		('KENANGA', 90, 5, 1),
		('KETAPANG', 90, 5, 1),
		('PETIR', 90, 5, 1),
		('PORIS PLAWAD', 90, 5, 1),
		('PORIS PLAWAD INDAH', 90, 5, 1),
		('PORIS PLAWAD UTARA', 90, 5, 1),
		('CIBOGO', 91, 5, 1),
		('CISAUK', 91, 5, 1),
		('DANGDANG', 91, 5, 1),
		('MEKARWANGI', 91, 5, 1),
		('SAMPORA', 91, 5, 1),
		('SURADITA', 91, 5, 1),
		('BOJONGLOA', 92, 5, 1),
		('CARENANG', 92, 5, 1),
		('CARINGIN', 92, 5, 1),
		('CEMPAKA', 92, 5, 1),
		('CIBUGEL', 92, 5, 1),
		('CISOKA', 92, 5, 1),
		('JEUNG JING', 92, 5, 1),
		('KARANGHARJA', 92, 5, 1),
		('SELAPAJANG', 92, 5, 1),
		('SUKATANI', 92, 5, 1),
		('CIBETOK', 93, 5, 1),
		('CIPAEH', 93, 5, 1),
		('GUNUNG KALER', 93, 5, 1),
		('KANDA WATI', 93, 5, 1),
		('KEDUNG', 93, 5, 1),
		('ONYAM', 93, 5, 1),
		('RANCA GEDE', 93, 5, 1),
		('SIDOKO', 93, 5, 1),
		('TAMIANG', 93, 5, 1),
		('ANCOL PASIR', 94, 5, 1),
		('DARU', 94, 5, 1),
		('JAMBE', 94, 5, 1),
		('KUTRUK', 94, 5, 1),
		('MEKARSARI', 94, 5, 1),
		('PASIR BARAT', 94, 5, 1),
		('RANCABUAYA', 94, 5, 1),
		('SUKA MANAH', 94, 5, 1),
		('TABAN', 94, 5, 1),
		('TIPAR JAYA (TIPARRAYA)', 94, 5, 1),
		('ALAM JAYA', 95, 5, 1),
		('GANDASARI', 95, 5, 1),
		('JATAKE', 95, 5, 1),
		('KERONCONG', 95, 5, 1),
		('MANIS JAYA', 95, 5, 1),
		('PASIR JAYA', 95, 5, 1),
		('CIKANDE', 96, 5, 1),
		('DANG DEUR', 96, 5, 1),
		('JAYANTI', 96, 5, 1),
		('PABUARAN', 96, 5, 1),
		('PANGKAT', 96, 5, 1),
		('PASIR GINTUNG', 96, 5, 1),
		('PASIR MUNCANG', 96, 5, 1),
		('SUMUR BANDUNG', 96, 5, 1),
		('BOJONG JAYA', 97, 5, 1),
		('BUGEL', 97, 5, 1),
		('CIMONE', 97, 5, 1),
		('CIMONE JAYA', 97, 5, 1),
		('GERENDENG', 97, 5, 1),
		('KARAWACI', 97, 5, 1),
		('KARAWACI BARU', 97, 5, 1),
		('KOANG JAYA', 97, 5, 1),
		('MARGA SARI', 97, 5, 1),
		('NAMBO JAYA', 97, 5, 1),
		('NUSA JAYA', 97, 5, 1),
		('PABUARAN', 97, 5, 1),
		('PABUARAN TUMPENG', 97, 5, 1),
		('PASAR BARU', 97, 5, 1),
		('SUKA JADI', 97, 5, 1),
		('SUMUR PANCING (PACING)', 97, 5, 1),
		('BENCONGAN', 98, 5, 1),
		('BENCONGAN INDAH', 98, 5, 1),
		('BOJONG NANGKA', 98, 5, 1),
		('CURUG SANGERANG (CURUG SANGERENG)', 98, 5, 1),
		('KELAPA DUA', 98, 5, 1),
		('PAKULONAN BARAT', 98, 5, 1),
		('BELIMBING', 99, 5, 1),
		('CENGKLONG', 99, 5, 1),
		('DADAP', 99, 5, 1),
		('JATIMULYA', 99, 5, 1),
		('KOSAMBI BARAT', 99, 5, 1),
		('KOSAMBI TIMUR', 99, 5, 1),
		('RAWA BURUNG', 99, 5, 1),
		('RAWA RENGAS', 99, 5, 1),
		('SALEMBARAN JATI', 99, 5, 1),
		('SALEMBARAN RAYA', 99, 5, 1),
		('JENGKOL', 100, 5, 1),
		('KEMUNING', 100, 5, 1),
		('KOPER', 100, 5, 1),
		('KRESEK', 100, 5, 1),
		('PASIR AMPO', 100, 5, 1),
		('PATRA SANA', 100, 5, 1),
		('RANCA ILAT', 100, 5, 1),
		('RENGED', 100, 5, 1),
		('TALOK', 100, 5, 1),
		('BAKUNG', 101, 5, 1),
		('BLUKBUK', 101, 5, 1),
		('CIRUMPAK', 101, 5, 1),
		('KRONJO', 101, 5, 1),
		('MUNCUNG', 101, 5, 1),
		('PAGEDANGAN ILIR', 101, 5, 1),
		('PAGEDANGAN UDIK', 101, 5, 1),
		('PAGENJAHAN', 101, 5, 1),
		('PASILIAN', 101, 5, 1),
		('PASIR', 101, 5, 1),
		('BABAKAN', 102, 5, 1),
		('BABAT', 102, 5, 1),
		('BOJONG KAMAL', 102, 5, 1),
		('CARINGIN', 102, 5, 1),
		('CIANGIR', 102, 5, 1),
		('CIRARAB', 102, 5, 1),
		('KAMUNING (KEMUNING)', 102, 5, 1),
		('LEGOK', 102, 5, 1),
		('PALA SARI', 102, 5, 1),
		('RANCAGONG', 102, 5, 1),
		('SERDANG WETAN', 102, 5, 1),
		('BANYU ASIH', 103, 5, 1),
		('GUNUNG SARI', 103, 5, 1),
		('JATI WARINGIN', 103, 5, 1),
		('KEDUNG DALEM', 103, 5, 1),
		('KETAPANG', 103, 5, 1),
		('MARGA MULYA', 103, 5, 1),
		('MAUK BARAT', 103, 5, 1),
		('MAUK TIMUR', 103, 5, 1),
		('SASAK', 103, 5, 1),
		('TANJUNG ANOM', 103, 5, 1),
		('TEGAL KUNIR KIDUL', 103, 5, 1),
		('TEGAL KUNIR LOR', 103, 5, 1),
		('CIJERUK', 104, 5, 1),
		('GANDA RIA', 104, 5, 1),
		('JENGGOT', 104, 5, 1),
		('KEDAUNG', 104, 5, 1),
		('KLUTUK', 104, 5, 1),
		('KOSAMBI DALAM', 104, 5, 1),
		('MEKAR BARU', 104, 5, 1),
		('WALIWIS', 104, 5, 1),
		('KARANG ANYAR', 105, 5, 1),
		('KARANG SARI', 105, 5, 1),
		('KEDAUNG BARU', 105, 5, 1),
		('KEDAUNG WETAN', 105, 5, 1),
		('MEKAR SARI', 105, 5, 1),
		('NEGLASARI', 105, 5, 1),
		('SELAPAJANG JAYA', 105, 5, 1),
		('CICALENGKA', 106, 5, 1),
		('CIHUNI', 106, 5, 1),
		('CIJANTRA', 106, 5, 1),
		('JATAKE', 106, 5, 1),
		('KADU SIRUNG', 106, 5, 1),
		('KARANG TENGAH', 106, 5, 1),
		('LENGKONG KULON', 106, 5, 1),
		('MALANG NENGAH', 106, 5, 1),
		('MEDANG', 106, 5, 1),
		('PAGEDANGAN', 106, 5, 1),
		('SITU GADUNG', 106, 5, 1),
		('BUARAN BAMBU', 107, 5, 1),
		('BUARAN MANGGA', 107, 5, 1),
		('BUNISARI (BONASARI)', 107, 5, 1),
		('GAGA', 107, 5, 1),
		('KALIBARU', 107, 5, 1),
		('KIARA PAYUNG', 107, 5, 1),
		('KOHOD', 107, 5, 1),
		('KRAMAT', 107, 5, 1),
		('LAKSANA', 107, 5, 1),
		('PAKU ALAM', 107, 5, 1),
		('PAKUHAJI', 107, 5, 1),
		('RAWA BONI', 107, 5, 1),
		('SUKAWALI', 107, 5, 1),
		('SURYA BAHARI', 107, 5, 1),
		('CIAKAR', 108, 5, 1),
		('MEKAR BAKTI', 108, 5, 1),
		('MEKAR JAYA', 108, 5, 1),
		('PANONGAN', 108, 5, 1),
		('PEUSAR', 108, 5, 1),
		('RANCA IYUH', 108, 5, 1),
		('RANCA KALAPA', 108, 5, 1),
		('SERDANG KULON', 108, 5, 1),
		('GELAM JAYA', 109, 5, 1),
		('KUTA BARU', 109, 5, 1),
		('KUTA BUMI', 109, 5, 1),
		('KUTA JAYA', 109, 5, 1),
		('PANGADEGAN', 109, 5, 1),
		('PASAR KEMIS', 109, 5, 1),
		('SINDANG SARI', 109, 5, 1),
		('SUKAASIH', 109, 5, 1),
		('SUKAMANTRI', 109, 5, 1),
		('GEBANG RAYA', 110, 5, 1),
		('GEMBOR', 110, 5, 1),
		('PERIUK', 110, 5, 1),
		('PERIUK JAYA', 110, 5, 1),
		('SANGIANG JAYA', 110, 5, 1),
		('CIPETE', 111, 5, 1),
		('KUNCIRAN', 111, 5, 1),
		('KUNCIRAN INDAH', 111, 5, 1),
		('KUNCIRAN JAYA', 111, 5, 1),
		('NEROGTOG', 111, 5, 1),
		('PAKOJAN', 111, 5, 1),
		('PANUNGGANGAN', 111, 5, 1),
		('PANUNGGANGAN TIMUR', 111, 5, 1),
		('PANUNGGANGAN UTARA', 111, 5, 1),
		('PINANG', 111, 5, 1),
		('SUDIMARA PINANG', 111, 5, 1),
		('DAON', 112, 5, 1),
		('JAMBU KARYA', 112, 5, 1),
		('LEMBANG SARI', 112, 5, 1),
		('MEKARSARI', 112, 5, 1),
		('PANGARENGAN', 112, 5, 1),
		('RAJEG', 112, 5, 1),
		('RAJEGMULYA', 112, 5, 1),
		('RANCA BANGO', 112, 5, 1),
		('SUKA MANAH', 112, 5, 1),
		('SUKA SARI', 112, 5, 1),
		('SUKA TANI', 112, 5, 1),
		('TANJAKAN', 112, 5, 1),
		('TANJAKAN MEKAR', 112, 5, 1),
		('KARET', 113, 5, 1),
		('KAYU AGUNG', 113, 5, 1),
		('KAYU BONGKOK', 113, 5, 1),
		('MEKAR JAYA', 113, 5, 1),
		('PISANGAN JAYA', 113, 5, 1),
		('PONDOK JAYA', 113, 5, 1),
		('SARAKAN', 113, 5, 1),
		('SEPATAN', 113, 5, 1),
		('GEMPOL SARI', 114, 5, 1),
		('JATI MULYA', 114, 5, 1),
		('KAMPUNG KELOR', 114, 5, 1),
		('KEDAUNG BARAT', 114, 5, 1),
		('LEBAK WANGI', 114, 5, 1),
		('PONDOK KELOR', 114, 5, 1),
		('SANGIANG', 114, 5, 1),
		('TANAH MERAH', 114, 5, 1),
		('BADAK ANOM', 115, 5, 1),
		('SINDANG ASIH', 115, 5, 1),
		('SINDANG JAYA', 115, 5, 1),
		('SINDANG PANON', 115, 5, 1),
		('SINDANG SONO', 115, 5, 1),
		('SUKA HARJA', 115, 5, 1),
		('WANA KERTA', 115, 5, 1),
		('CIKAREO', 116, 5, 1),
		('CIKASUNGKA', 116, 5, 1),
		('CIKUYA', 116, 5, 1),
		('CIREUNDEU', 116, 5, 1),
		('MUNJUL', 116, 5, 1),
		('PASANGGRAHAN', 116, 5, 1),
		('SOLEAR', 116, 5, 1),
		('BUARAN JATI', 117, 5, 1),
		('GINTUNG', 117, 5, 1),
		('KARANG SERANG', 117, 5, 1),
		('KOSAMBI', 117, 5, 1),
		('MEKAR KONDANG', 117, 5, 1),
		('PEKAYON', 117, 5, 1),
		('RAWA KIDANG', 117, 5, 1),
		('SUKADIRI', 117, 5, 1),
		('BENDA', 118, 5, 1),
		('BUNAR', 118, 5, 1),
		('BUNI AYU', 118, 5, 1),
		('KALI ASIN', 118, 5, 1),
		('KUBANG', 118, 5, 1),
		('MERAK', 118, 5, 1),
		('PARAHU', 118, 5, 1),
		('SUKA MULYA', 118, 5, 1),
		('BABAKAN', 119, 5, 1),
		('BUARAN INDAH', 119, 5, 1),
		('CIKOKOL', 119, 5, 1),
		('KELAPA INDAH', 119, 5, 1),
		('SUKA ASIH', 119, 5, 1),
		('SUKARASA', 119, 5, 1),
		('SUKASARI', 119, 5, 1),
		('TANAH TINGGI', 119, 5, 1),
		('BABAKAN ASEM', 120, 5, 1),
		('BOJONG RENGED', 120, 5, 1),
		('KAMPUNG BESAR', 120, 5, 1),
		('KAMPUNG MELAYU BARAT', 120, 5, 1),
		('KAMPUNG MELAYU TIMUR', 120, 5, 1),
		('KEBON CAU', 120, 5, 1),
		('LEMO', 120, 5, 1),
		('MUARA', 120, 5, 1),
		('PANGKALAN', 120, 5, 1),
		('TANJUNG BURUNG', 120, 5, 1),
		('TANJUNG PASIR', 120, 5, 1),
		('TEGAL ANGUS', 120, 5, 1),
		('TELUK NAGA', 120, 5, 1),
		('BANTAR PANJANG', 121, 5, 1),
		('CILELES', 121, 5, 1),
		('CISEREH', 121, 5, 1),
		('KADU AGUNG', 121, 5, 1),
		('MARGA SARI', 121, 5, 1),
		('MATA GARA', 121, 5, 1),
		('PASIR BOLANG', 121, 5, 1),
		('PASIR NANGKA', 121, 5, 1),
		('PEMATANG', 121, 5, 1),
		('PETE', 121, 5, 1),
		('SODONG', 121, 5, 1),
		('TAPOS', 121, 5, 1),
		('TEGALSARI', 121, 5, 1),
		('TIGARAKSA', 121, 5, 1),
		('CIPAYUNG', 122, 6, 1),
		('CIPUTAT', 122, 6, 1),
		('JOMBANG', 122, 6, 1),
		('SARUA (SERUA)', 122, 6, 1),
		('SARUA INDAH', 122, 6, 1),
		('SAWAH BARU', 122, 6, 1),
		('SAWAH LAMA', 122, 6, 1),
		('CEMPAKA PUTIH', 123, 6, 1),
		('CIREUNDEU', 123, 6, 1),
		('PISANGAN', 123, 6, 1),
		('PONDOK RANJI', 123, 6, 1),
		('REMPOA', 123, 6, 1),
		('RENGAS', 123, 6, 1),
		('BAMBU APUS', 124, 6, 1),
		('BENDA BARU', 124, 6, 1),
		('KEDAUNG', 124, 6, 1),
		('PAMULANG BARAT', 124, 6, 1),
		('PAMULANG TIMUR', 124, 6, 1),
		('PONDOK BENDA', 124, 6, 1),
		('PONDOK CABE ILIR', 124, 6, 1),
		('PONDOK CABE UDIK', 124, 6, 1),
		('JURANG MANGU BARAT', 125, 6, 1),
		('JURANG MANGU TIMUR', 125, 6, 1),
		('PERIGI BARU', 125, 6, 1),
		('PERIGI LAMA', 125, 6, 1),
		('PONDOK AREN', 125, 6, 1),
		('PONDOK BETUNG', 125, 6, 1),
		('PONDOK JAYA', 125, 6, 1),
		('PONDOK KACANG BARAT', 125, 6, 1),
		('PONDOK KACANG TIMUR', 125, 6, 1),
		('PONDOK KARYA', 125, 6, 1),
		('PONDOK PUCUNG', 125, 6, 1),
		('BUARAN', 126, 6, 1),
		('CIATER', 126, 6, 1),
		('CILENGGANG', 126, 6, 1),
		('LENGKONG GUDANG', 126, 6, 1),
		('LENGKONG GUDANG TIMUR', 126, 6, 1),
		('LENGKONG WETAN', 126, 6, 1),
		('RAWA BUNTU', 126, 6, 1),
		('RAWA MEKAR JAYA', 126, 6, 1),
		('SERPONG', 126, 6, 1),
		('JELUPANG', 127, 6, 1),
		('LENGKONG KARYA', 127, 6, 1),
		('PAKU JAYA', 127, 6, 1),
		('PAKUALAM', 127, 6, 1),
		('PAKULONAN', 127, 6, 1),
		('PONDOK JAGUNG', 127, 6, 1),
		('PONDOK JAGUNG TIMUR', 127, 6, 1)
	`)
	if err != nil {
		log.Fatal(err)
	}
}
