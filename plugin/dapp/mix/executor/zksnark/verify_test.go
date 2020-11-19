// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zksnark

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/33cn/plugin/plugin/dapp/mix/types"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	verifyKey := "1f8b08000000000000ff6c517b4c937714fdddf6536116149022a8a3bc9c1544a84a1086c8a3b85a278ac1c7b0bc6ba90386b41511910928e5fdb43a3543610c5164109c5a37314689c4218e1158650a5b47c4884e05991714bea51fe292657f9dfc7ef7dc7bceb9973d93b003e8436c00ce5669b27c57aa3c512696a6029d49800210029d4d80b54e00748101dd812e21c0d9a48a8e97c7881293544a05d01a42881d9dc506604bdd05531dc00a7003fa08015680bb010959421f6603b0a41ecc371b58fe6e40ab09b0fcdddfa1c08084f0e81c862898fa06969f1bd07904587eee0624c48ace058059c27869823451c9d466821121523a9f0de0ad5026ab6294bc34debaa88484a88d52192f3a51b0cac3759dc06fd72e79a2d49b17288d57fe4f81973e95128ca63b812e2260344d37bc0871a40bd90046d34d0c8505b09d710b3ba63208e86236c0a2f756c4bc30c93b31f77fc54a08008018e80a42882d5dce0298fb5f9ea108749941b794d17dff5f36ad9b37a56b580d973eca02300a932894c9f24499e136c0216492dd4a671ab4286c53bcacf11bbf7402b5c8b1bf38b157895983a14badf436e9c8d56dc9dfe3dad50814a6d7465af46d1d6b46fd8f631e51dd8a6eb475a914b4da44f4e0acc7cd8e23034b55c4308d2babfdeb79a64316aa3afbefa83fffec28b6b9253b4ddcb694e3f23f4bbe7f312af1040a858bf5837e5b9e05a1282d29b377c53e37bcdddfbda2ddab2d16cd35f10feecf7bb29999569bab1f14cb1a8250f1ddf15e9de3ad37b87afdd5868ad50331b85055de2adb861140a1f596a37d73bcc79d715fc3b5039dbf695c50af9eedefebba488d0b9d87d7ea0fdeb22184c99a127c5912ee59c147a37dc4f711e78c1faa95675b3571bbbdd0aee297932e21c72e00858f46474da21b440118776a81e92585731c66151c08af3a625f8f1f5856342a0e167833ee62ed33124b7d9bae628d71dcf84513761e0ee51e5e3522d18ee1525d7fd34a21270b282c7ec9bd32af3af62966fadc0968896d7f831da27a0f9c61721a3979e79dd76ed85dcb4c8b793b5278236d38199f8d17a55766478e608ac7b9a61b55430a9cbb725b7647df09c31d2e3717efecfd4317832b1af4c1d5b2949ff1a979a5981f91148a663ec3f9658157761026ac617b7dfc21e31f469e63ca3703ed3983ca5ff1edf916cfbfcd343bd03498fbed3dd18510a0f06e58754fd7ecb4e5c8ad5bbfeac927ea022c4f900f08db8aacd179a2382092bb7f82f1b75db24cad0d71f0c20f87261cc336a51aa36bf3ab84cc25cd34f2e77c35daf8f6740350d885d228301eeec19b35053e2b2d861ad1f361e30bfe0d311f4dc7b5f7be36b33ff4ee167517ebfcc54109f7b163f9c2f98bae0b63519dbdc6f59497d923345fd3e1131aedd70214d2faf0979ffaabcab0c3a92ff4ead90336585d3e1974f96cc831b4a56d23afb53d2c63dcd5db25f10ba1528fbd5d15ca8e415e2796f0ac449a1289095a65e8acf78449cf0185339b1c6c3c4c1d6c5113d9fe65a05a721c4fd66ff5dfe9fdc065cce3f8e6ebc67539ccf2584061a9aea78aeff97b2796b6cdcfe8945adc43575fdd82bb65af65386372e7fec843a79d80c2d541813f6d0a1bbf8939cbae770657a92cf1c45e4e8b57ef17f968fdf1f90bb1c9dc3b04288cef1739693ff2a731f3f5e35ccbf4c200bcbbf95542c46bef545cbc41606773b3520b1456172936a64cca2dd06e8e437dcb68773876d5f4cfc88834f34273f119b714edc33442806514bc5118b14d14228454f24f000000ffffd7ca51e5ae050000"
	pubInput := "7b2279223a223335227d0a"
	proof := "1f8b08000000000000ff62666360d6f9dfc8ccc8c81a50949f9fc6f8bf8981919991c9b188f17f0b0323b377513198c1e454ccf8bf83818141e57f3333232387bba1635a5a665e2a4492913182f17f1b0323632488626010ffdfcac8c8c8ee9a939a9b9a5702966263e400696e076b368269ee8069ee8268ee62606050f8dfc9ccc8c8946a04166562647234009bc0e4680831fd1fa3edff264646961ff2dac6726f5e0770fdd87d50ae413cc4a0fe87884ade7de7bfdd0b7eb0cdb27dd07755c99691e5879fb6c43f4b39a7a93fda32c5ae1c3d3cf5d28f79f35d4d3e955999fee099f6694e45d40e65069069468dd1b6cb04f3847ffc4ff8e2753ff3a2e48f77f24b59ff96d69df921b371399fffd5f5bf19597e249ebdf6e77fd444f71f079466ad774ef3abfe71a07bdd077ed695e13fe46ebe9e6ba4bb51888111645cd4abcfee7d2a26277fdc0a3d533b695a8bcc0faffdbcb71c4d3fb8fd50ac280dd7b253bbc2c8f26375a28ed446a98fb77f1cf8735e594cfbefff1feebdbbd25317759ef9211221d33433eb7c33d87147ef898b06a44ebffaa3227cdd5fddedaf56ff5860397f86c0559f673f647be744dc5fa6f49891e5475335e7a317d32df87e1c2f3f7352ef5339cb8f2dd11ffeff09ebbcf7437ff6e6873f8fa96c64606000040000ffff023acbaff4010000"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)

}

