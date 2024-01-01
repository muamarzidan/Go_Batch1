#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void main_admin();
void main_user();
void menu_admin();
void menu_user();
void lihat_datauser();
void lihat_penghasilan();
void lihat_riwayatuser();
void ganti_password();
void registrasi_user();
void login_user();
void dailycleaning();
void sedottungau();
void sofacleaning();
void topup_saldo();
void topup_saldo_menu();

void main(){
    int pilihan;
    printf("1. admin\n2. user\n3. fitur without login and register\n");
    printf("\nmasukkan pilihan : ");
    scanf("%d", &pilihan);

    switch (pilihan)
    {
    case 1:
        // main_admin();
        break;
    
    case 2:
        // main_user();
        break;
    //code fitur without login and register
    case 3:
        menu_user();
        break;
    
    default:
        printf("pilihan anda tidak tersedia!\n");
        main();
        break;
    }
}

int saldo = 0;
void topup_saldo_menu() {
    int topup;
    printf("Masukkan jumlah saldo yang ingin ditambahkan: ");
    scanf("%d", &topup);
    saldo = saldo + topup;
    printf("Saldo Anda sekarang: %d\n", saldo);
}

void topup_saldo() {
    int topup;
    printf("Masukkan jumlah saldo yang ingin ditambahkan: ");
    scanf("%d", &topup);
    saldo = saldo + topup;
    printf("Saldo Anda sekarang: %d\n", saldo);
    printf("Apakah Anda ingin kembali ke menu home?");
    printf("1. Iya\n");
    printf("2. Tidak\n");
    int pilihanKembali;
    scanf("%c", &pilihanKembali);
    if (pilihanKembali == 1) {
        menu_user();
    } else if (pilihanKembali == 2) {
        return;
    } else {
        printf("Pilihan tidak tersedia.\n");
        printf("Anda akan otomatis masuk ke menu\n");
        menu_user();
    }
}

