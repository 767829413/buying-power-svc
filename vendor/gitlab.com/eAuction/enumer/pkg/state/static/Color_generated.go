package static

import "gitlab.com/eAuction/enumer/pkg/state"

type Color int32

const (
	ColorOther                      Color = 1
	ColorAirForceBlue               Color = 2
	ColorAliceBlue                  Color = 3
	ColorAlizarinCrimson            Color = 4
	ColorAlmond                     Color = 5
	ColorAmaranth                   Color = 6
	ColorAmber                      Color = 7
	ColorAmericanRose               Color = 8
	ColorAmethyst                   Color = 9
	ColorAndroidGreen               Color = 10
	ColorAntiFlashWhite             Color = 11
	ColorAntiqueBrass               Color = 12
	ColorAntiqueFuchsia             Color = 13
	ColorAntiqueWhite               Color = 14
	ColorAo                         Color = 15
	ColorAppleGreen                 Color = 16
	ColorApricot                    Color = 17
	ColorAqua                       Color = 18
	ColorAquamarine                 Color = 19
	ColorArmyGreen                  Color = 20
	ColorArylideYellow              Color = 21
	ColorAshGrey                    Color = 22
	ColorAsparagus                  Color = 23
	ColorAtomicTangerine            Color = 24
	ColorAuburn                     Color = 25
	ColorAureolin                   Color = 26
	ColorAurometalsaurus            Color = 27
	ColorAwesome                    Color = 28
	ColorAzure                      Color = 29
	ColorAzureMistWeb               Color = 30
	ColorBabyBlue                   Color = 31
	ColorBabyBlueEyes               Color = 32
	ColorBabyPink                   Color = 33
	ColorBallBlue                   Color = 34
	ColorBananaMania                Color = 35
	ColorBananaYellow               Color = 36
	ColorBattleshipGrey             Color = 37
	ColorBazaar                     Color = 38
	ColorBeauBlue                   Color = 39
	ColorBeaver                     Color = 40
	ColorBeige                      Color = 41
	ColorBisque                     Color = 42
	ColorBistre                     Color = 43
	ColorBittersweet                Color = 44
	ColorBlack                      Color = 45
	ColorBlanchedAlmond             Color = 46
	ColorBleuDeFrance               Color = 47
	ColorBlizzardBlue               Color = 48
	ColorBlond                      Color = 49
	ColorBlue                       Color = 50
	ColorBlueBell                   Color = 51
	ColorBlueGray                   Color = 52
	ColorBlueGreen                  Color = 53
	ColorBluePurple                 Color = 54
	ColorBlueViolet                 Color = 55
	ColorBlush                      Color = 56
	ColorBole                       Color = 57
	ColorBondiBlue                  Color = 58
	ColorBone                       Color = 59
	ColorBostonUniversityRed        Color = 60
	ColorBottleGreen                Color = 61
	ColorBoysenberry                Color = 62
	ColorBrandeisBlue               Color = 63
	ColorBrass                      Color = 64
	ColorBrickRed                   Color = 65
	ColorBrightCerulean             Color = 66
	ColorBrightGreen                Color = 67
	ColorBrightLavender             Color = 68
	ColorBrightMaroon               Color = 69
	ColorBrightPink                 Color = 70
	ColorBrightTurquoise            Color = 71
	ColorBrightUbe                  Color = 72
	ColorBrilliantLavender          Color = 73
	ColorBrilliantRose              Color = 74
	ColorBrinkPink                  Color = 75
	ColorBritishRacingGreen         Color = 76
	ColorBronze                     Color = 77
	ColorBrown                      Color = 78
	ColorBubbleGum                  Color = 79
	ColorBubbles                    Color = 80
	ColorBuff                       Color = 81
	ColorBulgarianRose              Color = 82
	ColorBurgundy                   Color = 83
	ColorBurlywood                  Color = 84
	ColorBurntOrange                Color = 85
	ColorBurntSienna                Color = 86
	ColorBurntUmber                 Color = 87
	ColorByzantine                  Color = 88
	ColorByzantium                  Color = 89
	ColorCgBlue                     Color = 90
	ColorCgRed                      Color = 91
	ColorCadet                      Color = 92
	ColorCadetBlue                  Color = 93
	ColorCadetGrey                  Color = 94
	ColorCadmiumGreen               Color = 95
	ColorCadmiumOrange              Color = 96
	ColorCadmiumRed                 Color = 97
	ColorCadmiumYellow              Color = 98
	ColorCafeAuLait                 Color = 99
	ColorCafeNoir                   Color = 100
	ColorCalPolyPomonaGreen         Color = 101
	ColorCambridgeBlue              Color = 102
	ColorCamel                      Color = 103
	ColorCamouflageGreen            Color = 104
	ColorCanary                     Color = 105
	ColorCanaryYellow               Color = 106
	ColorCandyAppleRed              Color = 107
	ColorCandyPink                  Color = 108
	ColorCapri                      Color = 109
	ColorCaputMortuum               Color = 110
	ColorCardinal                   Color = 111
	ColorCaribbeanGreen             Color = 112
	ColorCarmine                    Color = 113
	ColorCarminePink                Color = 114
	ColorCarmineRed                 Color = 115
	ColorCarnationPink              Color = 116
	ColorCarnelian                  Color = 117
	ColorCarolinaBlue               Color = 118
	ColorCarrotOrange               Color = 119
	ColorCeladon                    Color = 120
	ColorCeleste                    Color = 121
	ColorCelestialBlue              Color = 122
	ColorCerise                     Color = 123
	ColorCerisePink                 Color = 124
	ColorCerulean                   Color = 125
	ColorCeruleanBlue               Color = 126
	ColorChamoisee                  Color = 127
	ColorChampagne                  Color = 128
	ColorCharcoal                   Color = 129
	ColorChartreuse                 Color = 130
	ColorCherry                     Color = 131
	ColorCherryBlossomPink          Color = 132
	ColorChestnut                   Color = 133
	ColorChocolate                  Color = 134
	ColorChromeYellow               Color = 135
	ColorCinereous                  Color = 136
	ColorCinnabar                   Color = 137
	ColorCinnamon                   Color = 138
	ColorCitrine                    Color = 139
	ColorClassicRose                Color = 140
	ColorCobalt                     Color = 141
	ColorCocoaBrown                 Color = 142
	ColorCoffee                     Color = 143
	ColorColumbiaBlue               Color = 144
	ColorCoolBlack                  Color = 145
	ColorCoolGrey                   Color = 146
	ColorCopper                     Color = 147
	ColorCopperRose                 Color = 148
	ColorCoquelicot                 Color = 149
	ColorCoral                      Color = 150
	ColorCoralPink                  Color = 151
	ColorCoralRed                   Color = 152
	ColorCordovan                   Color = 153
	ColorCorn                       Color = 154
	ColorCornellRed                 Color = 155
	ColorCornflower                 Color = 156
	ColorCornflowerBlue             Color = 157
	ColorCornsilk                   Color = 158
	ColorCosmicLatte                Color = 159
	ColorCottonCandy                Color = 160
	ColorCream                      Color = 161
	ColorCrimson                    Color = 162
	ColorCrimsonRed                 Color = 163
	ColorCrimsonGlory               Color = 164
	ColorCyan                       Color = 165
	ColorDaffodil                   Color = 166
	ColorDandelion                  Color = 167
	ColorDarkBlue                   Color = 168
	ColorDarkBrown                  Color = 169
	ColorDarkByzantium              Color = 170
	ColorDarkCandyAppleRed          Color = 171
	ColorDarkCerulean               Color = 172
	ColorDarkChestnut               Color = 173
	ColorDarkCoral                  Color = 174
	ColorDarkCyan                   Color = 175
	ColorDarkElectricBlue           Color = 176
	ColorDarkGoldenrod              Color = 177
	ColorDarkGray                   Color = 178
	ColorDarkGreen                  Color = 179
	ColorDarkJungleGreen            Color = 180
	ColorDarkKhaki                  Color = 181
	ColorDarkLava                   Color = 182
	ColorDarkLavender               Color = 183
	ColorDarkMagenta                Color = 184
	ColorDarkMidnightBlue           Color = 185
	ColorDarkOliveGreen             Color = 186
	ColorDarkOrange                 Color = 187
	ColorDarkOrchid                 Color = 188
	ColorDarkPastelBlue             Color = 189
	ColorDarkPastelGreen            Color = 190
	ColorDarkPastelPurple           Color = 191
	ColorDarkPastelRed              Color = 192
	ColorDarkPink                   Color = 193
	ColorDarkPowderBlue             Color = 194
	ColorDarkRaspberry              Color = 195
	ColorDarkRed                    Color = 196
	ColorDarkSalmon                 Color = 197
	ColorDarkScarlet                Color = 198
	ColorDarkSeaGreen               Color = 199
	ColorDarkSienna                 Color = 200
	ColorDarkSlateBlue              Color = 201
	ColorDarkSlateGray              Color = 202
	ColorDarkSpringGreen            Color = 203
	ColorDarkTan                    Color = 204
	ColorDarkTangerine              Color = 205
	ColorDarkTaupe                  Color = 206
	ColorDarkTerraCotta             Color = 207
	ColorDarkTurquoise              Color = 208
	ColorDarkViolet                 Color = 209
	ColorDartmouthGreen             Color = 210
	ColorDavyGrey                   Color = 211
	ColorDebianRed                  Color = 212
	ColorDeepCarmine                Color = 213
	ColorDeepCarminePink            Color = 214
	ColorDeepCarrotOrange           Color = 215
	ColorDeepCerise                 Color = 216
	ColorDeepChampagne              Color = 217
	ColorDeepChestnut               Color = 218
	ColorDeepCoffee                 Color = 219
	ColorDeepFuchsia                Color = 220
	ColorDeepJungleGreen            Color = 221
	ColorDeepLilac                  Color = 222
	ColorDeepMagenta                Color = 223
	ColorDeepPeach                  Color = 224
	ColorDeepPink                   Color = 225
	ColorDeepSaffron                Color = 226
	ColorDeepSkyBlue                Color = 227
	ColorDenim                      Color = 228
	ColorDesert                     Color = 229
	ColorDesertSand                 Color = 230
	ColorDimGray                    Color = 231
	ColorDodgerBlue                 Color = 232
	ColorDogwoodRose                Color = 233
	ColorDollarBill                 Color = 234
	ColorDrab                       Color = 235
	ColorDukeBlue                   Color = 236
	ColorEarthYellow                Color = 237
	ColorEcru                       Color = 238
	ColorEggplant                   Color = 239
	ColorEggshell                   Color = 240
	ColorEgyptianBlue               Color = 241
	ColorElectricBlue               Color = 242
	ColorElectricCrimson            Color = 243
	ColorElectricCyan               Color = 244
	ColorElectricGreen              Color = 245
	ColorElectricIndigo             Color = 246
	ColorElectricLavender           Color = 247
	ColorElectricLime               Color = 248
	ColorElectricPurple             Color = 249
	ColorElectricUltramarine        Color = 250
	ColorElectricViolet             Color = 251
	ColorElectricYellow             Color = 252
	ColorEmerald                    Color = 253
	ColorEtonBlue                   Color = 254
	ColorFallow                     Color = 255
	ColorFaluRed                    Color = 256
	ColorFamous                     Color = 257
	ColorFandango                   Color = 258
	ColorFashionFuchsia             Color = 259
	ColorFawn                       Color = 260
	ColorFeldgrau                   Color = 261
	ColorFern                       Color = 262
	ColorFernGreen                  Color = 263
	ColorFerrariRed                 Color = 264
	ColorFieldDrab                  Color = 265
	ColorFireEngineRed              Color = 266
	ColorFirebrick                  Color = 267
	ColorFlame                      Color = 268
	ColorFlamingoPink               Color = 269
	ColorFlavescent                 Color = 270
	ColorFlax                       Color = 271
	ColorFloralWhite                Color = 272
	ColorFluorescentOrange          Color = 273
	ColorFluorescentPink            Color = 274
	ColorFluorescentYellow          Color = 275
	ColorFolly                      Color = 276
	ColorForestGreen                Color = 277
	ColorFrenchBeige                Color = 278
	ColorFrenchBlue                 Color = 279
	ColorFrenchLilac                Color = 280
	ColorFrenchRose                 Color = 281
	ColorFuchsia                    Color = 282
	ColorFuchsiaPink                Color = 283
	ColorFulvous                    Color = 284
	ColorFuzzyWuzzy                 Color = 285
	ColorGainsboro                  Color = 286
	ColorGamboge                    Color = 287
	ColorGhostWhite                 Color = 288
	ColorGinger                     Color = 289
	ColorGlaucous                   Color = 290
	ColorGlitter                    Color = 291
	ColorGold                       Color = 292
	ColorGoldenBrown                Color = 293
	ColorGoldenPoppy                Color = 294
	ColorGoldenYellow               Color = 295
	ColorGoldenrod                  Color = 296
	ColorGrannySmithApple           Color = 297
	ColorGray                       Color = 298
	ColorGrayAsparagus              Color = 299
	ColorGreen                      Color = 300
	ColorGreenBlue                  Color = 301
	ColorGreenYellow                Color = 302
	ColorGrullo                     Color = 303
	ColorGuppieGreen                Color = 304
	ColorHalayaUbe                  Color = 305
	ColorHanBlue                    Color = 306
	ColorHanPurple                  Color = 307
	ColorHansaYellow                Color = 308
	ColorHarlequin                  Color = 309
	ColorHarvardCrimson             Color = 310
	ColorHarvestGold                Color = 311
	ColorHeartGold                  Color = 312
	ColorHeliotrope                 Color = 313
	ColorHollywoodCerise            Color = 314
	ColorHoneydew                   Color = 315
	ColorHookerGreen                Color = 316
	ColorHotMagenta                 Color = 317
	ColorHotPink                    Color = 318
	ColorHunterGreen                Color = 319
	ColorIcterine                   Color = 320
	ColorInchworm                   Color = 321
	ColorIndiaGreen                 Color = 322
	ColorIndianRed                  Color = 323
	ColorIndianYellow               Color = 324
	ColorIndigo                     Color = 325
	ColorInternationalKleinBlue     Color = 326
	ColorInternationalOrange        Color = 327
	ColorIris                       Color = 328
	ColorIsabelline                 Color = 329
	ColorIslamicGreen               Color = 330
	ColorIvory                      Color = 331
	ColorJade                       Color = 332
	ColorJasmine                    Color = 333
	ColorJasper                     Color = 334
	ColorJazzberryJam               Color = 335
	ColorJonquil                    Color = 336
	ColorJuneBud                    Color = 337
	ColorJungleGreen                Color = 338
	ColorKuCrimson                  Color = 339
	ColorKellyGreen                 Color = 340
	ColorKhaki                      Color = 341
	ColorLaSalleGreen               Color = 342
	ColorLanguidLavender            Color = 343
	ColorLapisLazuli                Color = 344
	ColorLaserLemon                 Color = 345
	ColorLaurelGreen                Color = 346
	ColorLava                       Color = 347
	ColorLavender                   Color = 348
	ColorLavenderBlue               Color = 349
	ColorLavenderBlush              Color = 350
	ColorLavenderGray               Color = 351
	ColorLavenderIndigo             Color = 352
	ColorLavenderMagenta            Color = 353
	ColorLavenderMist               Color = 354
	ColorLavenderPink               Color = 355
	ColorLavenderPurple             Color = 356
	ColorLavenderRose               Color = 357
	ColorLawnGreen                  Color = 358
	ColorLemon                      Color = 359
	ColorLemonYellow                Color = 360
	ColorLemonChiffon               Color = 361
	ColorLemonLime                  Color = 362
	ColorLightCrimson               Color = 363
	ColorLightThulianPink           Color = 364
	ColorLightApricot               Color = 365
	ColorLightBlue                  Color = 366
	ColorLightBrown                 Color = 367
	ColorLightCarminePink           Color = 368
	ColorLightCoral                 Color = 369
	ColorLightCornflowerBlue        Color = 370
	ColorLightCyan                  Color = 371
	ColorLightFuchsiaPink           Color = 372
	ColorLightGoldenrodYellow       Color = 373
	ColorLightGray                  Color = 374
	ColorLightGreen                 Color = 375
	ColorLightKhaki                 Color = 376
	ColorLightPastelPurple          Color = 377
	ColorLightPink                  Color = 378
	ColorLightSalmon                Color = 379
	ColorLightSalmonPink            Color = 380
	ColorLightSeaGreen              Color = 381
	ColorLightSkyBlue               Color = 382
	ColorLightSlateGray             Color = 383
	ColorLightTaupe                 Color = 384
	ColorLightYellow                Color = 385
	ColorLilac                      Color = 386
	ColorLime                       Color = 387
	ColorLimeGreen                  Color = 388
	ColorLincolnGreen               Color = 389
	ColorLinen                      Color = 390
	ColorLion                       Color = 391
	ColorLiver                      Color = 392
	ColorLust                       Color = 393
	ColorMsuGreen                   Color = 394
	ColorMacaroniAndCheese          Color = 395
	ColorMagenta                    Color = 396
	ColorMagicMint                  Color = 397
	ColorMagnolia                   Color = 398
	ColorMahogany                   Color = 399
	ColorMaize                      Color = 400
	ColorMajorelleBlue              Color = 401
	ColorMalachite                  Color = 402
	ColorManatee                    Color = 403
	ColorMangoTango                 Color = 404
	ColorMantis                     Color = 405
	ColorMaroon                     Color = 406
	ColorMauve                      Color = 407
	ColorMauveTaupe                 Color = 408
	ColorMauvelous                  Color = 409
	ColorMayaBlue                   Color = 410
	ColorMeatBrown                  Color = 411
	ColorMediumPersianBlue          Color = 412
	ColorMediumAquamarine           Color = 413
	ColorMediumBlue                 Color = 414
	ColorMediumCandyAppleRed        Color = 415
	ColorMediumCarmine              Color = 416
	ColorMediumChampagne            Color = 417
	ColorMediumElectricBlue         Color = 418
	ColorMediumJungleGreen          Color = 419
	ColorMediumLavenderMagenta      Color = 420
	ColorMediumOrchid               Color = 421
	ColorMediumPurple               Color = 422
	ColorMediumRedViolet            Color = 423
	ColorMediumSeaGreen             Color = 424
	ColorMediumSlateBlue            Color = 425
	ColorMediumSpringBud            Color = 426
	ColorMediumSpringGreen          Color = 427
	ColorMediumTaupe                Color = 428
	ColorMediumTealBlue             Color = 429
	ColorMediumTurquoise            Color = 430
	ColorMediumVioletRed            Color = 431
	ColorMelon                      Color = 432
	ColorMidnightBlue               Color = 433
	ColorMidnightGreen              Color = 434
	ColorMikadoYellow               Color = 435
	ColorMint                       Color = 436
	ColorMintCream                  Color = 437
	ColorMintGreen                  Color = 438
	ColorMistyRose                  Color = 439
	ColorMoccasin                   Color = 440
	ColorModeBeige                  Color = 441
	ColorMoonstoneBlue              Color = 442
	ColorMordantRed19               Color = 443
	ColorMossGreen                  Color = 444
	ColorMountainMeadow             Color = 445
	ColorMountbattenPink            Color = 446
	ColorMulberry                   Color = 447
	ColorMunsell                    Color = 448
	ColorMustard                    Color = 449
	ColorMyrtle                     Color = 450
	ColorNadeshikoPink              Color = 451
	ColorNapierGreen                Color = 452
	ColorNaplesYellow               Color = 453
	ColorNavajoWhite                Color = 454
	ColorNavyBlue                   Color = 455
	ColorNeonCarrot                 Color = 456
	ColorNeonFuchsia                Color = 457
	ColorNeonGreen                  Color = 458
	ColorNonPhotoBlue               Color = 459
	ColorNorthTexasGreen            Color = 460
	ColorOceanBoatBlue              Color = 461
	ColorOchre                      Color = 462
	ColorOfficeGreen                Color = 463
	ColorOldGold                    Color = 464
	ColorOldLace                    Color = 465
	ColorOldLavender                Color = 466
	ColorOldMauve                   Color = 467
	ColorOldRose                    Color = 468
	ColorOlive                      Color = 469
	ColorOliveDrab                  Color = 470
	ColorOliveGreen                 Color = 471
	ColorOlivine                    Color = 472
	ColorOnyx                       Color = 473
	ColorOperaMauve                 Color = 474
	ColorOrange                     Color = 475
	ColorOrangeYellow               Color = 476
	ColorOrangePeel                 Color = 477
	ColorOrangeRed                  Color = 478
	ColorOrchid                     Color = 479
	ColorOtterBrown                 Color = 480
	ColorOuterSpace                 Color = 481
	ColorOutrageousOrange           Color = 482
	ColorOxfordBlue                 Color = 483
	ColorPacificBlue                Color = 484
	ColorPakistanGreen              Color = 485
	ColorPalatinateBlue             Color = 486
	ColorPalatinatePurple           Color = 487
	ColorPaleAqua                   Color = 488
	ColorPaleBlue                   Color = 489
	ColorPaleBrown                  Color = 490
	ColorPaleCarmine                Color = 491
	ColorPaleCerulean               Color = 492
	ColorPaleChestnut               Color = 493
	ColorPaleCopper                 Color = 494
	ColorPaleCornflowerBlue         Color = 495
	ColorPaleGold                   Color = 496
	ColorPaleGoldenrod              Color = 497
	ColorPaleGreen                  Color = 498
	ColorPaleLavender               Color = 499
	ColorPaleMagenta                Color = 500
	ColorPalePink                   Color = 501
	ColorPalePlum                   Color = 502
	ColorPaleRedViolet              Color = 503
	ColorPaleRobinEggBlue           Color = 504
	ColorPaleSilver                 Color = 505
	ColorPaleSpringBud              Color = 506
	ColorPaleTaupe                  Color = 507
	ColorPaleVioletRed              Color = 508
	ColorPansyPurple                Color = 509
	ColorPapayaWhip                 Color = 510
	ColorParisGreen                 Color = 511
	ColorPastelBlue                 Color = 512
	ColorPastelBrown                Color = 513
	ColorPastelGray                 Color = 514
	ColorPastelGreen                Color = 515
	ColorPastelMagenta              Color = 516
	ColorPastelOrange               Color = 517
	ColorPastelPink                 Color = 518
	ColorPastelPurple               Color = 519
	ColorPastelRed                  Color = 520
	ColorPastelViolet               Color = 521
	ColorPastelYellow               Color = 522
	ColorPatriarch                  Color = 523
	ColorPayneGrey                  Color = 524
	ColorPeach                      Color = 525
	ColorPeachPuff                  Color = 526
	ColorPeachYellow                Color = 527
	ColorPear                       Color = 528
	ColorPearl                      Color = 529
	ColorPearlAqua                  Color = 530
	ColorPeridot                    Color = 531
	ColorPeriwinkle                 Color = 532
	ColorPersianBlue                Color = 533
	ColorPersianIndigo              Color = 534
	ColorPersianOrange              Color = 535
	ColorPersianPink                Color = 536
	ColorPersianPlum                Color = 537
	ColorPersianRed                 Color = 538
	ColorPersianRose                Color = 539
	ColorPhlox                      Color = 540
	ColorPhthaloBlue                Color = 541
	ColorPhthaloGreen               Color = 542
	ColorPiggyPink                  Color = 543
	ColorPineGreen                  Color = 544
	ColorPink                       Color = 545
	ColorPinkFlamingo               Color = 546
	ColorPinkSherbet                Color = 547
	ColorPinkPearl                  Color = 548
	ColorPistachio                  Color = 549
	ColorPlatinum                   Color = 550
	ColorPlum                       Color = 551
	ColorPortlandOrange             Color = 552
	ColorPowderBlue                 Color = 553
	ColorPrincetonOrange            Color = 554
	ColorPrussianBlue               Color = 555
	ColorPsychedelicPurple          Color = 556
	ColorPuce                       Color = 557
	ColorPumpkin                    Color = 558
	ColorPurple                     Color = 559
	ColorPurpleHeart                Color = 560
	ColorPurpleMountainsMajesty     Color = 561
	ColorPurpleMountainMajesty      Color = 562
	ColorPurplePizzazz              Color = 563
	ColorPurpleTaupe                Color = 564
	ColorRackley                    Color = 565
	ColorRadicalRed                 Color = 566
	ColorRaspberry                  Color = 567
	ColorRaspberryGlace             Color = 568
	ColorRaspberryPink              Color = 569
	ColorRaspberryRose              Color = 570
	ColorRawSienna                  Color = 571
	ColorRazzleDazzleRose           Color = 572
	ColorRazzmatazz                 Color = 573
	ColorRed                        Color = 574
	ColorRedOrange                  Color = 575
	ColorRedBrown                   Color = 576
	ColorRedViolet                  Color = 577
	ColorRichBlack                  Color = 578
	ColorRichCarmine                Color = 579
	ColorRichElectricBlue           Color = 580
	ColorRichLilac                  Color = 581
	ColorRichMaroon                 Color = 582
	ColorRifleGreen                 Color = 583
	ColorRobinsEggBlue              Color = 584
	ColorRose                       Color = 585
	ColorRoseBonbon                 Color = 586
	ColorRoseEbony                  Color = 587
	ColorRoseGold                   Color = 588
	ColorRoseMadder                 Color = 589
	ColorRosePink                   Color = 590
	ColorRoseQuartz                 Color = 591
	ColorRoseTaupe                  Color = 592
	ColorRoseVale                   Color = 593
	ColorRosewood                   Color = 594
	ColorRossoCorsa                 Color = 595
	ColorRosyBrown                  Color = 596
	ColorRoyalAzure                 Color = 597
	ColorRoyalBlue                  Color = 598
	ColorRoyalFuchsia               Color = 599
	ColorRoyalPurple                Color = 600
	ColorRuby                       Color = 601
	ColorRuddy                      Color = 602
	ColorRuddyBrown                 Color = 603
	ColorRuddyPink                  Color = 604
	ColorRufous                     Color = 605
	ColorRusset                     Color = 606
	ColorRust                       Color = 607
	ColorSacramentoStateGreen       Color = 608
	ColorSaddleBrown                Color = 609
	ColorSafetyOrange               Color = 610
	ColorSaffron                    Color = 611
	ColorSaintPatrickBlue           Color = 612
	ColorSalmon                     Color = 613
	ColorSalmonPink                 Color = 614
	ColorSand                       Color = 615
	ColorSandDune                   Color = 616
	ColorSandstorm                  Color = 617
	ColorSandyBrown                 Color = 618
	ColorSandyTaupe                 Color = 619
	ColorSapGreen                   Color = 620
	ColorSapphire                   Color = 621
	ColorSatinSheenGold             Color = 622
	ColorScarlet                    Color = 623
	ColorSchoolBusYellow            Color = 624
	ColorScreaminGreen              Color = 625
	ColorSeaBlue                    Color = 626
	ColorSeaGreen                   Color = 627
	ColorSealBrown                  Color = 628
	ColorSeashell                   Color = 629
	ColorSelectiveYellow            Color = 630
	ColorSepia                      Color = 631
	ColorShadow                     Color = 632
	ColorShamrock                   Color = 633
	ColorShamrockGreen              Color = 634
	ColorShockingPink               Color = 635
	ColorSienna                     Color = 636
	ColorSilver                     Color = 637
	ColorSinopia                    Color = 638
	ColorSkobeloff                  Color = 639
	ColorSkyBlue                    Color = 640
	ColorSkyMagenta                 Color = 641
	ColorSlateBlue                  Color = 642
	ColorSlateGray                  Color = 643
	ColorSmalt                      Color = 644
	ColorSmokeyTopaz                Color = 645
	ColorSmokyBlack                 Color = 646
	ColorSnow                       Color = 647
	ColorSpiroDiscoBall             Color = 648
	ColorSpringBud                  Color = 649
	ColorSpringGreen                Color = 650
	ColorSteelBlue                  Color = 651
	ColorStilDeGrainYellow          Color = 652
	ColorStizza                     Color = 653
	ColorStormcloud                 Color = 654
	ColorStraw                      Color = 655
	ColorSunglow                    Color = 656
	ColorSunset                     Color = 657
	ColorSunsetOrange               Color = 658
	ColorTan                        Color = 659
	ColorTangelo                    Color = 660
	ColorTangerine                  Color = 661
	ColorTangerineYellow            Color = 662
	ColorTaupe                      Color = 663
	ColorTaupeGray                  Color = 664
	ColorTawny                      Color = 665
	ColorTeaGreen                   Color = 666
	ColorTeaRose                    Color = 667
	ColorTeal                       Color = 668
	ColorTealBlue                   Color = 669
	ColorTealGreen                  Color = 670
	ColorTerraCotta                 Color = 671
	ColorThistle                    Color = 672
	ColorThulianPink                Color = 673
	ColorTickleMePink               Color = 674
	ColorTiffanyBlue                Color = 675
	ColorTigerEye                   Color = 676
	ColorTimberwolf                 Color = 677
	ColorTitaniumYellow             Color = 678
	ColorTomato                     Color = 679
	ColorToolbox                    Color = 680
	ColorTopaz                      Color = 681
	ColorTractorRed                 Color = 682
	ColorTrolleyGrey                Color = 683
	ColorTropicalRainForest         Color = 684
	ColorTrueBlue                   Color = 685
	ColorTuftsBlue                  Color = 686
	ColorTumbleweed                 Color = 687
	ColorTurkishRose                Color = 688
	ColorTurquoise                  Color = 689
	ColorTurquoiseBlue              Color = 690
	ColorTurquoiseGreen             Color = 691
	ColorTuscanRed                  Color = 692
	ColorTwilightLavender           Color = 693
	ColorTyrianPurple               Color = 694
	ColorUaBlue                     Color = 695
	ColorUaRed                      Color = 696
	ColorUclaBlue                   Color = 697
	ColorUclaGold                   Color = 698
	ColorUfoGreen                   Color = 699
	ColorUpForestGreen              Color = 700
	ColorUpMaroon                   Color = 701
	ColorUscCardinal                Color = 702
	ColorUscGold                    Color = 703
	ColorUbe                        Color = 704
	ColorUltraPink                  Color = 705
	ColorUltramarine                Color = 706
	ColorUltramarineBlue            Color = 707
	ColorUmber                      Color = 708
	ColorUnitedNationsBlue          Color = 709
	ColorUniversityOfCaliforniaGold Color = 710
	ColorUnmellowYellow             Color = 711
	ColorUpsdellRed                 Color = 712
	ColorUrobilin                   Color = 713
	ColorUtahCrimson                Color = 714
	ColorVanilla                    Color = 715
	ColorVegasGold                  Color = 716
	ColorVenetianRed                Color = 717
	ColorVerdigris                  Color = 718
	ColorVermilion                  Color = 719
	ColorVeronica                   Color = 720
	ColorViolet                     Color = 721
	ColorVioletBlue                 Color = 722
	ColorVioletRed                  Color = 723
	ColorViridian                   Color = 724
	ColorVividAuburn                Color = 725
	ColorVividBurgundy              Color = 726
	ColorVividCerise                Color = 727
	ColorVividTangerine             Color = 728
	ColorVividViolet                Color = 729
	ColorWarmBlack                  Color = 730
	ColorWaterspout                 Color = 731
	ColorWenge                      Color = 732
	ColorWheat                      Color = 733
	ColorWhite                      Color = 734
	ColorWhiteSmoke                 Color = 735
	ColorWildStrawberry             Color = 736
	ColorWildWatermelon             Color = 737
	ColorWildBlueYonder             Color = 738
	ColorWine                       Color = 739
	ColorWisteria                   Color = 740
	ColorXanadu                     Color = 741
	ColorYaleBlue                   Color = 742
	ColorYellow                     Color = 743
	ColorYellowOrange               Color = 744
	ColorYellowGreen                Color = 745
	ColorZaffre                     Color = 746
	ColorZinnwalditeBrown           Color = 747
	ColorMetallic                   Color = 748
	ColorDarkPurple                 Color = 749
	ColorPewter                     Color = 750
	ColorBurn                       Color = 751
	ColorVinous                     Color = 752
	ColorTwoTone                    Color = 753
	ColorLightRed                   Color = 754
	ColorLightOrange                Color = 755
)