func TestDeposit(t *testing.T) {
	verifyKey := "1f8b08000000000000ff6c517d34d47b1efe7ee6f7c35ccdb5dcbde18a8cb7dd417b9941e38eba1bc2e25cee76dd625dad89c1d4cc2466dc1c160945de95e965cb59b2c7944a75923a3615dae294262c6de985517adfa87c52faed9991f69c3df7afe77c3f2fcff379be0f654ca840660b05c0592d499326654a15c9e1924c60f209d00041c014106085088029d5231f984a029cef55eb64d2845045aa4a990e4c2d21c481d94a015012be606e0358819ec014116005f2f548088f29a4005892a5863205ac004f60b6136005f03fa2408f8470996d8641c15c1958fe9ec0941060f9f3f5488815530c00264132895ca2501a7ac6c02644c2eca000fcd29569aa0425378b1b2296cbc5119264ee3a85c067e9d72102ffa424a942e2c75d2991297fa1c1fdcb9c4b60cf6f02534e803d3fae7f11e2cc945100ecf925c3080b20da702dc4cc791030151480dda753c2b9b1711fc5f8ff13ab240000e1c0ec2484d833352c00f3ff9fd33781a9d6eb5619743fd5abe7754be674f55f63c9ec6201b063e3d295695245b23e1be010f281e632f97a2d1ad3c03f3da2f46d17169fb22d9fea68ba8f75f195897745d4185a9c297399e8715e0534be6d4996699f0ff6e088139a1909a7aab17d5f7b996679712eba059ded7a1063ed45f46cd75aad7eb7e9e9a3dbe81998b6d63e7398416560618250f38519da7032f277be9cce061acf560f468de78c6b30c9f2c66dc1f9f119ec160e3f3df10fbb52e49c883325cace21031bfd61e554f2e5e976bcdf4bed4cd40832d0abb6312ccf66b315c2af374878da4b8f81c682d23ddb467beaed30eb492fe6fc14528effdecd5fe8b54dd98156d7ffee15f0e0d9578418bc8e99ccae9f6db9f4012fddc97532af5921c53e5b9b7dfd4f223390fbf902381499de0b340e05fcd521d147148de6a5233029e7d9a0769c37f32e3eeb32dadd97731b2f46ce79b512b9ebfa52dbf858618f5322edcd0aec8d78d6601473611c7f23905e7858d06a0134b6df685eb46cafe94bb434ea0efcd7e0678b3165ff1bd3d4f257416f4dd5b7fe16d23c6920130be927392e65228cd05dddff65a4ff0c76ae5547d057a607d055d53a3030dc6f0b34bedf6527dbe119e980df0cae92a6a39020631b75bbbead448dac3e971a595e858a18bc028d1ab578f0ace6a91c3d2658a5e73ac2697c67c5733c1e9cdb8fb613acc3a35fb8ca80c689510b9277ddf121baa67554769ff3a847773757379163d36eb4d6960bb6e72dfa6a2e58f5fb5f79d5de1160fbf500e36c75e66bf4ac4ba3fbb4b273685d95bbc2d72dae10681cd95bb65c3d63d78b192d939f9b305b7af1e754f72529de2fcc90b791d7a0733ee3f7318a49d5c91fffb420271a2baadf38e83cfa8fe2a1850b527f3f921285bc1a6ffb7b4e17478146e1e203d6d5c6c7aa90fa63dd6707c25e68708833db2db6187f8eb6dbabbe195c16e363b86eaa257c83f99ed55b51dd9709be9c3b0ec8715cffdb0197ab42b415999cfaaee05103d018a67b10e234b1c11f83c7629ba71b6cf6a2c2a72c3969b3ae14174e67c72ac23c36197e8f021ac3f9fbda4e666d3e885ea7357dbbc7a62af15e3ebb2bec4ded4b74097eb54c788ffd03d0b8b6af68f1a354b5007d971ed659d5dd1dc4e1fa75514bfabd8fa06570e8ea235f0efd4cf45144b71d9fb9e52dc5c6a6c84db35d2543a83de8ddbcfceb6b1968b9c2c99d75a6c91d685cfff8bb42d3d1e2839864326dbf48ac15e1cd1ffd4ace17294ea34782597d66e7ec313dd99503df7e5fd9f63a0a7f8851173defd1e460e3acb1aa92d34aa17be712e1d1672bd7008d71bed3ba3a2f6706c3d784c627fd33fe14d698cb85d951908836ff09a47fba69f42d2140b1232382febc26745590b158be51a550b2151b95923f88d353c87f030000ffffeebab67209060000"
	pubInput := "7b22616d6f756e74223a223238323432303438222c226e6f746548617368223a223136333038373933333937303234363632383332303634353233383932343138393038313435393030383636353731353234313234303933353337313939303335383038353530323535363439227d"
	proof := "1f8b08000000000000ff62666360d6f9dfc8ccc8c81a50949f9fc6f8bf8981919991c9b188f17f0b0323b377513198c1e454ccf8bf83818141e57f3333232387bba1635a5a665e2a4492913182f17f1b0323632488626010ffdfcac8c8c8ee9a939a9b9a5702966263e400696e076b368269ee8069ee8268ee62606050f8dfc9ccc8c8946a04166562647234009bc0e4680831fd1fa3edff264646961f8a8b3aa3e4d764fcfab1ab72fab48e09b6453f9a26ac8b173a73ecd40fa5c6268bdb3bafdc6464f971ab2a30fb6bcec3693f5efef5ea5a9139efe30fce6bb53e4b6fdcbbf943a4f06b5a93dd2c79069069be4297b6f331096cfc31ad6b41d186058bd6fee0eaaedd71adf247d30fc1eb66fdb7a4441733b2fc98d1bc75ae44eb1f831fbba76505182ef894f2e3dd9189b16e8ca25b7fc86e7bbae1e6c2a5120c8c20e366589fec28d4f3daf8e398fdd176ad7d56393f7c1dd6ce9ed93451ef078f0f43d28e29160e8c2c3fa6274c177839df4ce7c7cf1fd354bbee08bfff319f3ffacdfe97461b7f2836c69e4ad15262013bcefb0e7b69dd8dad353fbceecf379d9eb477d18fe0ab2f1e25722b7dfd2175a22ea14a38da8291e5c7969baa8d1eeb369afeb0f2cd168ac899b7f107df7989c2d4d7cb797e485c71bed89e71888b818101100000ffff7c81eeadf4010000"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)
}