void dailyCleaning() {
    int Jam, Harga, Uang;
    char Lokasi[100];

    printf("Pengerjaan 1 jam seharga Rp. 50.000\n");
    printf("Masukkan berapa jam yang ingin dikerjakan: ");
    scanf("%d", &Jam);
    Harga = Jam * 50000;
    printf("Masukkan lokasi Anda tinggal: ");
    scanf("%s", Lokasi);
    printf("Total harga yang harus dibayar: %d\n", Harga);
    
    printf("Pilih Metode Pembayaran\n");
    printf("1. Bayar Ditempat\n");
    printf("2. Topup Online\n");
    int pilihMetode;
    scanf(" %d", &pilihMetode);

    if (pilihMetode == 1) {
        int uangTunai;
        printf("Masukkan uang tunai Anda: ");
        scanf("%d", &uangTunai);
        printf("\n");
        printf("\n");
        printf("************ Struk Pembayaran ************\n");
        printf("Nama: %s\n", "Nama"); //custom ganti ke user
        printf("Lokasi: %s\n", Lokasi);
        printf("Total harga: %d\n", Harga);
        printf("Uang yang dibayarkan: %d\n", uangTunai);
        printf("Kembalian: %d\n", uangTunai - Harga);
        printf("************ Terima Kasih ************\n");
        printf("\n");
        printf("\n");
        printf("Apakah Anda ingin kembali ke menu home?\n");
        printf("1. Iya\n");
        printf("2. Tidak\n");
        int pilihanKembali;
        scanf("%d", &pilihanKembali);
        if (pilihanKembali == 1) {
            menu_user();
        } else if (pilihanKembali == 2) {
            printf("Terima kasih telah menggunakan jasa kami.\n");
            return;
        } else {
            printf("Pilihan tidak tersedia.\n");
            printf("Anda akan otomatis masuk ke menu\n");
            menu_user();
        }
    } else if (pilihMetode == 2) {
        if (saldo >= Harga) {
            printf("Saldo anda cukup\n");
            printf("Apakah anda ingin melanjutkan pembayaran?\n");
            printf("1. Iya\n");
            printf("2.  Tidak\n");
            int pilihanPembayaran;
            scanf("%d", &pilihanPembayaran);
            if (pilihanPembayaran == 1) {
                saldo = saldo - Harga;
                printf("\n");
                printf("\n");
                printf("************ Struk Pembayaran ************\n");
                printf("Nama: %s\n", "Nama"); //custom ganti ke user
                printf("Lokasi: %s\n", Lokasi);
                printf("Total harga: %d\n", Harga);
                printf("Uang yang dibayarkan: %d\n", Harga);
                printf("Sisa Saldo %d\n", saldo);
                printf("************ Terima Kasih ************\n");
                printf("\n");
                printf("\n");

                printf("Apakah Anda ingin kembali ke menu home?");
                printf("1. Iya\n");
                printf("2. Tidak\n");
                int pilihanKembali;
                scanf(" %d", &pilihanKembali);
                if (pilihanKembali == 1) {
                    menu_user();
                } else if (pilihanKembali == 2) {
                    printf("Terima kasih telah menggunakan jasa kami.\n");
                    return;
                } else {
                    printf("Pilihan tidak tersedia.\n");
                    printf("Anda akan otomatis masuk ke menu\n");
                    menu_user();
                }
                return;
            } else if (pilihanPembayaran == 2) {
                printf("Pembayaran dibatalkan.\n");
                printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
                printf("1. Kembali ke menu home\n");
                printf("2. Melanjutkan pembayaran\n");
                int pilihanKembali;
                scanf("%d", &pilihanKembali);
                if (pilihanKembali == 1) {
                    menu_user();
                } else if (pilihanKembali == 2) {
                    dailyCleaning();
                } else {
                    printf("Pilihan tidak tersedia.\n");
                    printf("Anda akan otomatis masuk ke menu\n");
                    menu_user();
                }
                return;
            } else {
                printf("Pilihan tidak tersedia.\n");
                printf("Anda akan otomatis masuk ke menu\n");
                menu_user();
                return;
            }
        } else {
            printf("Saldo Anda saat ini: %d\n", saldo);
            printf("Saldo Anda tidak mencukupi untuk melakukan pembayaran.\n");
            printf("Apakah Anda ingin melanjutkan pembayaran? Jika Y maka Anda akan dialihkan ke menu top up saldo.\n");
            int pilihanKembali;
            printf("1. Iya\n");
            printf("2. Tidak\n");
            scanf("%d", &pilihanKembali);
            if (pilihanKembali == 1) {
                topup_saldo_menu();
                dailyCleaning();
            } else if (pilihanKembali == 2) {
                printf("Pembayaran dibatalkan.\n");
                printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
                printf("1. Kembali ke menu home\n");
                printf("2. Melanjutkan pembayaran\n");
                int pilihanKembali;
                scanf("%d", &pilihanKembali);
                if (pilihanKembali == 1) {
                    menu_user();
                } else if (pilihanKembali == 2) {
                    dailyCleaning();
                } else {
                    printf("Pilihan tidak tersedia.\n");
                    printf("Anda akan otomatis masuk ke menu\n");
                    menu_user();
                }
                return;
            } else {
                printf("Pilihan tidak tersedia.\n");
                printf("Anda akan otomatis masuk ke menu\n");
                menu_user();
                return;
            }
        }
    } else {
        printf("Pilihan tidak tersedia.\n");
        printf("Anda akan otomatis masuk ke menu\n");
        menu_user();
    }
}