var ColorNames = map[int32]string{
	1:   "OTHER",
	2:   "AIR_FORCE_BLUE",
	3:   "ALICE_BLUE",
	4:   "ALIZARIN_CRIMSON",
	5:   "ALMOND",
	6:   "AMARANTH",
	7:   "AMBER",
	8:   "AMERICAN_ROSE",
	9:   "AMETHYST",
	10:  "ANDROID_GREEN",
	11:  "ANTI_FLASH_WHITE",
	12:  "ANTIQUE_BRASS",
	13:  "ANTIQUE_FUCHSIA",
	14:  "ANTIQUE_WHITE",
	15:  "AO",
	16:  "APPLE_GREEN",
	17:  "APRICOT",
	18:  "AQUA",
	19:  "AQUAMARINE",
	20:  "ARMY_GREEN",
	21:  "ARYLIDE_YELLOW",
	22:  "ASH_GREY",
	23:  "ASPARAGUS",
	24:  "ATOMIC_TANGERINE",
	25:  "AUBURN",
	26:  "AUREOLIN",
	27:  "AUROMETALSAURUS",
	28:  "AWESOME",
	29:  "AZURE",
	30:  "AZURE_MIST_WEB",
	31:  "BABY_BLUE",
	32:  "BABY_BLUE_EYES",
	33:  "BABY_PINK",
	34:  "BALL_BLUE",
	35:  "BANANA_MANIA",
	36:  "BANANA_YELLOW",
	37:  "BATTLESHIP_GREY",
	38:  "BAZAAR",
	39:  "BEAU_BLUE",
	40:  "BEAVER",
	41:  "BEIGE",
	42:  "BISQUE",
	43:  "BISTRE",
	44:  "BITTERSWEET",
	45:  "BLACK",
	46:  "BLANCHED_ALMOND",
	47:  "BLEU_DE_FRANCE",
	48:  "BLIZZARD_BLUE",
	49:  "BLOND",
	50:  "BLUE",
	51:  "BLUE_BELL",
	52:  "BLUE_GRAY",
	53:  "BLUE_GREEN",
	54:  "BLUE_PURPLE",
	55:  "BLUE_VIOLET",
	56:  "BLUSH",
	57:  "BOLE",
	58:  "BONDI_BLUE",
	59:  "BONE",
	60:  "BOSTON_UNIVERSITY_RED",
	61:  "BOTTLE_GREEN",
	62:  "BOYSENBERRY",
	63:  "BRANDEIS_BLUE",
	64:  "BRASS",
	65:  "BRICK_RED",
	66:  "BRIGHT_CERULEAN",
	67:  "BRIGHT_GREEN",
	68:  "BRIGHT_LAVENDER",
	69:  "BRIGHT_MAROON",
	70:  "BRIGHT_PINK",
	71:  "BRIGHT_TURQUOISE",
	72:  "BRIGHT_UBE",
	73:  "BRILLIANT_LAVENDER",
	74:  "BRILLIANT_ROSE",
	75:  "BRINK_PINK",
	76:  "BRITISH_RACING_GREEN",
	77:  "BRONZE",
	78:  "BROWN",
	79:  "BUBBLE_GUM",
	80:  "BUBBLES",
	81:  "BUFF",
	82:  "BULGARIAN_ROSE",
	83:  "BURGUNDY",
	84:  "BURLYWOOD",
	85:  "BURNT_ORANGE",
	86:  "BURNT_SIENNA",
	87:  "BURNT_UMBER",
	88:  "BYZANTINE",
	89:  "BYZANTIUM",
	90:  "CG_BLUE",
	91:  "CG_RED",
	92:  "CADET",
	93:  "CADET_BLUE",
	94:  "CADET_GREY",
	95:  "CADMIUM_GREEN",
	96:  "CADMIUM_ORANGE",
	97:  "CADMIUM_RED",
	98:  "CADMIUM_YELLOW",
	99:  "CAFE_AU_LAIT",
	100: "CAFE_NOIR",
	101: "CAL_POLY_POMONA_GREEN",
	102: "CAMBRIDGE_BLUE",
	103: "CAMEL",
	104: "CAMOUFLAGE_GREEN",
	105: "CANARY",
	106: "CANARY_YELLOW",
	107: "CANDY_APPLE_RED",
	108: "CANDY_PINK",
	109: "CAPRI",
	110: "CAPUT_MORTUUM",
	111: "CARDINAL",
	112: "CARIBBEAN_GREEN",
	113: "CARMINE",
	114: "CARMINE_PINK",
	115: "CARMINE_RED",
	116: "CARNATION_PINK",
	117: "CARNELIAN",
	118: "CAROLINA_BLUE",
	119: "CARROT_ORANGE",
	120: "CELADON",
	121: "CELESTE",
	122: "CELESTIAL_BLUE",
	123: "CERISE",
	124: "CERISE_PINK",
	125: "CERULEAN",
	126: "CERULEAN_BLUE",
	127: "CHAMOISEE",
	128: "CHAMPAGNE",
	129: "CHARCOAL",
	130: "CHARTREUSE",
	131: "CHERRY",
	132: "CHERRY_BLOSSOM_PINK",
	133: "CHESTNUT",
	134: "CHOCOLATE",
	135: "CHROME_YELLOW",
	136: "CINEREOUS",
	137: "CINNABAR",
	138: "CINNAMON",
	139: "CITRINE",
	140: "CLASSIC_ROSE",
	141: "COBALT",
	142: "COCOA_BROWN",
	143: "COFFEE",
	144: "COLUMBIA_BLUE",
	145: "COOL_BLACK",
	146: "COOL_GREY",
	147: "COPPER",
	148: "COPPER_ROSE",
	149: "COQUELICOT",
	150: "CORAL",
	151: "CORAL_PINK",
	152: "CORAL_RED",
	153: "CORDOVAN",
	154: "CORN",
	155: "CORNELL_RED",
	156: "CORNFLOWER",
	157: "CORNFLOWER_BLUE",
	158: "CORNSILK",
	159: "COSMIC_LATTE",
	160: "COTTON_CANDY",
	161: "CREAM",
	162: "CRIMSON",
	163: "CRIMSON_RED",
	164: "CRIMSON_GLORY",
	165: "CYAN",
	166: "DAFFODIL",
	167: "DANDELION",
	168: "DARK_BLUE",
	169: "DARK_BROWN",
	170: "DARK_BYZANTIUM",
	171: "DARK_CANDY_APPLE_RED",
	172: "DARK_CERULEAN",
	173: "DARK_CHESTNUT",
	174: "DARK_CORAL",
	175: "DARK_CYAN",
	176: "DARK_ELECTRIC_BLUE",
	177: "DARK_GOLDENROD",
	178: "DARK_GRAY",
	179: "DARK_GREEN",
	180: "DARK_JUNGLE_GREEN",
	181: "DARK_KHAKI",
	182: "DARK_LAVA",
	183: "DARK_LAVENDER",
	184: "DARK_MAGENTA",
	185: "DARK_MIDNIGHT_BLUE",
	186: "DARK_OLIVE_GREEN",
	187: "DARK_ORANGE",
	188: "DARK_ORCHID",
	189: "DARK_PASTEL_BLUE",
	190: "DARK_PASTEL_GREEN",
	191: "DARK_PASTEL_PURPLE",
	192: "DARK_PASTEL_RED",
	193: "DARK_PINK",
	194: "DARK_POWDER_BLUE",
	195: "DARK_RASPBERRY",
	196: "DARK_RED",
	197: "DARK_SALMON",
	198: "DARK_SCARLET",
	199: "DARK_SEA_GREEN",
	200: "DARK_SIENNA",
	201: "DARK_SLATE_BLUE",
	202: "DARK_SLATE_GRAY",
	203: "DARK_SPRING_GREEN",
	204: "DARK_TAN",
	205: "DARK_TANGERINE",
	206: "DARK_TAUPE",
	207: "DARK_TERRA_COTTA",
	208: "DARK_TURQUOISE",
	209: "DARK_VIOLET",
	210: "DARTMOUTH_GREEN",
	211: "DAVY_GREY",
	212: "DEBIAN_RED",
	213: "DEEP_CARMINE",
	214: "DEEP_CARMINE_PINK",
	215: "DEEP_CARROT_ORANGE",
	216: "DEEP_CERISE",
	217: "DEEP_CHAMPAGNE",
	218: "DEEP_CHESTNUT",
	219: "DEEP_COFFEE",
	220: "DEEP_FUCHSIA",
	221: "DEEP_JUNGLE_GREEN",
	222: "DEEP_LILAC",
	223: "DEEP_MAGENTA",
	224: "DEEP_PEACH",
	225: "DEEP_PINK",
	226: "DEEP_SAFFRON",
	227: "DEEP_SKY_BLUE",
	228: "DENIM",
	229: "DESERT",
	230: "DESERT_SAND",
	231: "DIM_GRAY",
	232: "DODGER_BLUE",
	233: "DOGWOOD_ROSE",
	234: "DOLLAR_BILL",
	235: "DRAB",
	236: "DUKE_BLUE",
	237: "EARTH_YELLOW",
	238: "ECRU",
	239: "EGGPLANT",
	240: "EGGSHELL",
	241: "EGYPTIAN_BLUE",
	242: "ELECTRIC_BLUE",
	243: "ELECTRIC_CRIMSON",
	244: "ELECTRIC_CYAN",
	245: "ELECTRIC_GREEN",
	246: "ELECTRIC_INDIGO",
	247: "ELECTRIC_LAVENDER",
	248: "ELECTRIC_LIME",
	249: "ELECTRIC_PURPLE",
	250: "ELECTRIC_ULTRAMARINE",
	251: "ELECTRIC_VIOLET",
	252: "ELECTRIC_YELLOW",
	253: "EMERALD",
	254: "ETON_BLUE",
	255: "FALLOW",
	256: "FALU_RED",
	257: "FAMOUS",
	258: "FANDANGO",
	259: "FASHION_FUCHSIA",
	260: "FAWN",
	261: "FELDGRAU",
	262: "FERN",
	263: "FERN_GREEN",
	264: "FERRARI_RED",
	265: "FIELD_DRAB",
	266: "FIRE_ENGINE_RED",
	267: "FIREBRICK",
	268: "FLAME",
	269: "FLAMINGO_PINK",
	270: "FLAVESCENT",
	271: "FLAX",
	272: "FLORAL_WHITE",
	273: "FLUORESCENT_ORANGE",
	274: "FLUORESCENT_PINK",
	275: "FLUORESCENT_YELLOW",
	276: "FOLLY",
	277: "FOREST_GREEN",
	278: "FRENCH_BEIGE",
	279: "FRENCH_BLUE",
	280: "FRENCH_LILAC",
	281: "FRENCH_ROSE",
	282: "FUCHSIA",
	283: "FUCHSIA_PINK",
	284: "FULVOUS",
	285: "FUZZY_WUZZY",
	286: "GAINSBORO",
	287: "GAMBOGE",
	288: "GHOST_WHITE",
	289: "GINGER",
	290: "GLAUCOUS",
	291: "GLITTER",
	292: "GOLD",
	293: "GOLDEN_BROWN",
	294: "GOLDEN_POPPY",
	295: "GOLDEN_YELLOW",
	296: "GOLDENROD",
	297: "GRANNY_SMITH_APPLE",
	298: "GRAY",
	299: "GRAY_ASPARAGUS",
	300: "GREEN",
	301: "GREEN_BLUE",
	302: "GREEN_YELLOW",
	303: "GRULLO",
	304: "GUPPIE_GREEN",
	305: "HALAYA_UBE",
	306: "HAN_BLUE",
	307: "HAN_PURPLE",
	308: "HANSA_YELLOW",
	309: "HARLEQUIN",
	310: "HARVARD_CRIMSON",
	311: "HARVEST_GOLD",
	312: "HEART_GOLD",
	313: "HELIOTROPE",
	314: "HOLLYWOOD_CERISE",
	315: "HONEYDEW",
	316: "HOOKER_GREEN",
	317: "HOT_MAGENTA",
	318: "HOT_PINK",
	319: "HUNTER_GREEN",
	320: "ICTERINE",
	321: "INCHWORM",
	322: "INDIA_GREEN",
	323: "INDIAN_RED",
	324: "INDIAN_YELLOW",
	325: "INDIGO",
	326: "INTERNATIONAL_KLEIN_BLUE",
	327: "INTERNATIONAL_ORANGE",
	328: "IRIS",
	329: "ISABELLINE",
	330: "ISLAMIC_GREEN",
	331: "IVORY",
	332: "JADE",
	333: "JASMINE",
	334: "JASPER",
	335: "JAZZBERRY_JAM",
	336: "JONQUIL",
	337: "JUNE_BUD",
	338: "JUNGLE_GREEN",
	339: "KU_CRIMSON",
	340: "KELLY_GREEN",
	341: "KHAKI",
	342: "LA_SALLE_GREEN",
	343: "LANGUID_LAVENDER",
	344: "LAPIS_LAZULI",
	345: "LASER_LEMON",
	346: "LAUREL_GREEN",
	347: "LAVA",
	348: "LAVENDER",
	349: "LAVENDER_BLUE",
	350: "LAVENDER_BLUSH",
	351: "LAVENDER_GRAY",
	352: "LAVENDER_INDIGO",
	353: "LAVENDER_MAGENTA",
	354: "LAVENDER_MIST",
	355: "LAVENDER_PINK",
	356: "LAVENDER_PURPLE",
	357: "LAVENDER_ROSE",
	358: "LAWN_GREEN",
	359: "LEMON",
	360: "LEMON_YELLOW",
	361: "LEMON_CHIFFON",
	362: "LEMON_LIME",
	363: "LIGHT_CRIMSON",
	364: "LIGHT_THULIAN_PINK",
	365: "LIGHT_APRICOT",
	366: "LIGHT_BLUE",
	367: "LIGHT_BROWN",
	368: "LIGHT_CARMINE_PINK",
	369: "LIGHT_CORAL",
	370: "LIGHT_CORNFLOWER_BLUE",
	371: "LIGHT_CYAN",
	372: "LIGHT_FUCHSIA_PINK",
	373: "LIGHT_GOLDENROD_YELLOW",
	374: "LIGHT_GRAY",
	375: "LIGHT_GREEN",
	376: "LIGHT_KHAKI",
	377: "LIGHT_PASTEL_PURPLE",
	378: "LIGHT_PINK",
	379: "LIGHT_SALMON",
	380: "LIGHT_SALMON_PINK",
	381: "LIGHT_SEA_GREEN",
	382: "LIGHT_SKY_BLUE",
	383: "LIGHT_SLATE_GRAY",
	384: "LIGHT_TAUPE",
	385: "LIGHT_YELLOW",
	386: "LILAC",
	387: "LIME",
	388: "LIME_GREEN",
	389: "LINCOLN_GREEN",
	390: "LINEN",
	391: "LION",
	392: "LIVER",
	393: "LUST",
	394: "MSU_GREEN",
	395: "MACARONI_AND_CHEESE",
	396: "MAGENTA",
	397: "MAGIC_MINT",
	398: "MAGNOLIA",
	399: "MAHOGANY",
	400: "MAIZE",
	401: "MAJORELLE_BLUE",
	402: "MALACHITE",
	403: "MANATEE",
	404: "MANGO_TANGO",
	405: "MANTIS",
	406: "MAROON",
	407: "MAUVE",
	408: "MAUVE_TAUPE",
	409: "MAUVELOUS",
	410: "MAYA_BLUE",
	411: "MEAT_BROWN",
	412: "MEDIUM_PERSIAN_BLUE",
	413: "MEDIUM_AQUAMARINE",
	414: "MEDIUM_BLUE",
	415: "MEDIUM_CANDY_APPLE_RED",
	416: "MEDIUM_CARMINE",
	417: "MEDIUM_CHAMPAGNE",
	418: "MEDIUM_ELECTRIC_BLUE",
	419: "MEDIUM_JUNGLE_GREEN",
	420: "MEDIUM_LAVENDER_MAGENTA",
	421: "MEDIUM_ORCHID",
	422: "MEDIUM_PURPLE",
	423: "MEDIUM_RED_VIOLET",
	424: "MEDIUM_SEA_GREEN",
	425: "MEDIUM_SLATE_BLUE",
	426: "MEDIUM_SPRING_BUD",
	427: "MEDIUM_SPRING_GREEN",
	428: "MEDIUM_TAUPE",
	429: "MEDIUM_TEAL_BLUE",
	430: "MEDIUM_TURQUOISE",
	431: "MEDIUM_VIOLET_RED",
	432: "MELON",
	433: "MIDNIGHT_BLUE",
	434: "MIDNIGHT_GREEN",
	435: "MIKADO_YELLOW",
	436: "MINT",
	437: "MINT_CREAM",
	438: "MINT_GREEN",
	439: "MISTY_ROSE",
	440: "MOCCASIN",
	441: "MODE_BEIGE",
	442: "MOONSTONE_BLUE",
	443: "MORDANT_RED_19",
	444: "MOSS_GREEN",
	445: "MOUNTAIN_MEADOW",
	446: "MOUNTBATTEN_PINK",
	447: "MULBERRY",
	448: "MUNSELL",
	449: "MUSTARD",
	450: "MYRTLE",
	451: "NADESHIKO_PINK",
	452: "NAPIER_GREEN",
	453: "NAPLES_YELLOW",
	454: "NAVAJO_WHITE",
	455: "NAVY_BLUE",
	456: "NEON_CARROT",
	457: "NEON_FUCHSIA",
	458: "NEON_GREEN",
	459: "NON_PHOTO_BLUE",
	460: "NORTH_TEXAS_GREEN",
	461: "OCEAN_BOAT_BLUE",
	462: "OCHRE",
	463: "OFFICE_GREEN",
	464: "OLD_GOLD",
	465: "OLD_LACE",
	466: "OLD_LAVENDER",
	467: "OLD_MAUVE",
	468: "OLD_ROSE",
	469: "OLIVE",
	470: "OLIVE_DRAB",
	471: "OLIVE_GREEN",
	472: "OLIVINE",
	473: "ONYX",
	474: "OPERA_MAUVE",
	475: "ORANGE",
	476: "ORANGE_YELLOW",
	477: "ORANGE_PEEL",
	478: "ORANGE_RED",
	479: "ORCHID",
	480: "OTTER_BROWN",
	481: "OUTER_SPACE",
	482: "OUTRAGEOUS_ORANGE",
	483: "OXFORD_BLUE",
	484: "PACIFIC_BLUE",
	485: "PAKISTAN_GREEN",
	486: "PALATINATE_BLUE",
	487: "PALATINATE_PURPLE",
	488: "PALE_AQUA",
	489: "PALE_BLUE",
	490: "PALE_BROWN",
	491: "PALE_CARMINE",
	492: "PALE_CERULEAN",
	493: "PALE_CHESTNUT",
	494: "PALE_COPPER",
	495: "PALE_CORNFLOWER_BLUE",
	496: "PALE_GOLD",
	497: "PALE_GOLDENROD",
	498: "PALE_GREEN",
	499: "PALE_LAVENDER",
	500: "PALE_MAGENTA",
	501: "PALE_PINK",
	502: "PALE_PLUM",
	503: "PALE_RED_VIOLET",
	504: "PALE_ROBIN_EGG_BLUE",
	505: "PALE_SILVER",
	506: "PALE_SPRING_BUD",
	507: "PALE_TAUPE",
	508: "PALE_VIOLET_RED",
	509: "PANSY_PURPLE",
	510: "PAPAYA_WHIP",
	511: "PARIS_GREEN",
	512: "PASTEL_BLUE",
	513: "PASTEL_BROWN",
	514: "PASTEL_GRAY",
	515: "PASTEL_GREEN",
	516: "PASTEL_MAGENTA",
	517: "PASTEL_ORANGE",
	518: "PASTEL_PINK",
	519: "PASTEL_PURPLE",
	520: "PASTEL_RED",
	521: "PASTEL_VIOLET",
	522: "PASTEL_YELLOW",
	523: "PATRIARCH",
	524: "PAYNE_GREY",
	525: "PEACH",
	526: "PEACH_PUFF",
	527: "PEACH_YELLOW",
	528: "PEAR",
	529: "PEARL",
	530: "PEARL_AQUA",
	531: "PERIDOT",
	532: "PERIWINKLE",
	533: "PERSIAN_BLUE",
	534: "PERSIAN_INDIGO",
	535: "PERSIAN_ORANGE",
	536: "PERSIAN_PINK",
	537: "PERSIAN_PLUM",
	538: "PERSIAN_RED",
	539: "PERSIAN_ROSE",
	540: "PHLOX",
	541: "PHTHALO_BLUE",
	542: "PHTHALO_GREEN",
	543: "PIGGY_PINK",
	544: "PINE_GREEN",
	545: "PINK",
	546: "PINK_FLAMINGO",
	547: "PINK_SHERBET",
	548: "PINK_PEARL",
	549: "PISTACHIO",
	550: "PLATINUM",
	551: "PLUM",
	552: "PORTLAND_ORANGE",
	553: "POWDER_BLUE",
	554: "PRINCETON_ORANGE",
	555: "PRUSSIAN_BLUE",
	556: "PSYCHEDELIC_PURPLE",
	557: "PUCE",
	558: "PUMPKIN",
	559: "PURPLE",
	560: "PURPLE_HEART",
	561: "PURPLE_MOUNTAINS_MAJESTY",
	562: "PURPLE_MOUNTAIN_MAJESTY",
	563: "PURPLE_PIZZAZZ",
	564: "PURPLE_TAUPE",
	565: "RACKLEY",
	566: "RADICAL_RED",
	567: "RASPBERRY",
	568: "RASPBERRY_GLACE",
	569: "RASPBERRY_PINK",
	570: "RASPBERRY_ROSE",
	571: "RAW_SIENNA",
	572: "RAZZLE_DAZZLE_ROSE",
	573: "RAZZMATAZZ",
	574: "RED",
	575: "RED_ORANGE",
	576: "RED_BROWN",
	577: "RED_VIOLET",
	578: "RICH_BLACK",
	579: "RICH_CARMINE",
	580: "RICH_ELECTRIC_BLUE",
	581: "RICH_LILAC",
	582: "RICH_MAROON",
	583: "RIFLE_GREEN",
	584: "ROBINS_EGG_BLUE",
	585: "ROSE",
	586: "ROSE_BONBON",
	587: "ROSE_EBONY",
	588: "ROSE_GOLD",
	589: "ROSE_MADDER",
	590: "ROSE_PINK",
	591: "ROSE_QUARTZ",
	592: "ROSE_TAUPE",
	593: "ROSE_VALE",
	594: "ROSEWOOD",
	595: "ROSSO_CORSA",
	596: "ROSY_BROWN",
	597: "ROYAL_AZURE",
	598: "ROYAL_BLUE",
	599: "ROYAL_FUCHSIA",
	600: "ROYAL_PURPLE",
	601: "RUBY",
	602: "RUDDY",
	603: "RUDDY_BROWN",
	604: "RUDDY_PINK",
	605: "RUFOUS",
	606: "RUSSET",
	607: "RUST",
	608: "SACRAMENTO_STATE_GREEN",
	609: "SADDLE_BROWN",
	610: "SAFETY_ORANGE",
	611: "SAFFRON",
	612: "SAINT_PATRICK_BLUE",
	613: "SALMON",
	614: "SALMON_PINK",
	615: "SAND",
	616: "SAND_DUNE",
	617: "SANDSTORM",
	618: "SANDY_BROWN",
	619: "SANDY_TAUPE",
	620: "SAP_GREEN",
	621: "SAPPHIRE",
	622: "SATIN_SHEEN_GOLD",
	623: "SCARLET",
	624: "SCHOOL_BUS_YELLOW",
	625: "SCREAMIN_GREEN",
	626: "SEA_BLUE",
	627: "SEA_GREEN",
	628: "SEAL_BROWN",
	629: "SEASHELL",
	630: "SELECTIVE_YELLOW",
	631: "SEPIA",
	632: "SHADOW",
	633: "SHAMROCK",
	634: "SHAMROCK_GREEN",
	635: "SHOCKING_PINK",
	636: "SIENNA",
	637: "SILVER",
	638: "SINOPIA",
	639: "SKOBELOFF",
	640: "SKY_BLUE",
	641: "SKY_MAGENTA",
	642: "SLATE_BLUE",
	643: "SLATE_GRAY",
	644: "SMALT",
	645: "SMOKEY_TOPAZ",
	646: "SMOKY_BLACK",
	647: "SNOW",
	648: "SPIRO_DISCO_BALL",
	649: "SPRING_BUD",
	650: "SPRING_GREEN",
	651: "STEEL_BLUE",
	652: "STIL_DE_GRAIN_YELLOW",
	653: "STIZZA",
	654: "STORMCLOUD",
	655: "STRAW",
	656: "SUNGLOW",
	657: "SUNSET",
	658: "SUNSET_ORANGE",
	659: "TAN",
	660: "TANGELO",
	661: "TANGERINE",
	662: "TANGERINE_YELLOW",
	663: "TAUPE",
	664: "TAUPE_GRAY",
	665: "TAWNY",
	666: "TEA_GREEN",
	667: "TEA_ROSE",
	668: "TEAL",
	669: "TEAL_BLUE",
	670: "TEAL_GREEN",
	671: "TERRA_COTTA",
	672: "THISTLE",
	673: "THULIAN_PINK",
	674: "TICKLE_ME_PINK",
	675: "TIFFANY_BLUE",
	676: "TIGER_EYE",
	677: "TIMBERWOLF",
	678: "TITANIUM_YELLOW",
	679: "TOMATO",
	680: "TOOLBOX",
	681: "TOPAZ",
	682: "TRACTOR_RED",
	683: "TROLLEY_GREY",
	684: "TROPICAL_RAIN_FOREST",
	685: "TRUE_BLUE",
	686: "TUFTS_BLUE",
	687: "TUMBLEWEED",
	688: "TURKISH_ROSE",
	689: "TURQUOISE",
	690: "TURQUOISE_BLUE",
	691: "TURQUOISE_GREEN",
	692: "TUSCAN_RED",
	693: "TWILIGHT_LAVENDER",
	694: "TYRIAN_PURPLE",
	695: "UA_BLUE",
	696: "UA_RED",
	697: "UCLA_BLUE",
	698: "UCLA_GOLD",
	699: "UFO_GREEN",
	700: "UP_FOREST_GREEN",
	701: "UP_MAROON",
	702: "USC_CARDINAL",
	703: "USC_GOLD",
	704: "UBE",
	705: "ULTRA_PINK",
	706: "ULTRAMARINE",
	707: "ULTRAMARINE_BLUE",
	708: "UMBER",
	709: "UNITED_NATIONS_BLUE",
	710: "UNIVERSITY_OF_CALIFORNIA_GOLD",
	711: "UNMELLOW_YELLOW",
	712: "UPSDELL_RED",
	713: "UROBILIN",
	714: "UTAH_CRIMSON",
	715: "VANILLA",
	716: "VEGAS_GOLD",
	717: "VENETIAN_RED",
	718: "VERDIGRIS",
	719: "VERMILION",
	720: "VERONICA",
	721: "VIOLET",
	722: "VIOLET_BLUE",
	723: "VIOLET_RED",
	724: "VIRIDIAN",
	725: "VIVID_AUBURN",
	726: "VIVID_BURGUNDY",
	727: "VIVID_CERISE",
	728: "VIVID_TANGERINE",
	729: "VIVID_VIOLET",
	730: "WARM_BLACK",
	731: "WATERSPOUT",
	732: "WENGE",
	733: "WHEAT",
	734: "WHITE",
	735: "WHITE_SMOKE",
	736: "WILD_STRAWBERRY",
	737: "WILD_WATERMELON",
	738: "WILD_BLUE_YONDER",
	739: "WINE",
	740: "WISTERIA",
	741: "XANADU",
	742: "YALE_BLUE",
	743: "YELLOW",
	744: "YELLOW_ORANGE",
	745: "YELLOW_GREEN",
	746: "ZAFFRE",
	747: "ZINNWALDITE_BROWN",
	748: "METALLIC",
	749: "DARK_PURPLE",
	750: "PEWTER",
	751: "BURN",
	752: "VINOUS",
	753: "TWO_TONE",
	754: "LIGHT_RED",
	755: "LIGHT_ORANGE",
}