func TestAuthorize(t *testing.T) {
	verifyKey := "1f8b08000000000000ff6c527b5493e7197f9f2f1f907211b52da88004a5520b38881005842306449688119058016dc40f890782cba531b371155141299030db5a2fb00896138ebaaed551f10647451c502f5b8528322f2005e7c28cf4d13abe9d2fa83b67677f3de7b9fd2ecffbf25c094fcc6ee7017866316a65be41a9da28610cc09610a00192812d2540a50881ade06224b0d5043c65baf585cabc54d5669d5603ece784902076070f80c7440a273680124700bb8b00258ee42221efb33b79001423729679402d8900b68c00b524f2551472911001bbdb39289c2803951801ec1e025462241709f165cb01c02db9902962545a67cf15f88430ec5e1e409c46abd6e569055b05298aa222451ab351b05e258c16cd4b1126e6e72b554c9c208929d4fe9f86c038e112f8af3781ad24c07f3dce658404b39ff100f8af979c2314c06aa75af870c28390ade20104bc91221164e7be228bfc2f593501009000fb7b4248205b43014cfedf39ae09ac99e3353979dfd4cdaf79f74cf072a7f161f75100fcec5c8d56ad546de4de063c0919a71fb1251c178dfa9d496b17f85422feea9f2f87a423d78ee0131cada950e856e32cdbe0899ba43b1b68dc2ed0777e50d168c3ea9ebcf5d316673cc60737121a0f7d622fc7b7fd36978b6cb763088796a2dcb622d7127e1a130b1766192b43fbb038b0e55c60e7cc0c14c47f11e558eb770e687c70fdfef4294cbc2b3e3a643a26fde9c20ebc77f8f89d63252d85e872f140cb8b539fa639d12272b26f9fcf84f3d814ee75bda920211b57e66def6e1c1e5987f342b242e69e765d06347e79326257666054077a94edfaca7269bc1e3bdf7d7c5754734186fc7b77add4d8f1a38438bdee35cb67278acfe4a12867eccacdaa2d7614c42d1ab81bd73b86ef247dfc83ebf002ceeba98bf57ef3670aba31b6cad15cf08773e1385a7f2333738f84f73c8cd41b668d7ce714b7eaada76db3a9e487f889696756e505aa0b459eb63b1efe9521f8bebbc85a26cf8a061a0f1cbedcd471cad28eb732665fea8125bfc1291a7bfba276ef4e9c91e86835e71ccc77a2f56e6c99f77371ff4a3cef6fbd79ea454f2bbe285efec7e21f8c160cd595382e8febfb81c6f0a943a962993400175a83bb46725d82b1a9779c9f93b4b6093d53676d98dc70f85be2f40a347e6f92a50dd448d7e2e57befc6fcf551a01c4fd6cf3adb31c36cc549a505ad0f97597381c6c72d065bc8826fe2f11dbb7ca14bd89c2eb42a1c5b64d7c73230587ac0ad30be7995531f4f6d4ce8f5dd4a63decae82c59db357f3cfbe3f461f99abffc0367b89fd937685ffe21d058b7dcf8a4bbb361182b6f04d9c99f9b776287bbe24fefa5f55ec1a0a26f4ea45f75bbf8ea29145dd33d1e565f52e2ee9325b21493ef5114596f0ffb746cbf8ae14679dfc269095381c68fff3e3e2762519b1d379d751c0fd81df1111eee130fea431c2538a9c2d72cb9ed93ea54b76f798f79ffad9c7e8c9e13877356ef388baaab43dac1c5fac7e8712bd8bbfa8a3e0468dc90ee7524aa94adc3a5eadf5d3fe635b71c1ff499032d096fdfc1b9512f2b92c66adb9cd773e1ae3734ee62e93fe3863da5f79ffefc9de43e1e4a88d9375a579689ae457b3b3a7d1a1440637677fb8dbf79493d9e37162b6bd5750f51d2b22db55c3c3080bcd2ea5d312bae7c4db80feafe6868748db706db5b12f67bffb2ea38baf83ea968adf59a86be4be5dbdcd28f9880465310f3eb07678c7de811363f11198b0afb0e289bbd03361dc2c94c1a6f77d7cc180e2ce958bc7bd8d218094efd9effe397d44f4da83404fa2db5515af4f874bf4f0d1ddb0334a66fea0af50bd46dc1f34fa56deccaa849b83f6acd36fdd69113e81ffdcbf3a7b66b991c58c8ec3ab1dcddff053e118a8c2f6d124017293520b4181c181a64fb5a186df80c68d4b8d99f3d93d736e0564633c5a7222c0a2fafb96628d3cd5f87a1fa291f7db0baee280736bf2a3f6bf3e20d295895dcf09eebe703528c8d9dfe6dffbf5be391cafcead6a4dac13aa0513cbaa3fa5f734d04edcf624f37bfe5db885fc8b22e894ff72930e0e041db90b1d09f1070e1af484b5e274f4d4ff652e8b405c56ae56f99650a4d81f79b4ca65b2f610c53dfe4199b19d5066ec453ab6698f4e2622d9790ff040000ffff593c5ebfd0060000"
	pubInput := "7b22617574686f72697a6548617368223a2231323637383235343336393337373636323339363330333430333333333439363835333230393237323536393638353931303536333733313235393436353833313834353438333535303730222c22617574686f72697a655075624b6579223a223133353139383833323637313431323531383731353237313032313033393939323035313739373134343836353138353033383835393039393438313932333634373732393737363631353833222c22617574686f72697a655370656e6448617368223a223134343638353132333635343338363133303436303238323831353838363631333531343335343736313638363130393334313635353437393030343733363039313937373833353437363633222c2274726565526f6f7448617368223a223130353331333231363134393930373937303334393231323832353835363631383639363134353536343837303536393531343835323635333230343634393236363330343939333431333130227d"
	proof := "1f8b08000000000000ff62666360d6f9dfc8ccc8c81a50949f9fc6f8bf8981919991c9b188f17f0b0323b377513198c1e454ccf8bf83818141e57f3333232387bba1635a5a665e2a4492913182f17f1b0323632488626010ffdfcac8c8c8ee9a939a9b9a5702966263e400696e076b368269ee8069ee8268ee62606050f8dfc9ccc8c8946a04166562647234009bc0e4680831fd1fa3edff264646961fdef959ec5357f7bbfdd82211cf3b3dd643e287baa480fcbdece6833f38bcb63aff76590d52f526c657a7e69dd2d91ffe356b42446f9d2eff51c8d1a7685d6f74e207cbeee7fe270dbd7b1840ea44be4c4bceae5dbeefc7777e17b9ed51f6963faa7e6c37b8d331ebe30fdeb29cbefd22e5a98c2c3f725f6c539c20d12af42364d39683712fe6b9fdf8b35f2b523ba82de007bfd0bf438e6d7f53191841c6fd9cc796ff25ce7ff98fc4c486339b7ce4e6fc70b35feab79973d19f1fbc938d224bcd9e5a30b2fc58773e3f4169da6d811fe758c396efaf950dfeb1e46de244b132862d3f44cfcdd4dbbf62d16fb0e38e4e5d9653ef1fcaf363dea3b01d8d7d6f96fd98bc7d83353beba9f41ffcdb9e4edd32fb813c23cb8fcf87debc5956e8c6fe631f9bfdb7509fbe2d3f762c327b9c60e97fff87ee7e17e6537cd9af19181800010000ffffc892b7eff4010000"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)
}

