// SPDX-FileCopyrightText: NOI Techpark <digital@noi.bz.it>
// SPDX-License-Identifier: AGPL-3.0-or-later
package sharing

// Polygons for service constraint zones.
// As per Netex spec they are in WGS84/EPSG4326
// You can use QGIS to load features from some source (e.g. shape file), then select the area and export the layer as GML, always with CRS WGS84
// Then open the file in an editor and copy the content of <gml:Polygon> tag to here

const GML_MUNICIPALITY_BZ = `<gml:exterior><gml:LinearRing><gml:posList>46.5275039702865 11.3587014047429 46.5272458003351 11.3642287367731 46.5126359020919 11.3560933767181 46.5082982592365 11.3667525334481 46.5090479750115 11.3742167406329 46.5071409207445 11.3841182097064 46.5012439911105 11.3889546612642 46.5014436496046 11.4062545307947 46.500335098342 11.4106001829749 46.4965317542823 11.413839915913 46.4924525153392 11.4226679266433 46.4940974099184 11.432725642416 46.4902452309409 11.4287138081684 46.4888846615736 11.4243361132799 46.4940122900117 11.4108296625322 46.4951671323593 11.3944223735971 46.4931732397073 11.3916471028257 46.474917317277 11.4055623251629 46.4731869458676 11.4043918060758 46.4667220518721 11.4079501612313 46.4598779107241 11.4164626382075 46.4531536703352 11.4094864931315 46.4510339709512 11.3861133855019 46.4493990034369 11.3810044987022 46.4730152638628 11.3588517366422 46.4655349062171 11.344920871054 46.465948929995 11.3383293829128 46.4590027449559 11.3258894044549 46.4467874268319 11.314476120754 46.4469945201697 11.3084630291187 46.4552412040007 11.3018578005043 46.4708260810831 11.2999978361319 46.4781238947376 11.3025136712178 46.4936863997566 11.2847142076253 46.4978323299235 11.2753897690417 46.50925525935 11.2917328440625 46.505623679384 11.2993583393625 46.5064787848292 11.3102895987477 46.5091018083264 11.318092311088 46.5149644943131 11.3218990625113 46.5207538124991 11.3319519359821 46.5228043621688 11.3377505366013 46.5216830123961 11.3454821063847 46.5259325337733 11.3515392400704 46.5303831876479 11.3537159748611 46.5275039702865 11.3587014047429</gml:posList></gml:LinearRing></gml:exterior>`
const GML_MUNICIPALITY_ME = `<gml:exterior><gml:LinearRing><gml:posList>46.6896835034245 11.1364290879322 46.6899058416621 11.1442653461878 46.6786968224318 11.1629050078299 46.6761345251936 11.1618155452526 46.67188085073 11.16664262685 46.6756631685687 11.1731532593277 46.6868735035905 11.178288955614 46.6868458850561 11.1809323167838 46.671557044163 11.1860234677975 46.6731021551046 11.198846485147 46.6708489494787 11.2081258049239 46.6760664554599 11.2397148499297 46.6716576302575 11.2396633624789 46.6673656886121 11.236858300506 46.661058704873 11.2124777025187 46.6526094283735 11.2065816528681 46.6422129476589 11.2059774134719 46.6419256793078 11.2087375611271 46.6367399595257 11.2127924006032 46.6280729781818 11.2020828237554 46.6191536508851 11.2056166024823 46.616972965164 11.2048259898535 46.6222757080356 11.2016908953922 46.6201056229653 11.1842403177782 46.6222796459246 11.1810560158268 46.637425986679 11.1723169967497 46.6575447167387 11.1458623576531 46.6657406403101 11.1416881732664 46.6675916026944 11.1389612203664 46.6707367120655 11.1457015453499 46.6788606070854 11.1453694733908 46.6858534085117 11.1369515053403 46.6896835034245 11.1364290879322</gml:posList></gml:LinearRing></gml:exterior>`
const GML_PROVINCE_BZ = `<gml:exterior><gml:LinearRing><gml:posList>47.0865306118466 12.2051127346201 47.0686112919437 12.2402943390512 47.0278888156892 12.2047904598862 47.0242436719994 12.1469326074475 47.0066535775866 12.1209882633712 46.9629470382391 12.1315082166952 46.9378891358908 12.1682053335818 46.9140110066375 12.1441333208624 46.9062339488672 12.190163939063 46.8738693778566 12.2162303227969 46.8910821637859 12.2421865100843 46.8843972821332 12.2740196718489 46.8409438396897 12.30692125541 46.8149916513324 12.2824306229562 46.7829021228497 12.2838431254067 46.7746633485312 12.3574065618128 46.7165576970333 12.3836912463737 46.6798212638204 12.4775943539877 46.6429073956601 12.3797735603076 46.6227817086307 12.3851633132073 46.6313778944669 12.3397048263089 46.6179260587001 12.2838365149815 46.6291248839294 12.2612199361179 46.5945638629152 12.1922397858842 46.6193326136138 12.1942640581225 46.6750712857559 12.0682696593627 46.6426223973064 12.074224775677 46.5847244182347 12.0466934998341 46.5328557179936 11.9984188961765 46.5446779215668 11.9661474707714 46.5294169498902 11.9454946931061 46.5329602544659 11.9118839483467 46.5089139467489 11.828312400847 46.5327642901224 11.8111847168627 46.5091947310819 11.7746402494044 46.4992884275568 11.6328077298869 46.4830664892938 11.650260069412 46.471092949453 11.6247616882655 46.3877733429757 11.5998184577991 46.3812833570515 11.5642078958628 46.3455768758891 11.5455664864725 46.3639314013008 11.4768968145951 46.3347265788764 11.4544716439232 46.3008446877193 11.3607463712603 46.2742338198765 11.3943610001384 46.261469874046 11.3884198700378 46.2656676251445 11.3587808196001 46.2987106471577 11.3380700284323 46.2587670891491 11.3056164893551 46.2283818007903 11.2394722033083 46.2197724030662 11.2064291843848 46.2326672968392 11.1746381717514 46.2556046331325 11.2023783920202 46.2833300590352 11.1388106354162 46.3420979645282 11.2036122314289 46.3597941641161 11.1915649457967 46.4620482061655 11.2204213327919 46.5087285878493 11.1870849756078 46.4816293604976 11.1296786204622 46.5018478856677 11.0894473968926 46.5311756333736 11.0806366472673 46.5137787391814 11.0346009570852 46.4423469283568 11.0748536101506 46.4840251355235 10.9757482199462 46.443998427718 10.9130442931385 46.4511555127019 10.883632574071 46.4361313339799 10.8613240558654 46.4429661483581 10.8003805352273 46.4859571970675 10.7649772934059 46.4514708660184 10.6846435956634 46.4479606434984 10.6218369823205 46.4914566598879 10.5519607918894 46.4936008152278 10.4845306965504 46.510589612536 10.4579404260749 46.615116392487 10.4922940591132 46.6411036664124 10.4459538129023 46.6370478944115 10.4016077150143 46.6801821268263 10.3861051908103 46.7142960980867 10.4175405854345 46.7340770892702 10.4010679090777 46.7528572685974 10.4421050235955 46.7887893022408 10.4246732711661 46.8587117327541 10.4783542822263 46.8497818662742 10.5526302921984 46.8391707319163 10.5507180050461 46.875885991914 10.6679335520621 46.8232720660066 10.7633624817478 46.7879070966077 10.730494129755 46.7947005793793 10.7887843938403 46.7754120785198 10.814115077222 46.7816919155649 10.8407049509755 46.7631095848559 10.8831955543511 46.7749173529939 10.9231956999466 46.7651748112482 11.0221353703125 46.8050842352593 11.039798522775 46.8222199255807 11.0834527306058 46.8528607474226 11.0714239586224 46.8898597101632 11.1015461061389 46.9122335303534 11.0959013548118 46.9310053783462 11.1150127325481 46.940713074559 11.1651948298158 46.9655681951271 11.1645252112337 46.962599678831 11.206869301652 46.9924271968353 11.3193058937635 46.9903616923194 11.3582865017718 46.9652447641731 11.4009085191906 46.9764999446439 11.4420906806552 47.0109094091641 11.4793158699107 46.9841080865442 11.5380769385088 47.012577350939 11.6272104537395 46.9926276105131 11.6641040539233 46.9930235592184 11.7112319154652 46.9689027227926 11.7471681648721 46.9920590539057 11.7817826665826 46.992896204962 11.8362104979517 47.0497542087876 11.978281084477 47.046760178266 12.0200610285333 47.0775349490327 12.0971308854383 47.0865306118466 12.2051127346201</gml:posList></gml:LinearRing></gml:exterior>`