var ColorValues = map[string]int32{
	"OTHER":                         1,
	"AIR_FORCE_BLUE":                2,
	"ALICE_BLUE":                    3,
	"ALIZARIN_CRIMSON":              4,
	"ALMOND":                        5,
	"AMARANTH":                      6,
	"AMBER":                         7,
	"AMERICAN_ROSE":                 8,
	"AMETHYST":                      9,
	"ANDROID_GREEN":                 10,
	"ANTI_FLASH_WHITE":              11,
	"ANTIQUE_BRASS":                 12,
	"ANTIQUE_FUCHSIA":               13,
	"ANTIQUE_WHITE":                 14,
	"AO":                            15,
	"APPLE_GREEN":                   16,
	"APRICOT":                       17,
	"AQUA":                          18,
	"AQUAMARINE":                    19,
	"ARMY_GREEN":                    20,
	"ARYLIDE_YELLOW":                21,
	"ASH_GREY":                      22,
	"ASPARAGUS":                     23,
	"ATOMIC_TANGERINE":              24,
	"AUBURN":                        25,
	"AUREOLIN":                      26,
	"AUROMETALSAURUS":               27,
	"AWESOME":                       28,
	"AZURE":                         29,
	"AZURE_MIST_WEB":                30,
	"BABY_BLUE":                     31,
	"BABY_BLUE_EYES":                32,
	"BABY_PINK":                     33,
	"BALL_BLUE":                     34,
	"BANANA_MANIA":                  35,
	"BANANA_YELLOW":                 36,
	"BATTLESHIP_GREY":               37,
	"BAZAAR":                        38,
	"BEAU_BLUE":                     39,
	"BEAVER":                        40,
	"BEIGE":                         41,
	"BISQUE":                        42,
	"BISTRE":                        43,
	"BITTERSWEET":                   44,
	"BLACK":                         45,
	"BLANCHED_ALMOND":               46,
	"BLEU_DE_FRANCE":                47,
	"BLIZZARD_BLUE":                 48,
	"BLOND":                         49,
	"BLUE":                          50,
	"BLUE_BELL":                     51,
	"BLUE_GRAY":                     52,
	"BLUE_GREEN":                    53,
	"BLUE_PURPLE":                   54,
	"BLUE_VIOLET":                   55,
	"BLUSH":                         56,
	"BOLE":                          57,
	"BONDI_BLUE":                    58,
	"BONE":                          59,
	"BOSTON_UNIVERSITY_RED":         60,
	"BOTTLE_GREEN":                  61,
	"BOYSENBERRY":                   62,
	"BRANDEIS_BLUE":                 63,
	"BRASS":                         64,
	"BRICK_RED":                     65,
	"BRIGHT_CERULEAN":               66,
	"BRIGHT_GREEN":                  67,
	"BRIGHT_LAVENDER":               68,
	"BRIGHT_MAROON":                 69,
	"BRIGHT_PINK":                   70,
	"BRIGHT_TURQUOISE":              71,
	"BRIGHT_UBE":                    72,
	"BRILLIANT_LAVENDER":            73,
	"BRILLIANT_ROSE":                74,
	"BRINK_PINK":                    75,
	"BRITISH_RACING_GREEN":          76,
	"BRONZE":                        77,
	"BROWN":                         78,
	"BUBBLE_GUM":                    79,
	"BUBBLES":                       80,
	"BUFF":                          81,
	"BULGARIAN_ROSE":                82,
	"BURGUNDY":                      83,
	"BURLYWOOD":                     84,
	"BURNT_ORANGE":                  85,
	"BURNT_SIENNA":                  86,
	"BURNT_UMBER":                   87,
	"BYZANTINE":                     88,
	"BYZANTIUM":                     89,
	"CG_BLUE":                       90,
	"CG_RED":                        91,
	"CADET":                         92,
	"CADET_BLUE":                    93,
	"CADET_GREY":                    94,
	"CADMIUM_GREEN":                 95,
	"CADMIUM_ORANGE":                96,
	"CADMIUM_RED":                   97,
	"CADMIUM_YELLOW":                98,
	"CAFE_AU_LAIT":                  99,
	"CAFE_NOIR":                     100,
	"CAL_POLY_POMONA_GREEN":         101,
	"CAMBRIDGE_BLUE":                102,
	"CAMEL":                         103,
	"CAMOUFLAGE_GREEN":              104,
	"CANARY":                        105,
	"CANARY_YELLOW":                 106,
	"CANDY_APPLE_RED":               107,
	"CANDY_PINK":                    108,
	"CAPRI":                         109,
	"CAPUT_MORTUUM":                 110,
	"CARDINAL":                      111,
	"CARIBBEAN_GREEN":               112,
	"CARMINE":                       113,
	"CARMINE_PINK":                  114,
	"CARMINE_RED":                   115,
	"CARNATION_PINK":                116,
	"CARNELIAN":                     117,
	"CAROLINA_BLUE":                 118,
	"CARROT_ORANGE":                 119,
	"CELADON":                       120,
	"CELESTE":                       121,
	"CELESTIAL_BLUE":                122,
	"CERISE":                        123,
	"CERISE_PINK":                   124,
	"CERULEAN":                      125,
	"CERULEAN_BLUE":                 126,
	"CHAMOISEE":                     127,
	"CHAMPAGNE":                     128,
	"CHARCOAL":                      129,
	"CHARTREUSE":                    130,
	"CHERRY":                        131,
	"CHERRY_BLOSSOM_PINK":           132,
	"CHESTNUT":                      133,
	"CHOCOLATE":                     134,
	"CHROME_YELLOW":                 135,
	"CINEREOUS":                     136,
	"CINNABAR":                      137,
	"CINNAMON":                      138,
	"CITRINE":                       139,
	"CLASSIC_ROSE":                  140,
	"COBALT":                        141,
	"COCOA_BROWN":                   142,
	"COFFEE":                        143,
	"COLUMBIA_BLUE":                 144,
	"COOL_BLACK":                    145,
	"COOL_GREY":                     146,
	"COPPER":                        147,
	"COPPER_ROSE":                   148,
	"COQUELICOT":                    149,
	"CORAL":                         150,
	"CORAL_PINK":                    151,
	"CORAL_RED":                     152,
	"CORDOVAN":                      153,
	"CORN":                          154,
	"CORNELL_RED":                   155,
	"CORNFLOWER":                    156,
	"CORNFLOWER_BLUE":               157,
	"CORNSILK":                      158,
	"COSMIC_LATTE":                  159,
	"COTTON_CANDY":                  160,
	"CREAM":                         161,
	"CRIMSON":                       162,
	"CRIMSON_RED":                   163,
	"CRIMSON_GLORY":                 164,
	"CYAN":                          165,
	"DAFFODIL":                      166,
	"DANDELION":                     167,
	"DARK_BLUE":                     168,
	"DARK_BROWN":                    169,
	"DARK_BYZANTIUM":                170,
	"DARK_CANDY_APPLE_RED":          171,
	"DARK_CERULEAN":                 172,
	"DARK_CHESTNUT":                 173,
	"DARK_CORAL":                    174,
	"DARK_CYAN":                     175,
	"DARK_ELECTRIC_BLUE":            176,
	"DARK_GOLDENROD":                177,
	"DARK_GRAY":                     178,
	"DARK_GREEN":                    179,
	"DARK_JUNGLE_GREEN":             180,
	"DARK_KHAKI":                    181,
	"DARK_LAVA":                     182,
	"DARK_LAVENDER":                 183,
	"DARK_MAGENTA":                  184,
	"DARK_MIDNIGHT_BLUE":            185,
	"DARK_OLIVE_GREEN":              186,
	"DARK_ORANGE":                   187,
	"DARK_ORCHID":                   188,
	"DARK_PASTEL_BLUE":              189,
	"DARK_PASTEL_GREEN":             190,
	"DARK_PASTEL_PURPLE":            191,
	"DARK_PASTEL_RED":               192,
	"DARK_PINK":                     193,
	"DARK_POWDER_BLUE":              194,
	"DARK_RASPBERRY":                195,
	"DARK_RED":                      196,
	"DARK_SALMON":                   197,
	"DARK_SCARLET":                  198,
	"DARK_SEA_GREEN":                199,
	"DARK_SIENNA":                   200,
	"DARK_SLATE_BLUE":               201,
	"DARK_SLATE_GRAY":               202,
	"DARK_SPRING_GREEN":             203,
	"DARK_TAN":                      204,
	"DARK_TANGERINE":                205,
	"DARK_TAUPE":                    206,
	"DARK_TERRA_COTTA":              207,
	"DARK_TURQUOISE":                208,
	"DARK_VIOLET":                   209,
	"DARTMOUTH_GREEN":               210,
	"DAVY_GREY":                     211,
	"DEBIAN_RED":                    212,
	"DEEP_CARMINE":                  213,
	"DEEP_CARMINE_PINK":             214,
	"DEEP_CARROT_ORANGE":            215,
	"DEEP_CERISE":                   216,
	"DEEP_CHAMPAGNE":                217,
	"DEEP_CHESTNUT":                 218,
	"DEEP_COFFEE":                   219,
	"DEEP_FUCHSIA":                  220,
	"DEEP_JUNGLE_GREEN":             221,
	"DEEP_LILAC":                    222,
	"DEEP_MAGENTA":                  223,
	"DEEP_PEACH":                    224,
	"DEEP_PINK":                     225,
	"DEEP_SAFFRON":                  226,
	"DEEP_SKY_BLUE":                 227,
	"DENIM":                         228,
	"DESERT":                        229,
	"DESERT_SAND":                   230,
	"DIM_GRAY":                      231,
	"DODGER_BLUE":                   232,
	"DOGWOOD_ROSE":                  233,
	"DOLLAR_BILL":                   234,
	"DRAB":                          235,
	"DUKE_BLUE":                     236,
	"EARTH_YELLOW":                  237,
	"ECRU":                          238,
	"EGGPLANT":                      239,
	"EGGSHELL":                      240,
	"EGYPTIAN_BLUE":                 241,
	"ELECTRIC_BLUE":                 242,
	"ELECTRIC_CRIMSON":              243,
	"ELECTRIC_CYAN":                 244,
	"ELECTRIC_GREEN":                245,
	"ELECTRIC_INDIGO":               246,
	"ELECTRIC_LAVENDER":             247,
	"ELECTRIC_LIME":                 248,
	"ELECTRIC_PURPLE":               249,
	"ELECTRIC_ULTRAMARINE":          250,
	"ELECTRIC_VIOLET":               251,
	"ELECTRIC_YELLOW":               252,
	"EMERALD":                       253,
	"ETON_BLUE":                     254,
	"FALLOW":                        255,
	"FALU_RED":                      256,
	"FAMOUS":                        257,
	"FANDANGO":                      258,
	"FASHION_FUCHSIA":               259,
	"FAWN":                          260,
	"FELDGRAU":                      261,
	"FERN":                          262,
	"FERN_GREEN":                    263,
	"FERRARI_RED":                   264,
	"FIELD_DRAB":                    265,
	"FIRE_ENGINE_RED":               266,
	"FIREBRICK":                     267,
	"FLAME":                         268,
	"FLAMINGO_PINK":                 269,
	"FLAVESCENT":                    270,
	"FLAX":                          271,
	"FLORAL_WHITE":                  272,
	"FLUORESCENT_ORANGE":            273,
	"FLUORESCENT_PINK":              274,
	"FLUORESCENT_YELLOW":            275,
	"FOLLY":                         276,
	"FOREST_GREEN":                  277,
	"FRENCH_BEIGE":                  278,
	"FRENCH_BLUE":                   279,
	"FRENCH_LILAC":                  280,
	"FRENCH_ROSE":                   281,
	"FUCHSIA":                       282,
	"FUCHSIA_PINK":                  283,
	"FULVOUS":                       284,
	"FUZZY_WUZZY":                   285,
	"GAINSBORO":                     286,
	"GAMBOGE":                       287,
	"GHOST_WHITE":                   288,
	"GINGER":                        289,
	"GLAUCOUS":                      290,
	"GLITTER":                       291,
	"GOLD":                          292,
	"GOLDEN_BROWN":                  293,
	"GOLDEN_POPPY":                  294,
	"GOLDEN_YELLOW":                 295,
	"GOLDENROD":                     296,
	"GRANNY_SMITH_APPLE":            297,
	"GRAY":                          298,
	"GRAY_ASPARAGUS":                299,
	"GREEN":                         300,
	"GREEN_BLUE":                    301,
	"GREEN_YELLOW":                  302,
	"GRULLO":                        303,
	"GUPPIE_GREEN":                  304,
	"HALAYA_UBE":                    305,
	"HAN_BLUE":                      306,
	"HAN_PURPLE":                    307,
	"HANSA_YELLOW":                  308,
	"HARLEQUIN":                     309,
	"HARVARD_CRIMSON":               310,
	"HARVEST_GOLD":                  311,
	"HEART_GOLD":                    312,
	"HELIOTROPE":                    313,
	"HOLLYWOOD_CERISE":              314,
	"HONEYDEW":                      315,
	"HOOKER_GREEN":                  316,
	"HOT_MAGENTA":                   317,
	"HOT_PINK":                      318,
	"HUNTER_GREEN":                  319,
	"ICTERINE":                      320,
	"INCHWORM":                      321,
	"INDIA_GREEN":                   322,
	"INDIAN_RED":                    323,
	"INDIAN_YELLOW":                 324,
	"INDIGO":                        325,
	"INTERNATIONAL_KLEIN_BLUE":      326,
	"INTERNATIONAL_ORANGE":          327,
	"IRIS":                          328,
	"ISABELLINE":                    329,
	"ISLAMIC_GREEN":                 330,
	"IVORY":                         331,
	"JADE":                          332,
	"JASMINE":                       333,
	"JASPER":                        334,
	"JAZZBERRY_JAM":                 335,
	"JONQUIL":                       336,
	"JUNE_BUD":                      337,
	"JUNGLE_GREEN":                  338,
	"KU_CRIMSON":                    339,
	"KELLY_GREEN":                   340,
	"KHAKI":                         341,
	"LA_SALLE_GREEN":                342,
	"LANGUID_LAVENDER":              343,
	"LAPIS_LAZULI":                  344,
	"LASER_LEMON":                   345,
	"LAUREL_GREEN":                  346,
	"LAVA":                          347,
	"LAVENDER":                      348,
	"LAVENDER_BLUE":                 349,
	"LAVENDER_BLUSH":                350,
	"LAVENDER_GRAY":                 351,
	"LAVENDER_INDIGO":               352,
	"LAVENDER_MAGENTA":              353,
	"LAVENDER_MIST":                 354,
	"LAVENDER_PINK":                 355,
	"LAVENDER_PURPLE":               356,
	"LAVENDER_ROSE":                 357,
	"LAWN_GREEN":                    358,
	"LEMON":                         359,
	"LEMON_YELLOW":                  360,
	"LEMON_CHIFFON":                 361,
	"LEMON_LIME":                    362,
	"LIGHT_CRIMSON":                 363,
	"LIGHT_THULIAN_PINK":            364,
	"LIGHT_APRICOT":                 365,
	"LIGHT_BLUE":                    366,
	"LIGHT_BROWN":                   367,
	"LIGHT_CARMINE_PINK":            368,
	"LIGHT_CORAL":                   369,
	"LIGHT_CORNFLOWER_BLUE":         370,
	"LIGHT_CYAN":                    371,
	"LIGHT_FUCHSIA_PINK":            372,
	"LIGHT_GOLDENROD_YELLOW":        373,
	"LIGHT_GRAY":                    374,
	"LIGHT_GREEN":                   375,
	"LIGHT_KHAKI":                   376,
	"LIGHT_PASTEL_PURPLE":           377,
	"LIGHT_PINK":                    378,
	"LIGHT_SALMON":                  379,
	"LIGHT_SALMON_PINK":             380,
	"LIGHT_SEA_GREEN":               381,
	"LIGHT_SKY_BLUE":                382,
	"LIGHT_SLATE_GRAY":              383,
	"LIGHT_TAUPE":                   384,
	"LIGHT_YELLOW":                  385,
	"LILAC":                         386,
	"LIME":                          387,
	"LIME_GREEN":                    388,
	"LINCOLN_GREEN":                 389,
	"LINEN":                         390,
	"LION":                          391,
	"LIVER":                         392,
	"LUST":                          393,
	"MSU_GREEN":                     394,
	"MACARONI_AND_CHEESE":           395,
	"MAGENTA":                       396,
	"MAGIC_MINT":                    397,
	"MAGNOLIA":                      398,
	"MAHOGANY":                      399,
	"MAIZE":                         400,
	"MAJORELLE_BLUE":                401,
	"MALACHITE":                     402,
	"MANATEE":                       403,
	"MANGO_TANGO":                   404,
	"MANTIS":                        405,
	"MAROON":                        406,
	"MAUVE":                         407,
	"MAUVE_TAUPE":                   408,
	"MAUVELOUS":                     409,
	"MAYA_BLUE":                     410,
	"MEAT_BROWN":                    411,
	"MEDIUM_PERSIAN_BLUE":           412,
	"MEDIUM_AQUAMARINE":             413,
	"MEDIUM_BLUE":                   414,
	"MEDIUM_CANDY_APPLE_RED":        415,
	"MEDIUM_CARMINE":                416,
	"MEDIUM_CHAMPAGNE":              417,
	"MEDIUM_ELECTRIC_BLUE":          418,
	"MEDIUM_JUNGLE_GREEN":           419,
	"MEDIUM_LAVENDER_MAGENTA":       420,
	"MEDIUM_ORCHID":                 421,
	"MEDIUM_PURPLE":                 422,
	"MEDIUM_RED_VIOLET":             423,
	"MEDIUM_SEA_GREEN":              424,
	"MEDIUM_SLATE_BLUE":             425,
	"MEDIUM_SPRING_BUD":             426,
	"MEDIUM_SPRING_GREEN":           427,
	"MEDIUM_TAUPE":                  428,
	"MEDIUM_TEAL_BLUE":              429,
	"MEDIUM_TURQUOISE":              430,
	"MEDIUM_VIOLET_RED":             431,
	"MELON":                         432,
	"MIDNIGHT_BLUE":                 433,
	"MIDNIGHT_GREEN":                434,
	"MIKADO_YELLOW":                 435,
	"MINT":                          436,
	"MINT_CREAM":                    437,
	"MINT_GREEN":                    438,
	"MISTY_ROSE":                    439,
	"MOCCASIN":                      440,
	"MODE_BEIGE":                    441,
	"MOONSTONE_BLUE":                442,
	"MORDANT_RED_19":                443,
	"MOSS_GREEN":                    444,
	"MOUNTAIN_MEADOW":               445,
	"MOUNTBATTEN_PINK":              446,
	"MULBERRY":                      447,
	"MUNSELL":                       448,
	"MUSTARD":                       449,
	"MYRTLE":                        450,
	"NADESHIKO_PINK":                451,
	"NAPIER_GREEN":                  452,
	"NAPLES_YELLOW":                 453,
	"NAVAJO_WHITE":                  454,
	"NAVY_BLUE":                     455,
	"NEON_CARROT":                   456,
	"NEON_FUCHSIA":                  457,
	"NEON_GREEN":                    458,
	"NON_PHOTO_BLUE":                459,
	"NORTH_TEXAS_GREEN":             460,
	"OCEAN_BOAT_BLUE":               461,
	"OCHRE":                         462,
	"OFFICE_GREEN":                  463,
	"OLD_GOLD":                      464,
	"OLD_LACE":                      465,
	"OLD_LAVENDER":                  466,
	"OLD_MAUVE":                     467,
	"OLD_ROSE":                      468,
	"OLIVE":                         469,
	"OLIVE_DRAB":                    470,
	"OLIVE_GREEN":                   471,
	"OLIVINE":                       472,
	"ONYX":                          473,
	"OPERA_MAUVE":                   474,
	"ORANGE":                        475,
	"ORANGE_YELLOW":                 476,
	"ORANGE_PEEL":                   477,
	"ORANGE_RED":                    478,
	"ORCHID":                        479,
	"OTTER_BROWN":                   480,
	"OUTER_SPACE":                   481,
	"OUTRAGEOUS_ORANGE":             482,
	"OXFORD_BLUE":                   483,
	"PACIFIC_BLUE":                  484,
	"PAKISTAN_GREEN":                485,
	"PALATINATE_BLUE":               486,
	"PALATINATE_PURPLE":             487,
	"PALE_AQUA":                     488,
	"PALE_BLUE":                     489,
	"PALE_BROWN":                    490,
	"PALE_CARMINE":                  491,
	"PALE_CERULEAN":                 492,
	"PALE_CHESTNUT":                 493,
	"PALE_COPPER":                   494,
	"PALE_CORNFLOWER_BLUE":          495,
	"PALE_GOLD":                     496,
	"PALE_GOLDENROD":                497,
	"PALE_GREEN":                    498,
	"PALE_LAVENDER":                 499,
	"PALE_MAGENTA":                  500,
	"PALE_PINK":                     501,
	"PALE_PLUM":                     502,
	"PALE_RED_VIOLET":               503,
	"PALE_ROBIN_EGG_BLUE":           504,
	"PALE_SILVER":                   505,
	"PALE_SPRING_BUD":               506,
	"PALE_TAUPE":                    507,
	"PALE_VIOLET_RED":               508,
	"PANSY_PURPLE":                  509,
	"PAPAYA_WHIP":                   510,
	"PARIS_GREEN":                   511,
	"PASTEL_BLUE":                   512,
	"PASTEL_BROWN":                  513,
	"PASTEL_GRAY":                   514,
	"PASTEL_GREEN":                  515,
	"PASTEL_MAGENTA":                516,
	"PASTEL_ORANGE":                 517,
	"PASTEL_PINK":                   518,
	"PASTEL_PURPLE":                 519,
	"PASTEL_RED":                    520,
	"PASTEL_VIOLET":                 521,
	"PASTEL_YELLOW":                 522,
	"PATRIARCH":                     523,
	"PAYNE_GREY":                    524,
	"PEACH":                         525,
	"PEACH_PUFF":                    526,
	"PEACH_YELLOW":                  527,
	"PEAR":                          528,
	"PEARL":                         529,
	"PEARL_AQUA":                    530,
	"PERIDOT":                       531,
	"PERIWINKLE":                    532,
	"PERSIAN_BLUE":                  533,
	"PERSIAN_INDIGO":                534,
	"PERSIAN_ORANGE":                535,
	"PERSIAN_PINK":                  536,
	"PERSIAN_PLUM":                  537,
	"PERSIAN_RED":                   538,
	"PERSIAN_ROSE":                  539,
	"PHLOX":                         540,
	"PHTHALO_BLUE":                  541,
	"PHTHALO_GREEN":                 542,
	"PIGGY_PINK":                    543,
	"PINE_GREEN":                    544,
	"PINK":                          545,
	"PINK_FLAMINGO":                 546,
	"PINK_SHERBET":                  547,
	"PINK_PEARL":                    548,
	"PISTACHIO":                     549,
	"PLATINUM":                      550,
	"PLUM":                          551,
	"PORTLAND_ORANGE":               552,
	"POWDER_BLUE":                   553,
	"PRINCETON_ORANGE":              554,
	"PRUSSIAN_BLUE":                 555,
	"PSYCHEDELIC_PURPLE":            556,
	"PUCE":                          557,
	"PUMPKIN":                       558,
	"PURPLE":                        559,
	"PURPLE_HEART":                  560,
	"PURPLE_MOUNTAINS_MAJESTY":      561,
	"PURPLE_MOUNTAIN_MAJESTY":       562,
	"PURPLE_PIZZAZZ":                563,
	"PURPLE_TAUPE":                  564,
	"RACKLEY":                       565,
	"RADICAL_RED":                   566,
	"RASPBERRY":                     567,
	"RASPBERRY_GLACE":               568,
	"RASPBERRY_PINK":                569,
	"RASPBERRY_ROSE":                570,
	"RAW_SIENNA":                    571,
	"RAZZLE_DAZZLE_ROSE":            572,
	"RAZZMATAZZ":                    573,
	"RED":                           574,
	"RED_ORANGE":                    575,
	"RED_BROWN":                     576,
	"RED_VIOLET":                    577,
	"RICH_BLACK":                    578,
	"RICH_CARMINE":                  579,
	"RICH_ELECTRIC_BLUE":            580,
	"RICH_LILAC":                    581,
	"RICH_MAROON":                   582,
	"RIFLE_GREEN":                   583,
	"ROBINS_EGG_BLUE":               584,
	"ROSE":                          585,
	"ROSE_BONBON":                   586,
	"ROSE_EBONY":                    587,
	"ROSE_GOLD":                     588,
	"ROSE_MADDER":                   589,
	"ROSE_PINK":                     590,
	"ROSE_QUARTZ":                   591,
	"ROSE_TAUPE":                    592,
	"ROSE_VALE":                     593,
	"ROSEWOOD":                      594,
	"ROSSO_CORSA":                   595,
	"ROSY_BROWN":                    596,
	"ROYAL_AZURE":                   597,
	"ROYAL_BLUE":                    598,
	"ROYAL_FUCHSIA":                 599,
	"ROYAL_PURPLE":                  600,
	"RUBY":                          601,
	"RUDDY":                         602,
	"RUDDY_BROWN":                   603,
	"RUDDY_PINK":                    604,
	"RUFOUS":                        605,
	"RUSSET":                        606,
	"RUST":                          607,
	"SACRAMENTO_STATE_GREEN":        608,
	"SADDLE_BROWN":                  609,
	"SAFETY_ORANGE":                 610,
	"SAFFRON":                       611,
	"SAINT_PATRICK_BLUE":            612,
	"SALMON":                        613,
	"SALMON_PINK":                   614,
	"SAND":                          615,
	"SAND_DUNE":                     616,
	"SANDSTORM":                     617,
	"SANDY_BROWN":                   618,
	"SANDY_TAUPE":                   619,
	"SAP_GREEN":                     620,
	"SAPPHIRE":                      621,
	"SATIN_SHEEN_GOLD":              622,
	"SCARLET":                       623,
	"SCHOOL_BUS_YELLOW":             624,
	"SCREAMIN_GREEN":                625,
	"SEA_BLUE":                      626,
	"SEA_GREEN":                     627,
	"SEAL_BROWN":                    628,
	"SEASHELL":                      629,
	"SELECTIVE_YELLOW":              630,
	"SEPIA":                         631,
	"SHADOW":                        632,
	"SHAMROCK":                      633,
	"SHAMROCK_GREEN":                634,
	"SHOCKING_PINK":                 635,
	"SIENNA":                        636,
	"SILVER":                        637,
	"SINOPIA":                       638,
	"SKOBELOFF":                     639,
	"SKY_BLUE":                      640,
	"SKY_MAGENTA":                   641,
	"SLATE_BLUE":                    642,
	"SLATE_GRAY":                    643,
	"SMALT":                         644,
	"SMOKEY_TOPAZ":                  645,
	"SMOKY_BLACK":                   646,
	"SNOW":                          647,
	"SPIRO_DISCO_BALL":              648,
	"SPRING_BUD":                    649,
	"SPRING_GREEN":                  650,
	"STEEL_BLUE":                    651,
	"STIL_DE_GRAIN_YELLOW":          652,
	"STIZZA":                        653,
	"STORMCLOUD":                    654,
	"STRAW":                         655,
	"SUNGLOW":                       656,
	"SUNSET":                        657,
	"SUNSET_ORANGE":                 658,
	"TAN":                           659,
	"TANGELO":                       660,
	"TANGERINE":                     661,
	"TANGERINE_YELLOW":              662,
	"TAUPE":                         663,
	"TAUPE_GRAY":                    664,
	"TAWNY":                         665,
	"TEA_GREEN":                     666,
	"TEA_ROSE":                      667,
	"TEAL":                          668,
	"TEAL_BLUE":                     669,
	"TEAL_GREEN":                    670,
	"TERRA_COTTA":                   671,
	"THISTLE":                       672,
	"THULIAN_PINK":                  673,
	"TICKLE_ME_PINK":                674,
	"TIFFANY_BLUE":                  675,
	"TIGER_EYE":                     676,
	"TIMBERWOLF":                    677,
	"TITANIUM_YELLOW":               678,
	"TOMATO":                        679,
	"TOOLBOX":                       680,
	"TOPAZ":                         681,
	"TRACTOR_RED":                   682,
	"TROLLEY_GREY":                  683,
	"TROPICAL_RAIN_FOREST":          684,
	"TRUE_BLUE":                     685,
	"TUFTS_BLUE":                    686,
	"TUMBLEWEED":                    687,
	"TURKISH_ROSE":                  688,
	"TURQUOISE":                     689,
	"TURQUOISE_BLUE":                690,
	"TURQUOISE_GREEN":               691,
	"TUSCAN_RED":                    692,
	"TWILIGHT_LAVENDER":             693,
	"TYRIAN_PURPLE":                 694,
	"UA_BLUE":                       695,
	"UA_RED":                        696,
	"UCLA_BLUE":                     697,
	"UCLA_GOLD":                     698,
	"UFO_GREEN":                     699,
	"UP_FOREST_GREEN":               700,
	"UP_MAROON":                     701,
	"USC_CARDINAL":                  702,
	"USC_GOLD":                      703,
	"UBE":                           704,
	"ULTRA_PINK":                    705,
	"ULTRAMARINE":                   706,
	"ULTRAMARINE_BLUE":              707,
	"UMBER":                         708,
	"UNITED_NATIONS_BLUE":           709,
	"UNIVERSITY_OF_CALIFORNIA_GOLD": 710,
	"UNMELLOW_YELLOW":               711,
	"UPSDELL_RED":                   712,
	"UROBILIN":                      713,
	"UTAH_CRIMSON":                  714,
	"VANILLA":                       715,
	"VEGAS_GOLD":                    716,
	"VENETIAN_RED":                  717,
	"VERDIGRIS":                     718,
	"VERMILION":                     719,
	"VERONICA":                      720,
	"VIOLET":                        721,
	"VIOLET_BLUE":                   722,
	"VIOLET_RED":                    723,
	"VIRIDIAN":                      724,
	"VIVID_AUBURN":                  725,
	"VIVID_BURGUNDY":                726,
	"VIVID_CERISE":                  727,
	"VIVID_TANGERINE":               728,
	"VIVID_VIOLET":                  729,
	"WARM_BLACK":                    730,
	"WATERSPOUT":                    731,
	"WENGE":                         732,
	"WHEAT":                         733,
	"WHITE":                         734,
	"WHITE_SMOKE":                   735,
	"WILD_STRAWBERRY":               736,
	"WILD_WATERMELON":               737,
	"WILD_BLUE_YONDER":              738,
	"WINE":                          739,
	"WISTERIA":                      740,
	"XANADU":                        741,
	"YALE_BLUE":                     742,
	"YELLOW":                        743,
	"YELLOW_ORANGE":                 744,
	"YELLOW_GREEN":                  745,
	"ZAFFRE":                        746,
	"ZINNWALDITE_BROWN":             747,
	"METALLIC":                      748,
	"DARK_PURPLE":                   749,
	"PEWTER":                        750,
	"BURN":                          751,
	"VINOUS":                        752,
	"TWO_TONE":                      753,
	"LIGHT_RED":                     754,
	"LIGHT_ORANGE":                  755,
}