func TestTransferInput(t *testing.T) {
	verifyKey := "1f8b08000000000000ff6c517b50937716fddde433044c512c0645a2a1c511110422a0426193208d54090c6e2c50740df801614362432205f109a2d0e521b0209465b174c5062bd6078b150505953a822bcb5b8b60712ba0b558d4abddf2ed246877c6e95f677ef79e7bceb9bfcbe6107620b3970dc0db48eb5471a92a4dfc3a3a15987d04288020603209b0642260fe62424f600a08f0c20c316a556cb0669b419f0c4c0921c491c96003b0694fd1f404b0023d80c922c00af4342121cecc7e36008bf63197d9c0927a007390004beaf90a45262444c81c301345d36560493c80c921c092789a90103b261b002c82d47412add19b7b1ce01242339fb201fc92f53a43ac5eb8432853262529e574bc304623f2f6592e1349e2e2541ada4fb88656eb7fa721dc39bd25705f4f02934780fb9a6e7a11e2c4e4b201b8af87cc14164084392d444eef2062f2d90082dfa2ac137eb4e99599e7ffcd0a0800c03a608a09218b982216c0ec3779a626308526df4366dfdfea85af7d73a67d4d5fc367feca02e07eb42959af5369e24db7011e215333c4cc3e931785b7169c4948f3cb6a47ca782e40d071ed3bfce5acf6e6b6726281c20b55be56de6f2f040a837b0abf51334f3ec7e81fb8fb4f2dbcfd3e6a7b63eb9724e04c5c263eaf2af3bffb1331a959c8df39f8b8fa9bc3e893fda3c26a4da602ef14fe337f68e4328dd69b6128f1b86a2e5028ec3ff0aebaead3010c3ef973db6151cb0dccd591b29e9800192ebaffc1ad6a87f600b35ab647c5171de31d0ee89a613d73f927bc593858d67ceba2ddae0c744890055b77c7e880c26687c9e8a2f97951f8d07ef98daefa43b9d8b3a5d2caada25b814eb87bcb9d0ab74242ccbb3e336edb9d1bdb7a1ffff047ead7f40767a538a48a5edd6e252f41574be3bfc7fdf35a81c217faacaf3e9f61770bdb6a576c48ffa1d912672d73b79df78f90ed68977df8fd63173a7dcde99c72aa643f1a32a7f08e7dd49c8a2bad1ccc1ab8f9d0de4b9689b645aaa79b4bbbfd80c2f03d7ebc55de8e2df8ee9c40754ba2ed7bb84977d1ddf28482838ed743359d8e799366b5c775612b3b17dbff1d8bbe3e73b7afadc1038f2a8ee7ba56baee429b6f5d7b3f18f82c0e282cd9f1507e2d25742e3a288bf8713cdb1ab458fbedf9a7069b429c3d5c9f1a70c97d949897050aaf8d1d89b8bec73b1ae50dc3820b47f7c5a2a35778444dfb090172bdde3976ce8e160285df19ef9d3d6318d6a2b83faf31b97a8714fbf887cb1d3c6539386f38f845832032c79cafd2d7e9ca3a55a3069b2c8676ad38f9ac0bdb1a64474abcb3f7a34dc7d1f5d5ba9a34a0f0a0cf53fad2b3cb93981a16b6a03afaf965ecf6e53cec8f193a88dcba2659a5d7a9de57b7e813cbd3234ae9462c908e0def6f9bd7880bf5e9b53b7f7215a35032759b67943403852e92915ceb409f1cac79501014ecda5780c58ac19b7dad8fc468b752d1fbf851949b395dd6de557ad6fc15ebb1b2d83215cfaede8e290fc2537dcb036bd02e6009b5840ab90d146a43fdd7688fde8947273f234abfccfd1803bf14090649f74bb44e96eb33c6461bcdbfc7010a236d260e49baf1676cec7951903df9962d864c38e66b473cf371d6a8e090c4cdc294af629cbd72e0fba56d7860e979a7da8d757d982277bf5279ac588ad657eeb6f88f5f1c274061fd848bec0b9f45fd5872ca4829949d8158c537bc2d6e9ab11d9d43f88c5368f873a0f0c9fd7d615d70691ca39dad9a06679232bcd76ad0bd572e9d8df332ae6fcf4b49093789f9dc4de989d6ffe714ae687f54da9573750afd57c63edf6ab756879cde2dab3e4912370285a7f583a322b6cb527cf175dbd8c2bcba34f43bcddd15e5fffd5738dbddbd491993b9da24f6b14b7fdd4926f12a2636ef7c29fff5aa1cf9670489a58b1d7d110c0e36354f6b4f0385e7f38bff1b545c3c89234aa5604cb3958fd1c72f15a438d7bec4c55bebbb6a9a241a93586279cbee23b739cb70fc17af668b51ef673835f16188fbe8d20074cab31f3d975b5d01141e1f7d3437c77f6f052e689feb9cae2c146261cbbfdae2e6df588e2ec6fed8b83f57959ac422ee2d227c17cb789446bab5ff6dfe093ee29ee121c98471082d1ac495563b8337038569d69f9d28ed4c2e429f7a45f8e6fb7ac0e29181a832edd806e4a9d7afbff9a4638810e07043e5417ffa30383c688ed2a04fd0ea5469f4866db466eb5a6572022f569b94a4d26f54aa0d74e45b1a835aad8a53d13a734bafa3e970ad56ff262f82fc2f0000ffff7848a88329070000"
	pubInput := "7b22617574686f72697a655370656e6448617368223a223134343638353132333635343338363133303436303238323831353838363631333531343335343736313638363130393334313635353437393030343733363039313937373833353437363633222c22636f6d6d697456616c756558223a223134303837393735383637323735393131303737333731323331333435323237383234363131393531343336383232313332373632343633373837313330353538393537383338333230333438222c22636f6d6d697456616c756559223a223135313133353139393630333834323034363234383739363432303639353230343831333336323234333131393738303335323839323336363933363538363033363735333835323939383739222c226e756c6c696669657248617368223a2236373437353138373831363439303638333130373935363737343035383538333533303037343432333236353239363235343530383630363638393434313536313632303532333335313935222c2274726565526f6f7448617368223a223130353331333231363134393930373937303334393231323832353835363631383639363134353536343837303536393531343835323635333230343634393236363330343939333431333130227d"
	proof := "1f8b08000000000000ff62666360d6f9dfc8ccc8c81a50949f9fc6f8bf8981919991c9b188f17f0b0323b377513198c1e454ccf8bf83818141e57f3333232387bba1635a5a665e2a4492913182f17f1b0323632488626010ffdfcac8c8c8ee9a939a9b9a5702966263e400696e076b368269ee8069ee8268ee62606050f8dfc9ccc8c8946a04166562647234009bc0e4680831fd1fa3edff264646961f4b17bc559e92c1d8f343fd7330fbad90f4901f97c5f29e7b6f62b9fd43e587655550bcb31423cb0fab1c79a6ac03d96b7f1c6fd818f42cbae4d70f7b9e7b47fdc2af6ff921c674d2b77f9d161303c8b4596ad70f652df5baff43cfeec0ddaf7587a7ffe83a7269cdc49ffa413fe4163ac71d115a5fc9c8f223df3efbc7e28a05cf7eb85cdd357fc55b2ec11f36f6a7ce1de6ddc7f7436af2e2c499a2c72f3030828c7b147f655e75b493f90f693ecf2c815d1eae3faee795b77bfe5aa1f8438c63d5b21fe1930f30b2fc5876eaba89a3b482c40f5eb1dc85ba179a0eff70f233b30bf065dff143716ed5bc4b67f63a831d9771b9ed95834dfb991ffcd75efd67620d9dfea37945d5e6d825c6c77e30b93cdb6aeb60b98891e5c7a51f32a55bfddecbfef87aba7845448616d78fbbebff9a6fe417f6f9c1a31abf8dbf22c09c818101100000ffff5e148892f4010000"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)
}