void sedotTungu() {
    char Lokasi[100];
    int pilihan, Harga;

    printf("Masukkan lokasi Anda tinggal: \n");
    scanf("%s", Lokasi);

    printf("Silakan pilih ukuran:\n");
    printf("1. King size - Rp. 250.000\n");
    printf("2. Queen size - Rp. 180.000\n");
    printf("3. Full size - Rp. 150.000\n");
    printf("Masukkan pilihan Anda: ");
    scanf("%d", &pilihan);

    switch (pilihan) {
        case 1:
            Harga = 250000;
            break;
        case 2:
            Harga = 180000;
            break;
        case 3:
            Harga = 150000;
            break;
        default:
            printf("Pilihan tidak tersedia.\n");
            menu_user();
            return;
    }

    printf("Total harga yang harus dibayar: %d\n", Harga);

    printf("Pilih Metode Pembayaran\n");
    printf("1. Bayar Ditempat\n");
    printf("2. Topup Online\n");
    int pilihMetode;
    scanf(" %d", &pilihMetode);

    if (pilihMetode == 1) {
        int uangTunai;
        printf("Masukkan uang tunai Anda: ");
        scanf("%d", &uangTunai);
        
        printf("\n");
        printf("\n");
        printf("************ Struk Pembayaran ************\n");
        printf("Nama: %s\n", "Nama"); //custom ganti ke user
        printf("Jenis ukuran: %d\n", pilihan);
        printf("Total harga: %d\n", Harga);
        printf("Uang yang dibayarkan: %d\n", uangTunai);
        printf("Kembalian: %d\n", uangTunai - Harga);
        printf("************ Terima Kasih ************\n");
        printf("\n");
        printf("\n");

        printf("Apakah Anda ingin kembali ke menu home?\n");
        printf("1. Iya\n");
        printf("2.  Tidak\n");

        int pilihanKembali;
        scanf(" %d", &pilihanKembali);
        if (pilihanKembali == 1) {
            menu_user();
        } else if (pilihanKembali == 2) {
            printf("Terima kasih telah menggunakan jasa kami.\n");
            return;
        } else {
            printf("Pilihan tidak tersedia.\n");
            printf("Anda akan otomatis masuk ke menu\n");
            menu_user();
        }
    } else if (pilihMetode == 2) {
        if (saldo >= Harga) {
            printf("Saldo Anda cukup.\n");
            printf("Apakah Anda ingin melanjutkan pembayaran?\n");
            int pilihanPembayaran;
            scanf(" %d", &pilihanPembayaran);
            printf("1. Iya\n");
            printf("2. Tidak\n");
            if (pilihanPembayaran == 1) {
                saldo -= Harga;
                printf("\n");
                printf("\n");
                printf("************ Struk Pembayaran ************\n");
                printf("Nama: %s\n", "Nama"); // Ganti dengan nama pengguna
                printf("Jenis ukuran: %d\n", pilihan);
                printf("Total harga: %d\n", Harga);
                printf("Sisa Saldo %d\n", saldo);
                printf("************ Terima Kasih ************\n");
                printf("\n");
                printf("\n");

                printf("Apakah Anda ingin kembali ke menu home?");
                printf("1. Iya\n");
                printf("2. Tidak\n");

                int pilihanKembali;
                scanf(" %d", &pilihanKembali);
                if (pilihanKembali == 1) {
                    menu_user();
                } else if (pilihanKembali == 2) {
                    printf("Terima kasih telah menggunakan jasa kami.\n");
                    return;
                } else {
                    printf("Pilihan tidak tersedia.\n");
                    printf("Anda akan otomatis masuk ke menu\n");
                    menu_user();
                }
                return;
            } else if (pilihanPembayaran == 2) {
                printf("Pembayaran dibatalkan.\n");
                printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
                printf("1. Kembali ke menu home\n");
                printf("2. Melanjutkan pembayaran\n");
                int pilihanKembali;
                scanf("%d", &pilihanKembali);
                if (pilihanKembali == 1) {
                    menu_user();
                } else if (pilihanKembali == 2) {
                    sedotTungu();
                } else {
                    printf("Pilihan tidak tersedia.\n");
                    printf("Anda akan otomatis masuk ke menu\n");
                    menu_user();
                }
                return;
            } else {
                printf("Pilihan tidak tersedia.\n");
                printf("Anda akan otomatis masuk ke menu\n");
                menu_user();
                return;
            }
        } else {
            printf("Saldo Anda saat ini: %d\n", saldo);
            printf("Saldo Anda tidak mencukupi untuk melakukan pembayaran.\n");
            printf("Apakah Anda ingin melanjutkan pembayaran? Jika Y maka Anda akan dialihkan ke menu top up saldo.\n");

            int pilihanKembali;
            printf("1. Iya\n");
            printf("2. Tidak\n");
            scanf("%d", &pilihanKembali);
            if (pilihanKembali == 1) {
                topup_saldo_menu();
                sedotTungu();
            } else if (pilihanKembali == 2) {
                printf("\n");
                printf("\n");
                printf("Pembayaran dibatalkan.\n");
                printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
                printf("1. Kembali ke menu home\n");
                printf("2. Melanjutkan pembayaran\n");
                int pilihanKembali;
                scanf("%d", &pilihanKembali);
                if (pilihanKembali == 1) {
                    menu_user();
                } else if (pilihanKembali == 2) {
                    sedotTungu();
                } else {
                    printf("Pilihan tidak tersedia.\n");
                    printf("Anda akan otomatis masuk ke menu\n");
                    menu_user();
                }
                return;
            } else {
                printf("Pilihan tidak tersedia.\n");
                printf("Anda akan otomatis masuk ke menu\n");
                menu_user();
                return;
            }
        }
    } else {
        printf("Pilihan tidak tersedia.\n");
        printf("Anda akan otomatis masuk ke menu\n");
        menu_user();
    }
}