// ColorAliases is a set of regexp that defines
// everything that can be considered the same as enum
// name.
var ColorAliases = []state.Alias{
	{Alias: "^(other|unknown|unknownnotokforinv)$", Value: 1},
	{Alias: "^airforceblue$", Value: 2},
	{Alias: "^aliceblue$", Value: 3},
	{Alias: "^alizarincrimson$", Value: 4},
	{Alias: "^almond$", Value: 5},
	{Alias: "^amaranth$", Value: 6},
	{Alias: "^amber$", Value: 7},
	{Alias: "^americanrose$", Value: 8},
	{Alias: "^amethyst$", Value: 9},
	{Alias: "^androidgreen$", Value: 10},
	{Alias: "^antiflashwhite$", Value: 11},
	{Alias: "^antiquebrass$", Value: 12},
	{Alias: "^antiquefuchsia$", Value: 13},
	{Alias: "^antiquewhite$", Value: 14},
	{Alias: "^ao$", Value: 15},
	{Alias: "^applegreen$", Value: 16},
	{Alias: "^apricot$", Value: 17},
	{Alias: "^aqua$", Value: 18},
	{Alias: "^aquamarine$", Value: 19},
	{Alias: "^armygreen$", Value: 20},
	{Alias: "^arylideyellow$", Value: 21},
	{Alias: "^ashgrey$", Value: 22},
	{Alias: "^asparagus$", Value: 23},
	{Alias: "^atomictangerine$", Value: 24},
	{Alias: "^auburn$", Value: 25},
	{Alias: "^aureolin$", Value: 26},
	{Alias: "^aurometalsaurus$", Value: 27},
	{Alias: "^awesome$", Value: 28},
	{Alias: "^azure$", Value: 29},
	{Alias: "^azuremistweb$", Value: 30},
	{Alias: "^babyblue$", Value: 31},
	{Alias: "^babyblueeyes$", Value: 32},
	{Alias: "^babypink$", Value: 33},
	{Alias: "^ballblue$", Value: 34},
	{Alias: "^bananamania$", Value: 35},
	{Alias: "^bananayellow$", Value: 36},
	{Alias: "^battleshipgrey$", Value: 37},
	{Alias: "^bazaar$", Value: 38},
	{Alias: "^beaublue$", Value: 39},
	{Alias: "^beaver$", Value: 40},
	{Alias: "^(beige|lightbeige)$", Value: 41},
	{Alias: "^bisque$", Value: 42},
	{Alias: "^bistre$", Value: 43},
	{Alias: "^bittersweet$", Value: 44},
	{Alias: "^black$", Value: 45},
	{Alias: "^blanchedalmond$", Value: 46},
	{Alias: "^bleudefrance$", Value: 47},
	{Alias: "^blizzardblue$", Value: 48},
	{Alias: "^blond$", Value: 49},
	{Alias: "^blue$", Value: 50},
	{Alias: "^bluebell$", Value: 51},
	{Alias: "^bluegray$", Value: 52},
	{Alias: "^bluegreen$", Value: 53},
	{Alias: "^bluepurple$", Value: 54},
	{Alias: "^blueviolet$", Value: 55},
	{Alias: "^blush$", Value: 56},
	{Alias: "^bole$", Value: 57},
	{Alias: "^bondiblue$", Value: 58},
	{Alias: "^bone$", Value: 59},
	{Alias: "^bostonuniversityred$", Value: 60},
	{Alias: "^bottlegreen$", Value: 61},
	{Alias: "^boysenberry$", Value: 62},
	{Alias: "^brandeisblue$", Value: 63},
	{Alias: "^brass$", Value: 64},
	{Alias: "^brickred$", Value: 65},
	{Alias: "^brightcerulean$", Value: 66},
	{Alias: "^brightgreen$", Value: 67},
	{Alias: "^brightlavender$", Value: 68},
	{Alias: "^brightmaroon$", Value: 69},
	{Alias: "^brightpink$", Value: 70},
	{Alias: "^brightturquoise$", Value: 71},
	{Alias: "^brightube$", Value: 72},
	{Alias: "^brilliantlavender$", Value: 73},
	{Alias: "^brilliantrose$", Value: 74},
	{Alias: "^brinkpink$", Value: 75},
	{Alias: "^britishracinggreen$", Value: 76},
	{Alias: "^bronze$", Value: 77},
	{Alias: "^brown$", Value: 78},
	{Alias: "^bubblegum$", Value: 79},
	{Alias: "^bubbles$", Value: 80},
	{Alias: "^buff$", Value: 81},
	{Alias: "^bulgarianrose$", Value: 82},
	{Alias: "^burgundy$", Value: 83},
	{Alias: "^burlywood$", Value: 84},
	{Alias: "^burntorange$", Value: 85},
	{Alias: "^burntsienna$", Value: 86},
	{Alias: "^burntumber$", Value: 87},
	{Alias: "^byzantine$", Value: 88},
	{Alias: "^byzantium$", Value: 89},
	{Alias: "^cgblue$", Value: 90},
	{Alias: "^cgred$", Value: 91},
	{Alias: "^cadet$", Value: 92},
	{Alias: "^cadetblue$", Value: 93},
	{Alias: "^cadetgrey$", Value: 94},
	{Alias: "^cadmiumgreen$", Value: 95},
	{Alias: "^cadmiumorange$", Value: 96},
	{Alias: "^cadmiumred$", Value: 97},
	{Alias: "^cadmiumyellow$", Value: 98},
	{Alias: "^cafeaulait$", Value: 99},
	{Alias: "^cafenoir$", Value: 100},
	{Alias: "^calpolypomonagreen$", Value: 101},
	{Alias: "^cambridgeblue$", Value: 102},
	{Alias: "^camel$", Value: 103},
	{Alias: "^camouflagegreen$", Value: 104},
	{Alias: "^canary$", Value: 105},
	{Alias: "^canaryyellow$", Value: 106},
	{Alias: "^candyapplered$", Value: 107},
	{Alias: "^candypink$", Value: 108},
	{Alias: "^capri$", Value: 109},
	{Alias: "^caputmortuum$", Value: 110},
	{Alias: "^cardinal$", Value: 111},
	{Alias: "^caribbeangreen$", Value: 112},
	{Alias: "^carmine$", Value: 113},
	{Alias: "^carminepink$", Value: 114},
	{Alias: "^carminered$", Value: 115},
	{Alias: "^carnationpink$", Value: 116},
	{Alias: "^(carnelian|carnelianred)$", Value: 117},
	{Alias: "^carolinablue$", Value: 118},
	{Alias: "^carrotorange$", Value: 119},
	{Alias: "^celadon$", Value: 120},
	{Alias: "^celeste$", Value: 121},
	{Alias: "^celestialblue$", Value: 122},
	{Alias: "^cerise$", Value: 123},
	{Alias: "^cerisepink$", Value: 124},
	{Alias: "^cerulean$", Value: 125},
	{Alias: "^ceruleanblue$", Value: 126},
	{Alias: "^chamoisee$", Value: 127},
	{Alias: "^champagne$", Value: 128},
	{Alias: "^charcoal$", Value: 129},
	{Alias: "^chartreuse$", Value: 130},
	{Alias: "^cherry$", Value: 131},
	{Alias: "^cherryblossompink$", Value: 132},
	{Alias: "^chestnut$", Value: 133},
	{Alias: "^chocolate$", Value: 134},
	{Alias: "^chromeyellow$", Value: 135},
	{Alias: "^cinereous$", Value: 136},
	{Alias: "^cinnabar$", Value: 137},
	{Alias: "^cinnamon$", Value: 138},
	{Alias: "^citrine$", Value: 139},
	{Alias: "^classicrose$", Value: 140},
	{Alias: "^cobalt$", Value: 141},
	{Alias: "^cocoabrown$", Value: 142},
	{Alias: "^coffee$", Value: 143},
	{Alias: "^columbiablue$", Value: 144},
	{Alias: "^coolblack$", Value: 145},
	{Alias: "^coolgrey$", Value: 146},
	{Alias: "^copper$", Value: 147},
	{Alias: "^copperrose$", Value: 148},
	{Alias: "^coquelicot$", Value: 149},
	{Alias: "^coral$", Value: 150},
	{Alias: "^coralpink$", Value: 151},
	{Alias: "^coralred$", Value: 152},
	{Alias: "^cordovan$", Value: 153},
	{Alias: "^corn$", Value: 154},
	{Alias: "^cornellred$", Value: 155},
	{Alias: "^cornflower$", Value: 156},
	{Alias: "^cornflowerblue$", Value: 157},
	{Alias: "^cornsilk$", Value: 158},
	{Alias: "^cosmiclatte$", Value: 159},
	{Alias: "^cottoncandy$", Value: 160},
	{Alias: "^cream$", Value: 161},
	{Alias: "^crimson$", Value: 162},
	{Alias: "^crimsonred$", Value: 163},
	{Alias: "^crimsonglory$", Value: 164},
	{Alias: "^cyan$", Value: 165},
	{Alias: "^daffodil$", Value: 166},
	{Alias: "^dandelion$", Value: 167},
	{Alias: "^darkblue$", Value: 168},
	{Alias: "^darkbrown$", Value: 169},
	{Alias: "^darkbyzantium$", Value: 170},
	{Alias: "^darkcandyapplered$", Value: 171},
	{Alias: "^darkcerulean$", Value: 172},
	{Alias: "^darkchestnut$", Value: 173},
	{Alias: "^darkcoral$", Value: 174},
	{Alias: "^darkcyan$", Value: 175},
	{Alias: "^darkelectricblue$", Value: 176},
	{Alias: "^(darkgoldenrod|darkyellow)$", Value: 177},
	{Alias: "^(darkgray|darkgrey)$", Value: 178},
	{Alias: "^darkgreen$", Value: 179},
	{Alias: "^darkjunglegreen$", Value: 180},
	{Alias: "^darkkhaki$", Value: 181},
	{Alias: "^darklava$", Value: 182},
	{Alias: "^darklavender$", Value: 183},
	{Alias: "^darkmagenta$", Value: 184},
	{Alias: "^darkmidnightblue$", Value: 185},
	{Alias: "^darkolivegreen$", Value: 186},
	{Alias: "^darkorange$", Value: 187},
	{Alias: "^darkorchid$", Value: 188},
	{Alias: "^darkpastelblue$", Value: 189},
	{Alias: "^darkpastelgreen$", Value: 190},
	{Alias: "^darkpastelpurple$", Value: 191},
	{Alias: "^darkpastelred$", Value: 192},
	{Alias: "^darkpink$", Value: 193},
	{Alias: "^darkpowderblue$", Value: 194},
	{Alias: "^darkraspberry$", Value: 195},
	{Alias: "^darkred$", Value: 196},
	{Alias: "^darksalmon$", Value: 197},
	{Alias: "^darkscarlet$", Value: 198},
	{Alias: "^darkseagreen$", Value: 199},
	{Alias: "^darksienna$", Value: 200},
	{Alias: "^darkslateblue$", Value: 201},
	{Alias: "^darkslategray$", Value: 202},
	{Alias: "^darkspringgreen$", Value: 203},
	{Alias: "^darktan$", Value: 204},
	{Alias: "^darktangerine$", Value: 205},
	{Alias: "^darktaupe$", Value: 206},
	{Alias: "^darkterracotta$", Value: 207},
	{Alias: "^darkturquoise$", Value: 208},
	{Alias: "^darkviolet$", Value: 209},
	{Alias: "^dartmouthgreen$", Value: 210},
	{Alias: "^davygrey$", Value: 211},
	{Alias: "^debianred$", Value: 212},
	{Alias: "^deepcarmine$", Value: 213},
	{Alias: "^deepcarminepink$", Value: 214},
	{Alias: "^deepcarrotorange$", Value: 215},
	{Alias: "^deepcerise$", Value: 216},
	{Alias: "^deepchampagne$", Value: 217},
	{Alias: "^deepchestnut$", Value: 218},
	{Alias: "^deepcoffee$", Value: 219},
	{Alias: "^deepfuchsia$", Value: 220},
	{Alias: "^deepjunglegreen$", Value: 221},
	{Alias: "^deeplilac$", Value: 222},
	{Alias: "^deepmagenta$", Value: 223},
	{Alias: "^deeppeach$", Value: 224},
	{Alias: "^deeppink$", Value: 225},
	{Alias: "^deepsaffron$", Value: 226},
	{Alias: "^deepskyblue$", Value: 227},
	{Alias: "^denim$", Value: 228},
	{Alias: "^desert$", Value: 229},
	{Alias: "^desertsand$", Value: 230},
	{Alias: "^dimgray$", Value: 231},
	{Alias: "^dodgerblue$", Value: 232},
	{Alias: "^dogwoodrose$", Value: 233},
	{Alias: "^dollarbill$", Value: 234},
	{Alias: "^drab$", Value: 235},
	{Alias: "^dukeblue$", Value: 236},
	{Alias: "^earthyellow$", Value: 237},
	{Alias: "^ecru$", Value: 238},
	{Alias: "^eggplant$", Value: 239},
	{Alias: "^eggshell$", Value: 240},
	{Alias: "^egyptianblue$", Value: 241},
	{Alias: "^electricblue$", Value: 242},
	{Alias: "^electriccrimson$", Value: 243},
	{Alias: "^electriccyan$", Value: 244},
	{Alias: "^electricgreen$", Value: 245},
	{Alias: "^electricindigo$", Value: 246},
	{Alias: "^electriclavender$", Value: 247},
	{Alias: "^electriclime$", Value: 248},
	{Alias: "^electricpurple$", Value: 249},
	{Alias: "^electricultramarine$", Value: 250},
	{Alias: "^electricviolet$", Value: 251},
	{Alias: "^electricyellow$", Value: 252},
	{Alias: "^emerald$", Value: 253},
	{Alias: "^etonblue$", Value: 254},
	{Alias: "^fallow$", Value: 255},
	{Alias: "^falured$", Value: 256},
	{Alias: "^famous$", Value: 257},
	{Alias: "^fandango$", Value: 258},
	{Alias: "^fashionfuchsia$", Value: 259},
	{Alias: "^fawn$", Value: 260},
	{Alias: "^feldgrau$", Value: 261},
	{Alias: "^fern$", Value: 262},
	{Alias: "^ferngreen$", Value: 263},
	{Alias: "^ferrarired$", Value: 264},
	{Alias: "^fielddrab$", Value: 265},
	{Alias: "^fireenginered$", Value: 266},
	{Alias: "^firebrick$", Value: 267},
	{Alias: "^flame$", Value: 268},
	{Alias: "^flamingopink$", Value: 269},
	{Alias: "^flavescent$", Value: 270},
	{Alias: "^flax$", Value: 271},
	{Alias: "^floralwhite$", Value: 272},
	{Alias: "^fluorescentorange$", Value: 273},
	{Alias: "^fluorescentpink$", Value: 274},
	{Alias: "^fluorescentyellow$", Value: 275},
	{Alias: "^folly$", Value: 276},
	{Alias: "^forestgreen$", Value: 277},
	{Alias: "^frenchbeige$", Value: 278},
	{Alias: "^frenchblue$", Value: 279},
	{Alias: "^frenchlilac$", Value: 280},
	{Alias: "^frenchrose$", Value: 281},
	{Alias: "^fuchsia$", Value: 282},
	{Alias: "^fuchsiapink$", Value: 283},
	{Alias: "^fulvous$", Value: 284},
	{Alias: "^fuzzywuzzy$", Value: 285},
	{Alias: "^gainsboro$", Value: 286},
	{Alias: "^gamboge$", Value: 287},
	{Alias: "^ghostwhite$", Value: 288},
	{Alias: "^ginger$", Value: 289},
	{Alias: "^glaucous$", Value: 290},
	{Alias: "^glitter$", Value: 291},
	{Alias: "^(gold|golden)$", Value: 292},
	{Alias: "^goldenbrown$", Value: 293},
	{Alias: "^goldenpoppy$", Value: 294},
	{Alias: "^goldenyellow$", Value: 295},
	{Alias: "^goldenrod$", Value: 296},
	{Alias: "^grannysmithapple$", Value: 297},
	{Alias: "^(gray|grey)$", Value: 298},
	{Alias: "^grayasparagus$", Value: 299},
	{Alias: "^green$", Value: 300},
	{Alias: "^greenblue$", Value: 301},
	{Alias: "^greenyellow$", Value: 302},
	{Alias: "^grullo$", Value: 303},
	{Alias: "^guppiegreen$", Value: 304},
	{Alias: "^halayaube$", Value: 305},
	{Alias: "^hanblue$", Value: 306},
	{Alias: "^hanpurple$", Value: 307},
	{Alias: "^hansayellow$", Value: 308},
	{Alias: "^harlequin$", Value: 309},
	{Alias: "^harvardcrimson$", Value: 310},
	{Alias: "^harvestgold$", Value: 311},
	{Alias: "^heartgold$", Value: 312},
	{Alias: "^heliotrope$", Value: 313},
	{Alias: "^hollywoodcerise$", Value: 314},
	{Alias: "^honeydew$", Value: 315},
	{Alias: "^hookergreen$", Value: 316},
	{Alias: "^hotmagenta$", Value: 317},
	{Alias: "^hotpink$", Value: 318},
	{Alias: "^huntergreen$", Value: 319},
	{Alias: "^icterine$", Value: 320},
	{Alias: "^inchworm$", Value: 321},
	{Alias: "^indiagreen$", Value: 322},
	{Alias: "^indianred$", Value: 323},
	{Alias: "^indianyellow$", Value: 324},
	{Alias: "^indigo$", Value: 325},
	{Alias: "^internationalkleinblue$", Value: 326},
	{Alias: "^internationalorange$", Value: 327},
	{Alias: "^iris$", Value: 328},
	{Alias: "^isabelline$", Value: 329},
	{Alias: "^islamicgreen$", Value: 330},
	{Alias: "^ivory$", Value: 331},
	{Alias: "^jade$", Value: 332},
	{Alias: "^jasmine$", Value: 333},
	{Alias: "^jasper$", Value: 334},
	{Alias: "^jazzberryjam$", Value: 335},
	{Alias: "^jonquil$", Value: 336},
	{Alias: "^junebud$", Value: 337},
	{Alias: "^junglegreen$", Value: 338},
	{Alias: "^kucrimson$", Value: 339},
	{Alias: "^kellygreen$", Value: 340},
	{Alias: "^khaki$", Value: 341},
	{Alias: "^lasallegreen$", Value: 342},
	{Alias: "^languidlavender$", Value: 343},
	{Alias: "^lapislazuli$", Value: 344},
	{Alias: "^laserlemon$", Value: 345},
	{Alias: "^laurelgreen$", Value: 346},
	{Alias: "^lava$", Value: 347},
	{Alias: "^lavender$", Value: 348},
	{Alias: "^lavenderblue$", Value: 349},
	{Alias: "^lavenderblush$", Value: 350},
	{Alias: "^lavendergray$", Value: 351},
	{Alias: "^lavenderindigo$", Value: 352},
	{Alias: "^lavendermagenta$", Value: 353},
	{Alias: "^lavendermist$", Value: 354},
	{Alias: "^lavenderpink$", Value: 355},
	{Alias: "^lavenderpurple$", Value: 356},
	{Alias: "^lavenderrose$", Value: 357},
	{Alias: "^lawngreen$", Value: 358},
	{Alias: "^lemon$", Value: 359},
	{Alias: "^lemonyellow$", Value: 360},
	{Alias: "^lemonchiffon$", Value: 361},
	{Alias: "^lemonlime$", Value: 362},
	{Alias: "^lightcrimson$", Value: 363},
	{Alias: "^lightthulianpink$", Value: 364},
	{Alias: "^lightapricot$", Value: 365},
	{Alias: "^lightblue$", Value: 366},
	{Alias: "^lightbrown$", Value: 367},
	{Alias: "^lightcarminepink$", Value: 368},
	{Alias: "^lightcoral$", Value: 369},
	{Alias: "^lightcornflowerblue$", Value: 370},
	{Alias: "^lightcyan$", Value: 371},
	{Alias: "^lightfuchsiapink$", Value: 372},
	{Alias: "^lightgoldenrodyellow$", Value: 373},
	{Alias: "^(lightgray|lightgrey)$", Value: 374},
	{Alias: "^lightgreen$", Value: 375},
	{Alias: "^lightkhaki$", Value: 376},
	{Alias: "^lightpastelpurple$", Value: 377},
	{Alias: "^lightpink$", Value: 378},
	{Alias: "^lightsalmon$", Value: 379},
	{Alias: "^lightsalmonpink$", Value: 380},
	{Alias: "^lightseagreen$", Value: 381},
	{Alias: "^lightskyblue$", Value: 382},
	{Alias: "^lightslategray$", Value: 383},
	{Alias: "^lighttaupe$", Value: 384},
	{Alias: "^lightyellow$", Value: 385},
	{Alias: "^lilac$", Value: 386},
	{Alias: "^lime$", Value: 387},
	{Alias: "^limegreen$", Value: 388},
	{Alias: "^lincolngreen$", Value: 389},
	{Alias: "^linen$", Value: 390},
	{Alias: "^lion$", Value: 391},
	{Alias: "^liver$", Value: 392},
	{Alias: "^lust$", Value: 393},
	{Alias: "^msugreen$", Value: 394},
	{Alias: "^macaroniandcheese$", Value: 395},
	{Alias: "^magenta$", Value: 396},
	{Alias: "^magicmint$", Value: 397},
	{Alias: "^magnolia$", Value: 398},
	{Alias: "^mahogany$", Value: 399},
	{Alias: "^maize$", Value: 400},
	{Alias: "^majorelleblue$", Value: 401},
	{Alias: "^malachite$", Value: 402},
	{Alias: "^manatee$", Value: 403},
	{Alias: "^mangotango$", Value: 404},
	{Alias: "^mantis$", Value: 405},
	{Alias: "^maroon$", Value: 406},
	{Alias: "^mauve$", Value: 407},
	{Alias: "^mauvetaupe$", Value: 408},
	{Alias: "^mauvelous$", Value: 409},
	{Alias: "^mayablue$", Value: 410},
	{Alias: "^meatbrown$", Value: 411},
	{Alias: "^mediumpersianblue$", Value: 412},
	{Alias: "^mediumaquamarine$", Value: 413},
	{Alias: "^mediumblue$", Value: 414},
	{Alias: "^mediumcandyapplered$", Value: 415},
	{Alias: "^mediumcarmine$", Value: 416},
	{Alias: "^mediumchampagne$", Value: 417},
	{Alias: "^mediumelectricblue$", Value: 418},
	{Alias: "^mediumjunglegreen$", Value: 419},
	{Alias: "^mediumlavendermagenta$", Value: 420},
	{Alias: "^mediumorchid$", Value: 421},
	{Alias: "^mediumpurple$", Value: 422},
	{Alias: "^mediumredviolet$", Value: 423},
	{Alias: "^mediumseagreen$", Value: 424},
	{Alias: "^mediumslateblue$", Value: 425},
	{Alias: "^mediumspringbud$", Value: 426},
	{Alias: "^mediumspringgreen$", Value: 427},
	{Alias: "^mediumtaupe$", Value: 428},
	{Alias: "^mediumtealblue$", Value: 429},
	{Alias: "^mediumturquoise$", Value: 430},
	{Alias: "^mediumvioletred$", Value: 431},
	{Alias: "^melon$", Value: 432},
	{Alias: "^midnightblue$", Value: 433},
	{Alias: "^midnightgreen$", Value: 434},
	{Alias: "^mikadoyellow$", Value: 435},
	{Alias: "^mint$", Value: 436},
	{Alias: "^mintcream$", Value: 437},
	{Alias: "^mintgreen$", Value: 438},
	{Alias: "^mistyrose$", Value: 439},
	{Alias: "^moccasin$", Value: 440},
	{Alias: "^(modebeige|darkbeige)$", Value: 441},
	{Alias: "^moonstoneblue$", Value: 442},
	{Alias: "^mordantred19$", Value: 443},
	{Alias: "^mossgreen$", Value: 444},
	{Alias: "^mountainmeadow$", Value: 445},
	{Alias: "^mountbattenpink$", Value: 446},
	{Alias: "^mulberry$", Value: 447},
	{Alias: "^munsell$", Value: 448},
	{Alias: "^mustard$", Value: 449},
	{Alias: "^myrtle$", Value: 450},
	{Alias: "^nadeshikopink$", Value: 451},
	{Alias: "^napiergreen$", Value: 452},
	{Alias: "^naplesyellow$", Value: 453},
	{Alias: "^navajowhite$", Value: 454},
	{Alias: "^navyblue|navy$", Value: 455},
	{Alias: "^neoncarrot$", Value: 456},
	{Alias: "^neonfuchsia$", Value: 457},
	{Alias: "^neongreen$", Value: 458},
	{Alias: "^nonphotoblue$", Value: 459},
	{Alias: "^northtexasgreen$", Value: 460},
	{Alias: "^oceanboatblue$", Value: 461},
	{Alias: "^ochre$", Value: 462},
	{Alias: "^officegreen$", Value: 463},
	{Alias: "^oldgold$", Value: 464},
	{Alias: "^oldlace$", Value: 465},
	{Alias: "^oldlavender$", Value: 466},
	{Alias: "^oldmauve$", Value: 467},
	{Alias: "^oldrose$", Value: 468},
	{Alias: "^olive$", Value: 469},
	{Alias: "^olivedrab$", Value: 470},
	{Alias: "^olivegreen$", Value: 471},
	{Alias: "^olivine$", Value: 472},
	{Alias: "^onyx$", Value: 473},
	{Alias: "^operamauve$", Value: 474},
	{Alias: "^orange$", Value: 475},
	{Alias: "^orangeyellow$", Value: 476},
	{Alias: "^orangepeel$", Value: 477},
	{Alias: "^orangered$", Value: 478},
	{Alias: "^orchid$", Value: 479},
	{Alias: "^otterbrown$", Value: 480},
	{Alias: "^outerspace$", Value: 481},
	{Alias: "^outrageousorange$", Value: 482},
	{Alias: "^oxfordblue$", Value: 483},
	{Alias: "^pacificblue$", Value: 484},
	{Alias: "^pakistangreen$", Value: 485},
	{Alias: "^palatinateblue$", Value: 486},
	{Alias: "^palatinatepurple$", Value: 487},
	{Alias: "^paleaqua$", Value: 488},
	{Alias: "^paleblue$", Value: 489},
	{Alias: "^palebrown$", Value: 490},
	{Alias: "^palecarmine$", Value: 491},
	{Alias: "^palecerulean$", Value: 492},
	{Alias: "^palechestnut$", Value: 493},
	{Alias: "^palecopper$", Value: 494},
	{Alias: "^palecornflowerblue$", Value: 495},
	{Alias: "^palegold$", Value: 496},
	{Alias: "^palegoldenrod$", Value: 497},
	{Alias: "^palegreen$", Value: 498},
	{Alias: "^palelavender$", Value: 499},
	{Alias: "^palemagenta$", Value: 500},
	{Alias: "^palepink$", Value: 501},
	{Alias: "^paleplum$", Value: 502},
	{Alias: "^paleredviolet$", Value: 503},
	{Alias: "^palerobineggblue$", Value: 504},
	{Alias: "^palesilver$", Value: 505},
	{Alias: "^palespringbud$", Value: 506},
	{Alias: "^paletaupe$", Value: 507},
	{Alias: "^palevioletred$", Value: 508},
	{Alias: "^pansypurple$", Value: 509},
	{Alias: "^papayawhip$", Value: 510},
	{Alias: "^parisgreen$", Value: 511},
	{Alias: "^pastelblue$", Value: 512},
	{Alias: "^pastelbrown$", Value: 513},
	{Alias: "^pastelgray$", Value: 514},
	{Alias: "^pastelgreen$", Value: 515},
	{Alias: "^pastelmagenta$", Value: 516},
	{Alias: "^pastelorange$", Value: 517},
	{Alias: "^pastelpink$", Value: 518},
	{Alias: "^pastelpurple$", Value: 519},
	{Alias: "^pastelred$", Value: 520},
	{Alias: "^pastelviolet$", Value: 521},
	{Alias: "^pastelyellow$", Value: 522},
	{Alias: "^patriarch$", Value: 523},
	{Alias: "^paynegrey$", Value: 524},
	{Alias: "^peach$", Value: 525},
	{Alias: "^peachpuff$", Value: 526},
	{Alias: "^peachyellow$", Value: 527},
	{Alias: "^pear$", Value: 528},
	{Alias: "^pearl$", Value: 529},
	{Alias: "^pearlaqua$", Value: 530},
	{Alias: "^peridot$", Value: 531},
	{Alias: "^periwinkle$", Value: 532},
	{Alias: "^persianblue$", Value: 533},
	{Alias: "^persianindigo$", Value: 534},
	{Alias: "^persianorange$", Value: 535},
	{Alias: "^persianpink$", Value: 536},
	{Alias: "^persianplum$", Value: 537},
	{Alias: "^persianred$", Value: 538},
	{Alias: "^persianrose$", Value: 539},
	{Alias: "^phlox$", Value: 540},
	{Alias: "^phthaloblue$", Value: 541},
	{Alias: "^phthalogreen$", Value: 542},
	{Alias: "^piggypink$", Value: 543},
	{Alias: "^pinegreen$", Value: 544},
	{Alias: "^pink$", Value: 545},
	{Alias: "^pinkflamingo$", Value: 546},
	{Alias: "^pinksherbet$", Value: 547},
	{Alias: "^pinkpearl$", Value: 548},
	{Alias: "^pistachio$", Value: 549},
	{Alias: "^platinum$", Value: 550},
	{Alias: "^plum$", Value: 551},
	{Alias: "^portlandorange$", Value: 552},
	{Alias: "^powderblue$", Value: 553},
	{Alias: "^princetonorange$", Value: 554},
	{Alias: "^prussianblue$", Value: 555},
	{Alias: "^psychedelicpurple$", Value: 556},
	{Alias: "^puce$", Value: 557},
	{Alias: "^pumpkin$", Value: 558},
	{Alias: "^purple$", Value: 559},
	{Alias: "^purpleheart$", Value: 560},
	{Alias: "^purplemountainsmajesty$", Value: 561},
	{Alias: "^purplemountainmajesty$", Value: 562},
	{Alias: "^purplepizzazz$", Value: 563},
	{Alias: "^purpletaupe$", Value: 564},
	{Alias: "^rackley$", Value: 565},
	{Alias: "^radicalred$", Value: 566},
	{Alias: "^raspberry$", Value: 567},
	{Alias: "^raspberryglace$", Value: 568},
	{Alias: "^raspberrypink$", Value: 569},
	{Alias: "^raspberryrose$", Value: 570},
	{Alias: "^rawsienna$", Value: 571},
	{Alias: "^razzledazzlerose$", Value: 572},
	{Alias: "^razzmatazz$", Value: 573},
	{Alias: "^red$", Value: 574},
	{Alias: "^redorange$", Value: 575},
	{Alias: "^redbrown$", Value: 576},
	{Alias: "^redviolet$", Value: 577},
	{Alias: "^richblack$", Value: 578},
	{Alias: "^richcarmine$", Value: 579},
	{Alias: "^richelectricblue$", Value: 580},
	{Alias: "^richlilac$", Value: 581},
	{Alias: "^richmaroon$", Value: 582},
	{Alias: "^riflegreen$", Value: 583},
	{Alias: "^robinseggblue$", Value: 584},
	{Alias: "^rose$", Value: 585},
	{Alias: "^rosebonbon$", Value: 586},
	{Alias: "^roseebony$", Value: 587},
	{Alias: "^rosegold$", Value: 588},
	{Alias: "^rosemadder$", Value: 589},
	{Alias: "^rosepink$", Value: 590},
	{Alias: "^rosequartz$", Value: 591},
	{Alias: "^rosetaupe$", Value: 592},
	{Alias: "^rosevale$", Value: 593},
	{Alias: "^rosewood$", Value: 594},
	{Alias: "^rossocorsa$", Value: 595},
	{Alias: "^rosybrown$", Value: 596},
	{Alias: "^royalazure$", Value: 597},
	{Alias: "^royalblue$", Value: 598},
	{Alias: "^royalfuchsia$", Value: 599},
	{Alias: "^royalpurple$", Value: 600},
	{Alias: "^ruby$", Value: 601},
	{Alias: "^ruddy$", Value: 602},
	{Alias: "^ruddybrown$", Value: 603},
	{Alias: "^ruddypink$", Value: 604},
	{Alias: "^rufous$", Value: 605},
	{Alias: "^russet$", Value: 606},
	{Alias: "^rust$", Value: 607},
	{Alias: "^sacramentostategreen$", Value: 608},
	{Alias: "^saddlebrown$", Value: 609},
	{Alias: "^safetyorange$", Value: 610},
	{Alias: "^saffron$", Value: 611},
	{Alias: "^saintpatrickblue$", Value: 612},
	{Alias: "^salmon$", Value: 613},
	{Alias: "^salmonpink$", Value: 614},
	{Alias: "^sand$", Value: 615},
	{Alias: "^sanddune$", Value: 616},
	{Alias: "^sandstorm$", Value: 617},
	{Alias: "^sandybrown$", Value: 618},
	{Alias: "^sandytaupe$", Value: 619},
	{Alias: "^sapgreen$", Value: 620},
	{Alias: "^sapphire$", Value: 621},
	{Alias: "^satinsheengold$", Value: 622},
	{Alias: "^scarlet$", Value: 623},
	{Alias: "^schoolbusyellow$", Value: 624},
	{Alias: "^screamingreen$", Value: 625},
	{Alias: "^seablue$", Value: 626},
	{Alias: "^seagreen$", Value: 627},
	{Alias: "^sealbrown$", Value: 628},
	{Alias: "^seashell$", Value: 629},
	{Alias: "^selectiveyellow$", Value: 630},
	{Alias: "^sepia$", Value: 631},
	{Alias: "^shadow$", Value: 632},
	{Alias: "^shamrock$", Value: 633},
	{Alias: "^shamrockgreen$", Value: 634},
	{Alias: "^shockingpink$", Value: 635},
	{Alias: "^sienna$", Value: 636},
	{Alias: "^silver$", Value: 637},
	{Alias: "^sinopia$", Value: 638},
	{Alias: "^skobeloff$", Value: 639},
	{Alias: "^skyblue$", Value: 640},
	{Alias: "^skymagenta$", Value: 641},
	{Alias: "^slateblue$", Value: 642},
	{Alias: "^slategray$", Value: 643},
	{Alias: "^smalt$", Value: 644},
	{Alias: "^smokeytopaz$", Value: 645},
	{Alias: "^smokyblack$", Value: 646},
	{Alias: "^snow$", Value: 647},
	{Alias: "^spirodiscoball$", Value: 648},
	{Alias: "^springbud$", Value: 649},
	{Alias: "^springgreen$", Value: 650},
	{Alias: "^steelblue$", Value: 651},
	{Alias: "^stildegrainyellow$", Value: 652},
	{Alias: "^stizza$", Value: 653},
	{Alias: "^stormcloud$", Value: 654},
	{Alias: "^straw$", Value: 655},
	{Alias: "^sunglow$", Value: 656},
	{Alias: "^sunset$", Value: 657},
	{Alias: "^sunsetorange$", Value: 658},
	{Alias: "^tan$", Value: 659},
	{Alias: "^tangelo$", Value: 660},
	{Alias: "^tangerine$", Value: 661},
	{Alias: "^tangerineyellow$", Value: 662},
	{Alias: "^taupe$", Value: 663},
	{Alias: "^taupegray$", Value: 664},
	{Alias: "^tawny$", Value: 665},
	{Alias: "^teagreen$", Value: 666},
	{Alias: "^tearose$", Value: 667},
	{Alias: "^teal$", Value: 668},
	{Alias: "^tealblue$", Value: 669},
	{Alias: "^tealgreen$", Value: 670},
	{Alias: "^terracotta$", Value: 671},
	{Alias: "^thistle$", Value: 672},
	{Alias: "^thulianpink$", Value: 673},
	{Alias: "^ticklemepink$", Value: 674},
	{Alias: "^tiffanyblue$", Value: 675},
	{Alias: "^tigereye$", Value: 676},
	{Alias: "^timberwolf$", Value: 677},
	{Alias: "^titaniumyellow$", Value: 678},
	{Alias: "^tomato$", Value: 679},
	{Alias: "^toolbox$", Value: 680},
	{Alias: "^topaz$", Value: 681},
	{Alias: "^tractorred$", Value: 682},
	{Alias: "^trolleygrey$", Value: 683},
	{Alias: "^tropicalrainforest$", Value: 684},
	{Alias: "^trueblue$", Value: 685},
	{Alias: "^tuftsblue$", Value: 686},
	{Alias: "^tumbleweed$", Value: 687},
	{Alias: "^turkishrose$", Value: 688},
	{Alias: "^turquoise$", Value: 689},
	{Alias: "^turquoiseblue$", Value: 690},
	{Alias: "^turquoisegreen$", Value: 691},
	{Alias: "^tuscanred$", Value: 692},
	{Alias: "^twilightlavender$", Value: 693},
	{Alias: "^tyrianpurple$", Value: 694},
	{Alias: "^uablue$", Value: 695},
	{Alias: "^uared$", Value: 696},
	{Alias: "^uclablue$", Value: 697},
	{Alias: "^uclagold$", Value: 698},
	{Alias: "^ufogreen$", Value: 699},
	{Alias: "^upforestgreen$", Value: 700},
	{Alias: "^upmaroon$", Value: 701},
	{Alias: "^usccardinal$", Value: 702},
	{Alias: "^uscgold$", Value: 703},
	{Alias: "^ube$", Value: 704},
	{Alias: "^ultrapink$", Value: 705},
	{Alias: "^ultramarine$", Value: 706},
	{Alias: "^ultramarineblue$", Value: 707},
	{Alias: "^umber$", Value: 708},
	{Alias: "^unitednationsblue$", Value: 709},
	{Alias: "^universityofcaliforniagold$", Value: 710},
	{Alias: "^unmellowyellow$", Value: 711},
	{Alias: "^upsdellred$", Value: 712},
	{Alias: "^urobilin$", Value: 713},
	{Alias: "^utahcrimson$", Value: 714},
	{Alias: "^vanilla$", Value: 715},
	{Alias: "^vegasgold$", Value: 716},
	{Alias: "^venetianred$", Value: 717},
	{Alias: "^verdigris$", Value: 718},
	{Alias: "^vermilion$", Value: 719},
	{Alias: "^veronica$", Value: 720},
	{Alias: "^violet$", Value: 721},
	{Alias: "^violetblue$", Value: 722},
	{Alias: "^violetred$", Value: 723},
	{Alias: "^viridian$", Value: 724},
	{Alias: "^vividauburn$", Value: 725},
	{Alias: "^vividburgundy$", Value: 726},
	{Alias: "^vividcerise$", Value: 727},
	{Alias: "^vividtangerine$", Value: 728},
	{Alias: "^vividviolet$", Value: 729},
	{Alias: "^warmblack$", Value: 730},
	{Alias: "^waterspout$", Value: 731},
	{Alias: "^wenge$", Value: 732},
	{Alias: "^wheat$", Value: 733},
	{Alias: "^white$", Value: 734},
	{Alias: "^whitesmoke$", Value: 735},
	{Alias: "^wildstrawberry$", Value: 736},
	{Alias: "^wildwatermelon$", Value: 737},
	{Alias: "^wildblueyonder$", Value: 738},
	{Alias: "^wine$", Value: 739},
	{Alias: "^wisteria$", Value: 740},
	{Alias: "^xanadu$", Value: 741},
	{Alias: "^yaleblue$", Value: 742},
	{Alias: "^yellow$", Value: 743},
	{Alias: "^yelloworange$", Value: 744},
	{Alias: "^yellowgreen$", Value: 745},
	{Alias: "^zaffre$", Value: 746},
	{Alias: "^zinnwalditebrown$", Value: 747},
	{Alias: "^metallic$", Value: 748},
	{Alias: "^darkpurple$", Value: 749},
	{Alias: "^pewter$", Value: 750},
	{Alias: "^burn$", Value: 751},
	{Alias: "^vinous$", Value: 752},
	{Alias: "^(twotone|twotones)$", Value: 753},
	{Alias: "^lightred$", Value: 754},
	{Alias: "^lightorange$", Value: 755},
}