func TestTransferOutput(t *testing.T) {
	verifyKey := "1f8b08000000000000ff6c517b5493e7197f9f2f9f92145456b90c2beb97910de6412529322855ca25230c4a018b84d2400324342b0997844b8a45b9050541a508c74209acdc3ac54da954b13827eddada1db04015b40a94bb56a065f43ce0816f2741ba9d9dfef53bef73f95dde87b39170fcd93c0e80d50145ba4aa957699282157a60f309d00062600b09508122608f995008ec0902566119f1c9aa84204d6a864e0b6c152184cf167000380aa1686d03287f37600d04287fa1090971618b380094c2c35ce600e5e706ec1102949ff0098a4c4808c3169b07456b65a07cdd802d2140f90a4d48883d7b14002cc4c90ab542a333f736029710055bca01f0d6ead23312744c0e132857abe5a18a24265e23dae3b12b50e4ab54aa340a6f264091acfb9906f3f65a4ae0ae6f025b4e80bb3e6e7a112260cb3800dcf525f308052035bb85e8b50c22f63807c0f1272bc14c8cec8998f0bf622708004030b095849067d9772800ebff9f333581ad30e99e34ebfe54af58d72d59d3357d8d1d7b8a02e0c6c8b4ba749526c9741bb02264953ec4e69bb4683c20b63cdaf051451c8a777afe7ad4a2a6111d54c3bb6ed42d7f8d4ef2a521e77f958881c6a9ece473dd9d952adc6b8f737a1f4a8faacf07360a0e8f033ef3418e93ddc1e74e1113db8349cbae3b2e0f7ec4b7bfec2b6e316c1760c38edad1cba18a41fccdcbf97263e37424d0f85ee48601f2eddc2296c91999952d358ea91bf407e178ec37681f2f2cfd6ed1ed5766b65b499b0a4324335c2c492daef9a1f0b2173e3d517fca273d7c2f6e17367dda9a92550a3476ddbcb7f0e1e6540de6bd65bc1d27731a41267d8bcd0b11b4017f3bd36bad73ee1d20c49cd5921779fdda969a79bcd2f88d75e2fe33b33816d259b47de6ae230a44fa10bf30810268ec197499f451dfb5c55ef99feba6e5679ec5fca9132f3546cfce22f769073254f1418cd91d23688ff593949d45ff9271c5e16dafcee317bfffebf4a60ca706fc657be0b0fae3776781c6b18bb1a5f39b7c6af164024f387b52fa3166f6d68b659dda41b4f65af431d4f29566b660b46116c61f36e1e6955583ee48f80bd8fcdac3a2e16b775974ec8ed5afb64b8a81c68680c8b1f3b99b34b878e892dfada8e25194462496d5712d8750909dad6ca7f33b88392cd0b8b1b6ded5767079043583cad696eef2257cbfa23ec5e62d7bfba5462ef8b57bb9038df74801f7c51d3d17f06a5bd285fd1ff68c60577f887afe95d3a1e8ba3b6065c5f919b5d95e4b5779fb4e65b60173f5b21b3c55d6efd05a96a6d46c7db31a69abab13f5d16d9f028d67dd2ddab837176ab0dfeef1b9c2859666fceab6e0d260e6bc11ada63cba3ff953c4eb4f4e618c146cb89653178c9f7b495e7cd57bee7b0c8bf8f68646981f85fc4faabb3cba624f038dff8ef3f5f0ffbaf93c5e77b837909bea7a1fc7e8adbe47bd8403b8eb8c45b3e3c054a7d95d39d728716a6cd1636f9abaf2d0757e3c6edd6151fc6533b5015da3bcb54131b71f018ddd071fbba425e9a598bbc571e40fcf5996e236b7bf5d652ef082d0e598b3435847ffaaf9f368a0f1d1f81589edeef01cd42a1b9bcee6317c14b4d91618578bcb50309d9c5ca6b47e1e6814c738d655dbfcd0819925dbec6a2fed1fc0e9bf74f4a2cfbc273233799fd57ebf7d99008d2f9d4b8cdd13f95a0406839b17ed1ba7c1f4d3fb721e891daad1fe299b8bf1055fec031adf9fe6fdfd61db8379b4b78b5fa0de7db31fbfbb9331e45639c7a0c07a595cc5b89e3791653fafbe39b9794f183259fd7613e17ffc27b61a78f117fb8c89b8bbe89de1a2e3bc23402f5db99cf191f19627fe9835298d7e6369024baa5ab599af53ff409eefa46e748ae59bb8bab582c7b63d2136f85e0ff9aa4ada178b76090e2df7d3c2a4c8177e6678854f39018d21ed45fbf4b2e870dcf9d49d6345eec615ec3bdc34b437e0fe0cba24a946dd259ebf200468ab8414b55aa53b204fce5048b99a149d4222d7bef1bfd568eecba1e2b8a8a00831f94f000000ffffeb52afc268060000"
	pubInput := "7b22636f6d6d697456616c756558223a223134303837393735383637323735393131303737333731323331333435323237383234363131393531343336383232313332373632343633373837313330353538393537383338333230333438222c22636f6d6d697456616c756559223a223135313133353139393630333834323034363234383739363432303639353230343831333336323234333131393738303335323839323336363933363538363033363735333835323939383739222c226e6f746548617368223a223136333038373933333937303234363632383332303634353233383932343138393038313435393030383636353731353234313234303933353337313939303335383038353530323535363439227d"
	proof := "1f8b08000000000000ff62666360d6f9dfc8ccc8c81a50949f9fc6f8bf8981919991c9b188f17f0b0323b377513198c1e454ccf8bf83818141e57f3333232387bba1635a5a665e2a4492913182f17f1b0323632488626010ffdfcac8c8c8ee9a939a9b9a5702966263e400696e076b368269ee8069ee8268ee62606050f8dfc9ccc8c8946a04166562647234009bc0e4680831fd1fa3edff264646961f4ddc5f18955a16ccfba1906ff768ab2187fe8fd6dc5bdec1568cfb7ef0774915ab3759b833b2fc70d0df70ecdbc229d37ef031ea35c5ead87ef991a2de64615aa974e987988bd4810762d39d1940a6d56dd9197b33c1bce9c7e6cbf358e35e886ff8c179be6ffeeb599a677f30bdbebdecd18c8c7246961ffed93a5b3cd719c8fff8cc21bd6af7928ed21f46f53fdf7b37dd53f8c1ffe37dd0267ecd4a06469071611bb9839dde3c7bf383bb69def593f6f9b63ff6301fdc3439ce37f007afd8d50dd704d22519597ecc9af92af26ff9f99b3ff2bed878af787c30ed878dd10957162e158f1ffaaaafb86c8bcfed003b6ec5e9a3898bbfb708fdd0ba35713effa2e7753f7eff6f9ddfa6d858ff43e8a4d2d18b7edd5e8c2c3fa4774c15ce9c513fe5477df7f9836d519ef63fb436acf9b6563d2fed87e08463fb97acdbf69e818101100000ffff9650c672f4010000"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)
}