void sofaCleaning() {
    int pilihan, Harga;
    printf("Silakan pilih ukuran sofa:\n");
    printf("1. Sofa L - Rp. 300.000\n");
    printf("2. Sofa Bed - Rp. 240.000\n");
    printf("3. Sofa Jumbo - Rp. 180.000\n");
    printf("4. Sofa Standard - Rp. 120.000\n");
    printf("Masukkan pilihan Anda: ");
    scanf("%d", &pilihan);

    switch (pilihan) {
        case 1:
            Harga = 300000;
            break;
        case 2:
            Harga = 240000;
            break;
        case 3:
            Harga = 180000;
            break;
        case 4:
            Harga = 120000;
            break;
        default:
            printf("Pilihan tidak tersedia.\n");
            menu_user();
            return;
    }

    printf("Total harga yang harus dibayar: %d\n", Harga);
    char Lokasi[100];
    printf("Masukkan lokasi Anda tinggal: \n");
    scanf("%s", Lokasi);

    printf("Pilih Metode Pembayaran\n");
    printf("1. Bayar Ditempat\n");
    printf("2. Topup Online\n");
    int pilihMetode;
    scanf(" %d", &pilihMetode);

    if (pilihMetode  == 1) {
        int uangTunai;
        printf("Masukkan uang tunai Anda: ");
        scanf("%d", &uangTunai);
        printf("\n");
        printf("\n");
        printf("************ Struk Pembayaran ************\n");
        printf("Nama: %s\n", "Nama"); // Ganti dengan nama pengguna
        printf("Jenis ukuran sofa: %d\n", pilihan);
        printf("Total harga: %d\n", Harga);
        printf("Uang yang dibayarkan: %d\n", uangTunai);
        printf("Kembalian: %d\n", uangTunai - Harga);
        printf("************ Terima Kasih ************\n");
        printf("\n");
        printf("\n");

        printf("Apakah Anda ingin kembali ke menu home?\n");
        printf("1. Iya\n");
        printf("2.  Tidak\n");
        int pilihanKembali;
        scanf("%d", &pilihanKembali);
        if (pilihanKembali == 1) {
            menu_user();
        } else if (pilihanKembali == 2) {
            printf("Terima kasih telah menggunakan jasa kami.\n");
            return;
        } else {
            printf("Pilihan tidak tersedia.\n");
            printf("Anda akan otomatis masuk ke menu\n");
            menu_user();
        }
    } else if (pilihMetode == 2) {
        if (saldo >= Harga) {
            printf("Saldo Anda cukup.\n");
            printf("Apakah Anda ingin melanjutkan pembayaran?\n");
            printf("1. Iya\n");
            printf("2. Tidak\n");
            int pilihanPembayaran;
            scanf("%d", &pilihanPembayaran);
            if (pilihanPembayaran == 1) {
                saldo -= Harga;
                printf("\n");
                printf("\n");
                printf("************ Struk Pembayaran ************\n");
                printf("Nama: %s\n", "Nama"); // Ganti dengan nama pengguna
                printf("Jenis ukuran sofa: %d\n", pilihan);
                printf("Total harga: %d\n", Harga);
                printf("Sisa Saldo %d\n", saldo);
                printf("************ Terima Kasih ************\n");
                printf("\n");
                printf("\n");

                printf("Apakah Anda ingin kembali ke menu home?\n");
                printf("1. Iya\n");
                printf("2. Tidak\n");
                int pilihanKembali;
                scanf("%d", &pilihanKembali);
                if (pilihanKembali == 1) {
                    menu_user();
                } else if (pilihanKembali == 2) {
                    printf("Terima kasih telah menggunakan jasa kami.\n");
                    return;
                } else {
                    printf("Pilihan tidak tersedia.\n");
                    printf("Anda akan otomatis masuk ke menu\n");
                    menu_user();
                }
                return;
            } else if (pilihanPembayaran == 2) {
                printf("Pembayaran dibatalkan.\n");
                printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
                printf("1. Kembali ke menu home\n");
                printf("2. Melanjutkan pembayaran\n");
                int pilihanKembali;
                scanf("%d", &pilihanKembali);
                if (pilihanKembali == 1) {
                    menu_user();
                } else if (pilihanKembali == 2) {
                    sofaCleaning();
                } else {
                    printf("Pilihan tidak tersedia.\n");
                    printf("Anda akan otomatis masuk ke menu\n");
                    menu_user();
                }
                return;
            } else {
                printf("Pilihan tidak tersedia.\n");
                printf("Anda akan otomatis masuk ke menu\n");
                menu_user();
                return;
            }
        } else {
            printf("Saldo Anda saat ini: %d\n", saldo);
            printf("Saldo Anda tidak mencukupi untuk melakukan pembayaran.\n");
            printf("Apakah Anda ingin melanjutkan pembayaran? Jika Y maka Anda akan dialihkan ke menu top up saldo.\n");
            printf("1. Iya\n");
            printf("2. Tidak\n");
            int pilihanKembali;
            scanf("%d", &pilihanKembali);
            if (pilihanKembali == 1) {
                topup_saldo_menu();
                sofaCleaning();
            } else if (pilihanKembali == 2) {
                printf("Pembayaran dibatalkan.\n");
                printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
                printf("1. Kembali ke menu home\n");
                printf("2. Melanjutkan pembayaran\n");
                int pilihanKembali;
                scanf("%d", &pilihanKembali);
                if (pilihanKembali == 1) {
                    menu_user();
                } else if (pilihanKembali == 2) {
                    sofaCleaning();
                } else {
                    printf("Pilihan tidak tersedia.\n");
                    printf("Anda akan otomatis masuk ke menu\n");
                    menu_user();
                }
                return;
            } else {
                printf("Pilihan tidak tersedia.\n");
                printf("Anda akan otomatis masuk ke menu\n");
                menu_user();
                return;
            }
        }
    }
}

void menu_user() {
    int pilihan;
    printf("\n");
    printf("\n");
    printf("MENU USER\n");
    printf("1. Daily Cleaning\n2. Sedot Tungau\n3. Sofa Cleaning\n4. Top Up Saldo\n5. Keluar\n");
    printf("Masukkan pilihan: ");
    scanf("%d", &pilihan);

    switch (pilihan) {
        case 1:
            dailyCleaning();
            break;
        case 2:
            sedotTungu();
            break;
        case 3:
            sofaCleaning();
            break;
        case 4:
            topup_saldo();
            break;
        case 5:
            printf("Terima kasih, sampai jumpa!\n");
            break;
        default:
            printf("Pilihan Anda tidak tersedia!\n");
            break;
    }
}