// ColorOrderedValues is a list of sorted values
// of Color. Use OrderedValues when you want to iterate
// through Values/Aliases and want to be sure in the order.
var ColorOrderedValues = []int32{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	17,
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	25,
	26,
	27,
	28,
	29,
	30,
	31,
	32,
	33,
	34,
	35,
	36,
	37,
	38,
	39,
	40,
	41,
	42,
	43,
	44,
	45,
	46,
	47,
	48,
	49,
	50,
	51,
	52,
	53,
	54,
	55,
	56,
	57,
	58,
	59,
	60,
	61,
	62,
	63,
	64,
	65,
	66,
	67,
	68,
	69,
	70,
	71,
	72,
	73,
	74,
	75,
	76,
	77,
	78,
	79,
	80,
	81,
	82,
	83,
	84,
	85,
	86,
	87,
	88,
	89,
	90,
	91,
	92,
	93,
	94,
	95,
	96,
	97,
	98,
	99,
	100,
	101,
	102,
	103,
	104,
	105,
	106,
	107,
	108,
	109,
	110,
	111,
	112,
	113,
	114,
	115,
	116,
	117,
	118,
	119,
	120,
	121,
	122,
	123,
	124,
	125,
	126,
	127,
	128,
	129,
	130,
	131,
	132,
	133,
	134,
	135,
	136,
	137,
	138,
	139,
	140,
	141,
	142,
	143,
	144,
	145,
	146,
	147,
	148,
	149,
	150,
	151,
	152,
	153,
	154,
	155,
	156,
	157,
	158,
	159,
	160,
	161,
	162,
	163,
	164,
	165,
	166,
	167,
	168,
	169,
	170,
	171,
	172,
	173,
	174,
	175,
	176,
	177,
	178,
	179,
	180,
	181,
	182,
	183,
	184,
	185,
	186,
	187,
	188,
	189,
	190,
	191,
	192,
	193,
	194,
	195,
	196,
	197,
	198,
	199,
	200,
	201,
	202,
	203,
	204,
	205,
	206,
	207,
	208,
	209,
	210,
	211,
	212,
	213,
	214,
	215,
	216,
	217,
	218,
	219,
	220,
	221,
	222,
	223,
	224,
	225,
	226,
	227,
	228,
	229,
	230,
	231,
	232,
	233,
	234,
	235,
	236,
	237,
	238,
	239,
	240,
	241,
	242,
	243,
	244,
	245,
	246,
	247,
	248,
	249,
	250,
	251,
	252,
	253,
	254,
	255,
	256,
	257,
	258,
	259,
	260,
	261,
	262,
	263,
	264,
	265,
	266,
	267,
	268,
	269,
	270,
	271,
	272,
	273,
	274,
	275,
	276,
	277,
	278,
	279,
	280,
	281,
	282,
	283,
	284,
	285,
	286,
	287,
	288,
	289,
	290,
	291,
	292,
	293,
	294,
	295,
	296,
	297,
	298,
	299,
	300,
	301,
	302,
	303,
	304,
	305,
	306,
	307,
	308,
	309,
	310,
	311,
	312,
	313,
	314,
	315,
	316,
	317,
	318,
	319,
	320,
	321,
	322,
	323,
	324,
	325,
	326,
	327,
	328,
	329,
	330,
	331,
	332,
	333,
	334,
	335,
	336,
	337,
	338,
	339,
	340,
	341,
	342,
	343,
	344,
	345,
	346,
	347,
	348,
	349,
	350,
	351,
	352,
	353,
	354,
	355,
	356,
	357,
	358,
	359,
	360,
	361,
	362,
	363,
	364,
	365,
	366,
	367,
	368,
	369,
	370,
	371,
	372,
	373,
	374,
	375,
	376,
	377,
	378,
	379,
	380,
	381,
	382,
	383,
	384,
	385,
	386,
	387,
	388,
	389,
	390,
	391,
	392,
	393,
	394,
	395,
	396,
	397,
	398,
	399,
	400,
	401,
	402,
	403,
	404,
	405,
	406,
	407,
	408,
	409,
	410,
	411,
	412,
	413,
	414,
	415,
	416,
	417,
	418,
	419,
	420,
	421,
	422,
	423,
	424,
	425,
	426,
	427,
	428,
	429,
	430,
	431,
	432,
	433,
	434,
	435,
	436,
	437,
	438,
	439,
	440,
	441,
	442,
	443,
	444,
	445,
	446,
	447,
	448,
	449,
	450,
	451,
	452,
	453,
	454,
	455,
	456,
	457,
	458,
	459,
	460,
	461,
	462,
	463,
	464,
	465,
	466,
	467,
	468,
	469,
	470,
	471,
	472,
	473,
	474,
	475,
	476,
	477,
	478,
	479,
	480,
	481,
	482,
	483,
	484,
	485,
	486,
	487,
	488,
	489,
	490,
	491,
	492,
	493,
	494,
	495,
	496,
	497,
	498,
	499,
	500,
	501,
	502,
	503,
	504,
	505,
	506,
	507,
	508,
	509,
	510,
	511,
	512,
	513,
	514,
	515,
	516,
	517,
	518,
	519,
	520,
	521,
	522,
	523,
	524,
	525,
	526,
	527,
	528,
	529,
	530,
	531,
	532,
	533,
	534,
	535,
	536,
	537,
	538,
	539,
	540,
	541,
	542,
	543,
	544,
	545,
	546,
	547,
	548,
	549,
	550,
	551,
	552,
	553,
	554,
	555,
	556,
	557,
	558,
	559,
	560,
	561,
	562,
	563,
	564,
	565,
	566,
	567,
	568,
	569,
	570,
	571,
	572,
	573,
	574,
	575,
	576,
	577,
	578,
	579,
	580,
	581,
	582,
	583,
	584,
	585,
	586,
	587,
	588,
	589,
	590,
	591,
	592,
	593,
	594,
	595,
	596,
	597,
	598,
	599,
	600,
	601,
	602,
	603,
	604,
	605,
	606,
	607,
	608,
	609,
	610,
	611,
	612,
	613,
	614,
	615,
	616,
	617,
	618,
	619,
	620,
	621,
	622,
	623,
	624,
	625,
	626,
	627,
	628,
	629,
	630,
	631,
	632,
	633,
	634,
	635,
	636,
	637,
	638,
	639,
	640,
	641,
	642,
	643,
	644,
	645,
	646,
	647,
	648,
	649,
	650,
	651,
	652,
	653,
	654,
	655,
	656,
	657,
	658,
	659,
	660,
	661,
	662,
	663,
	664,
	665,
	666,
	667,
	668,
	669,
	670,
	671,
	672,
	673,
	674,
	675,
	676,
	677,
	678,
	679,
	680,
	681,
	682,
	683,
	684,
	685,
	686,
	687,
	688,
	689,
	690,
	691,
	692,
	693,
	694,
	695,
	696,
	697,
	698,
	699,
	700,
	701,
	702,
	703,
	704,
	705,
	706,
	707,
	708,
	709,
	710,
	711,
	712,
	713,
	714,
	715,
	716,
	717,
	718,
	719,
	720,
	721,
	722,
	723,
	724,
	725,
	726,
	727,
	728,
	729,
	730,
	731,
	732,
	733,
	734,
	735,
	736,
	737,
	738,
	739,
	740,
	741,
	742,
	743,
	744,
	745,
	746,
	747,
	748,
	749,
	750,
	751,
	752,
	753,
	754,
	755,
}