func TestWithdraw(t *testing.T) {
	verifyKey := "1f8b08000000000000ff6c537d5453e7197f9fe4960617156d0d29394ac05834a3986482b66c70f81a512a16da0382c50af1068321d0e4c615e99ce02818141094da16655a70c2290504c78208a3881b4c3a8182f4503eac43261f0317910715ee4e82b87376fad7efdce7e3f7f13ee772ed0837804de302f02368bd469da2d1c587d029c0a613a0008280fd3d014eb002d8135694039b4780ff8e314eab516dd7251b1903b085841017f61817804bcb158b1bc0099001fb09014e80dc8a846c6433b8001cdacb56e602c75f066c16018ebffc392aac488898cdb40d2a16cbc0f193016b22c0f1935b911047f63800bc1ca4a513691d63ebd9018f109acde602781b18bd51c58853c5c1b18989b1a174bc384ea7f0f4f20856f8a9d51a1ded2d0ea4b5cc4f34c4bf5d4c09bca54d607308f096c6ad5f8448d8935c00ded2926d8403b0dbe616a2163328d85c2ec0da175642c47b629e8bc9ff279647000042803d4d0871660b38000eff3f676d029b6fd53d65d37d51cf5fd2352dea5a9f46c09ee100f0f6c41818bd46176fbd0df00959a032d874ab16857da70f8f948e9e9062cad480bae6d2c91f314ddffbf61f3dbfbc882e376bba1b4cb5e140a1c279de9b2a7ae4809f383e35efeceddd8e5762fa7fd0d5d7e6a2c4e7508960f3d79f112bdbb27fd8d75feed17461a25b5b9ca083fa17666c3175be7be7f8f7e811754b109810390214bafc6aae30e395b25abc70eb03d9d17ede216c8e36a6fab945cfe14aa78faaa51bcd421b5b961d73362c31f25becee34956eca6b6fc7494b51ccf227ea87c86b34ba3a4538ee030ad71e0c3de5dec9cee07f627d974d2cf8f4a15f727c7f5282d8199da45d9464f5966384d8b2f2d95b4655ee9f5ec368a367b3ea2fd78598be22c43e3bae76083d4a0a15c9a2723350a8ec93a9aa72393378bb363ff59d2f72b23038fc656816dd8d40c1fdd2d991d47e6273d750b5ee8e2684e9c1936c7df4ba29d31bd8407f25ebacbd3f8bbcf1753efa82f759a030c8e0f94dd6a3732a6425f71fdc345beea1ef65cb4853dde35fa363e58ee552876f37d9d89c3e757bba55592cc5536be263c69ac27230aa5a317f042235e89c90d6ba71bd4f2750181c582fca3f3a50819bf2473f74383f5d867f086ad6cc4739d12879602fee72f06f26b6b040e1eaf4d95725d983efa3b7cebba0c1f2d1666cdc57d836fc79823dbaa7ac5afbc6d40d2150f8b9c9edd8d7ca7fa6e160e85fb7bf3ef4960555473f4d1d985eb3033dc646cd7f0e6996d8fc6d8aa00376c9ab8ad1b7473deebb42ae408972f377153f0b14a2ebf94b72fc37d90a14ee936953cdfe37bb716bfa7729bd97274670d5179e76fdbdca1fd1a573fa6ee3ed7dc1cf6ff1f09786d13b39f3edc8cb3e525692e1e58a5507ee659c0b379f46d1ceea71ed36f30da070e5973b43ab3e7eb30cf9173b8a6b5ecb34a0b9f1f543426676145dc7736987960ca5cd5ddab365439cc83477f4c81b8f8bcef4ff0dbaf9eda2b99b0f16e3aaaa3d1b8a64dbde030ac3455b8a249c4b73e8badc4e7455e0d8896196736f6b2b43a3903b30e6f4d54aee61dbeb5140e16ec589caddd3fa02bc56f5f784eb2bae8ce1f09903257b5d1f04a1d490a4ce9b7b7416286c6f167aed7fc53085d72a9203ce3a3947a04f56e68557071f77e0aaba4cd55edfa3ad0428f4722ec809134ce7e24bbf138dcaebf8223cf5445d9fddd1e28bee7dfdc3378ec8fd81c2ee69e6b3c727fbb7e05bc75f6a293edfba17ef2d7c2f56b70ad7a3507cd1e3e39e1db156b262972b3cff836d1770f2cc42d358dd8636740f7be63279b5c5841b8603cb4f145dcb060ad7284f4fecb08bfa10eba55793f9dffca2021766ae87cd6ccdaec1d5da1f982641f5b895ece0d0c4c3ae49bf687ca68d286f0ddb5d8aa5bd92a772fbfd4328e23c29b734e4fecdfa23aee0255696fe3c182d313b993d58b91e4705eae295b787df44e97bc9ae7707b60d1202d4ea5823732049af394cbf9b4cebf62b630d0778bb42833e88dc1e1ec467f4341d9e94c458ab76b18949461d43fe1b0000ffff42520e4f6e060000"
	pubInput := "7b22616d6f756e74223a223238323432303438222c22617574686f72697a655370656e6448617368223a223134343638353132333635343338363133303436303238323831353838363631333531343335343736313638363130393334313635353437393030343733363039313937373833353437363633222c226e756c6c696669657248617368223a2236373437353138373831363439303638333130373935363737343035383538333533303037343432333236353239363235343530383630363638393434313536313632303532333335313935222c2274726565526f6f7448617368223a223130353331333231363134393930373937303334393231323832353835363631383639363134353536343837303536393531343835323635333230343634393236363330343939333431333130227d"
	proof := "1f8b08000000000000ff62666360d6f9dfc8ccc8c81a50949f9fc6f8bf8981919991c9b188f17f0b0323b377513198c1e454ccf8bf83818141e57f3333232387bba1635a5a665e2a4492913182f17f1b0323632488626010ffdfcac8c8c8ee9a939a9b9a5702966263e400696e076b368269ee8069ee8268ee62606050f8dfc9ccc8c8946a04166562647234009bc0e4680831fd1fa3edff264646961fa722d83eb67ed939f5077b2a1bcf02bb2b1f7ea46b94bf7871efbbe80f96dc39b35e39e52631b2fcf8c19dee7fa7d6a6ea87f9e955bb362e7991f8e396d9a4159b6f7edcff43eaa3a4c0359798db0c20d3929e3efeee6c3cfdc08f9a933b17ddd9dd39ef47d52a69f9b079417b7e30cee4dcc35971ad8091e58708e71d06bb4606bd1f7b970b455d92669aff2372d2f4b86db2e73eff50ba24ff5089bff22d0323c8b83ced3a63a19cacb73fceabd4fe0d53976ff97144f8959e5f7c8cdf0f7e25ced5d65b1eda31b2fc60d79d2cedf0fd8ee70fff6d332ffc914d09fcb16169aa504dc4d4841fdc934f842cb54c7809765c4777e772d9a4979d3f0ecd7ba7bd3442d7f5c7f713e2bc336f6feafec1abb0f3d6b1bdede58c2c3fc2b24ef62e597a25f1c7926f3ed725fd5638ff38a3bdb9f435c779c51f5ceab3b64db1322861606000040000ffff877599d4f4010000"

	rst, err := Verify(verifyKey, proof, pubInput)
	assert.Nil(t, err)
	assert.Equal(t, true, rst)
}

//
//func TestDeserializeWitness(t *testing.T) {
//	//str := "7b0a202020202259223a202230783233222c0a202020202278223a202230783033220a7d0a"
//	str := "7b22636d7468617368223a2230393265613834303464323431303566326263616638616132656365633239626230303039353965376137313433393539663465623638373232656531653765222c22636d7476616c756578223a2231663235383330333630353866613432376338346364663264303361666234363233623361376566313365366637653465656239663762316335666563326463222c22636d7476616c756579223a2232313639663266626634623839653431343631643734326430373036663362313137643761613531376663313164333261336563313765666664346638336137222c22636f6d6d69745f72616e646f6d223a226138353566323233222c2268656c70657231223a223031222c2268656c70657232223a223031222c2268656c70657233223a223031222c2268656c70657234223a223031222c2268656c70657235223a223031222c227061746830223a2231653434633733626137393830623034353061386539393763396439633738626535613964376365616635393764663738313436396131633964623465346339222c227061746831223a2231393161383065333737616639653064303465316437356538643730326434633264623231623935326239666666366263636133316639653966643564653030222c227061746832223a2232323064633230343161386338313038366139626338303834653766396565303738386165326338623739323865306232653236373233333962323933336233222c227061746833223a2232393564363331666366306564306433343734326163353630666537613063303538353232356366663931393438303664346439643865316630306531373437222c227061746834223a2232313538343961643762643433343461383037663466306339616566643231333135373265336362323161386232316161393665386131316334613231346535222c227061746835223a22222c227072696b6579223a2232613831613633623731616534383132343836396661316537363664643637633164336236383431653561646132643362623831326661303432363730313733222c227075626b6579223a2231356639646364656362396432623963323635643936376534623135363532636564393862663530303566646236623562333030626630376335303666386239222c22726f6f7468617368223a2232616665613563323866373631663432663335636361343731313730636130373263626534623639623662313863346133623435363337653937343738336334222c227472616e736665725f72616e646f6d223a223233222c227472616e736665725f76616c7565223a223031616566303830222c2276616c696430223a223031222c2276616c696431223a223031222c2276616c696432223a223031222c2276616c696433223a223031222c2276616c696434223a223031222c2276616c696435223a22222c2276616c7565223a2233306537227d0a"
//	res, _ := hex.DecodeString(str)
//	var buf bytes.Buffer
//	buf.Write(res)
//	decoder := json.NewDecoder(&buf)
//
//	var ss mixTy.TestMsg2
//
//	if err := decoder.Decode(&ss); err != nil {
//		panic(err)
//	}
//	val, _ := hex.DecodeString(ss.Cmthash)
//	var a big.Int
//	a.SetBytes(val)
//	fmt.Println("cmthash=", ss.Cmthash, "val=", a.String(), "v=", a)
//	fmt.Println("cmtvaluex=", ss.Cmtvaluex)
//	fmt.Println("y=", ss.Cmtvaluey)
//	fmt.Println("roothash", ss.Roothash)
//	//v,err := strconv.ParseInt(ss.Value,16,0)
//	b, err := new(big.Int).SetString(ss.Value, 16)
//	fmt.Println("erri", err)
//	fmt.Println("value", ss.Value, "atoi", b)
//
//}
//
//func TestDeserial(t *testing.T) {
//	val := make(map[string]interface{})
//	val["cmthash"] = "1"
//	hash := "0x2afea5c28f761f42f35cca471170ca072cbe4b69b6b18c4a3b45637e974783c4"
//	hashb, _ := hex.DecodeString(hash[2:])
//	val["value"] = hashb
//
//	res, err := json.Marshal(val)
//	assert.Nil(t, err)
//	str := hex.EncodeToString(res)
//	fmt.Println("res", string(res), "str", str)
//
//	str2, _ := hex.DecodeString(str)
//	var msg mixTy.TestMsg3
//	json.Unmarshal(str2, &msg)
//	fmt.Println("msg", msg.Cmthash, "val", hex.EncodeToString(msg.Value))
//
//	val2 := make(map[string]interface{})
//	var buff2 bytes.Buffer
//	buff2.Write(str2)
//	decode := json.NewDecoder(&buff2)
//	err = decode.Decode(&val2)
//	assert.Nil(t, err)
//	for k, v := range val2 {
//		fmt.Println("k=", k, "v=", v)
//	}
//
//}

func TestDeserial(t *testing.T) {
	val := make(map[string]interface{})
	val["cmthash"] = "1"
	hash := "0x2afea5c28f761f42f35cca471170ca072cbe4b69b6b18c4a3b45637e974783c4"
	hashb, _ := hex.DecodeString(hash[2:])
	val["value"] = hashb

	var buf bytes.Buffer
	encode := json.NewEncoder(&buf)
	err := encode.Encode(val)
	assert.Nil(t, err)
	enstr := hex.EncodeToString(buf.Bytes())
	fmt.Println("encode", enstr)

	enbuff, _ := hex.DecodeString(enstr)
	var buff2 bytes.Buffer
	buff2.Write(enbuff)
	decode := json.NewDecoder(&buff2)

	val2 := make(map[string]interface{})
	err = decode.Decode(&val2)
	assert.Nil(t, err)
	for k, v := range val2 {
		fmt.Println("k=", k, "v=", v)
	}

}

func TestDeserialDeposit(t *testing.T) {
	pubInput := "7b22616d6f756e74223a223238323432303438222c226e6f746548617368223a223136333038373933333937303234363632383332303634353233383932343138393038313435393030383636353731353234313234303933353337313939303335383038353530323535363439227d"
	var deposit types.DepositPublicInput
	data, err := hex.DecodeString(pubInput)
	assert.Nil(t, err)
	err = json.Unmarshal(data, &deposit)
	assert.Nil(t, err)
	fmt.Println("notehash", deposit.NoteHash)
	fmt.Println("amout", deposit.Amount)
}