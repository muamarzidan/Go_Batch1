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

FILE *registrasi;

struct admin
{
    char username[50], password[50], namauser[50];
}admin;

struct datauser
{
    int id;
    char username[50], password[50];
}user;

void main(){
    int pilihan;
    printf("= = = = = = = = = = = = = = = = \n");
    printf("= = =  ~ D'BENGKEL TelU ~ = = =\n");
    printf("= = = = = = = = = = = = = = = = \n");
    printf("= = 1. ADMIN                = =\n");
    printf("= = 2. USER                 = =\n");
    printf("\nmasukkan pilihan : ");
    scanf("%d", &pilihan);

    switch (pilihan)
    {
    case 1:
        main_admin(1);
        break;
    
    case 2:
        main_user();
        break;
    
    default:
        printf("pilihan anda tidak tersedia!\n");
        main();
        break;
    }
}

void main_admin(int percobaan){
    char username[50], password[50];

    getchar();
    printf("-=====-----=====-----=====-----\n");
    printf("===     ~ Login Admin ~     ===\n");
    printf("-=====-----=====-----=====-----\n");
    printf("Username: "); gets(username);
    printf("Password: "); gets(password);

    if (strcmp(username, "admin")==0 && strcmp(password, "123")==0){
        printf("\n -------------------------------- \n");
        printf("Anda berhasil login sebagai Admin ^_^\n");
        system("pause");
        system("cls");
        menu_admin();
    }
    else{
        printf("\n -------------------------------- \n");
        printf("Anda gagal login !!\nTersisa %d kesempatan lagi !!\n", 3-percobaan);
        if (percobaan < 3){
            system("pause");
            system("cls");
            main_admin(percobaan+1);

         }else{
            printf("\n -------------------------------- \n");
            printf("Akun anda kami banned, karena sudah melebihi limit !!\n");
            printf("Terimakasih ^_^\n");
            system("pause");
            system("cls");
            main();
         }
    }
}

void lihat_datauser(){

    struct datauser user={0,"",""};
    FILE *lihatuser;

    if((lihatuser = fopen("users.dat", "rb")) == NULL){
        printf("File could not be opened\n");
    }else{
        printf("Id\t name\t Password\t\n");
        fseek(lihatuser,0, SEEK_SET);
        while(!feof (lihatuser)){
            while(fread(&user, sizeof(struct datauser), 1, lihatuser) == 1){
                if(user.id != 0){
                    printf("%d\t %s\t %s\t\n", user.id, user.username, user.password);
                }
            }
        }
        fclose(lihatuser);
    }
    menu_admin();
}

void lihat_penghasilan(){
    printf("x");
    menu_admin();
}

void hapus_user(){
    char hapus[30];

    printf("Masukkan username akun yang ingin dihapus: ");
        scanf("%s", hapus);

    FILE *file = fopen("users.dat", "rb");
    FILE *file2 = fopen("users2.dat", "wb");

    if (file == NULL || file2 == NULL)
    {
        printf("Gagal membuka file.\n");
        return;
    }

    
    while (fread(&user, sizeof(struct datauser), 1, file))
    {
        
        if (strcmp(user.username, hapus) != 0)
        {
            
            fwrite(&user, sizeof(struct datauser), 1, file2);
        }
    }

    fclose(file);
    fclose(file2);

    // Hapus file asli
    remove("users.dat");

    // Ganti nama file sementara menjadi datauser.dat
    rename("users2.dat", "users.dat");

    printf("Akun dengan username %s berhasil dihapus.\n", hapus);
}

void ganti_password(){
    printf("x");
    menu_admin();
}


void login_user(){
    FILE *file = fopen("users.dat", "rb");
    if (file == NULL) {
        printf("File tidak ditemukan. Silakan registrasi terlebih dahulu.\n");
        main_user();
    }

    struct datauser user;
    char username[20];
    char password[20];

    printf("Masukkan username: ");
    scanf("%s", username);
    printf("Masukkan password: ");
    scanf("%s", password);

    while (fread(&user, sizeof(struct datauser), 1, file) == 1) {
        if (strcmp(user.username, username) == 0 && strcmp(user.password, password) == 0) {
            fclose(file);
            printf("Login berhasil!\n");
            menu_user();
        }
    }

    fclose(file);
    printf("Login gagal. Silakan coba lagi.\n");
    main_user(); 
}

void main_user(){
    int i = 3;
    int pilihan;
    int percobaan = 0;

    do {
        printf("Pilih opsi:\n");
        printf("1. Registrasi\n");
        printf("2. Login\n");
        printf("3. Keluar\n");
        printf("Pilihan Anda: ");
        scanf("%d", &pilihan);

        switch (pilihan) {
            case 1:
                registrasi_user();
                break;
            case 2:
                login_user();
                    percobaan = 0;
                    percobaan++;
                    if (percobaan >= i) {
                        printf("Anda sudah mencoba login sebanyak %d kali. Program keluar.\n", i);
                        exit(0);
                    }
                
                break;
            case 3:
                printf("Program keluar.\n");
                main();
                break;
            default:
                printf("Pilihan tidak valid. Silakan coba lagi.\n");
        }
    } while (pilihan != 3);

    main_user();
}

void registrasi_user(){
    struct datauser user={0,"",""};
    FILE *insertUser;
    FILE *readUser;
    int recordCount;

    if((readUser = fopen("users.dat", "rb")) == NULL){
        printf("File could not be opened\n");
    }else{
       
        fseek(readUser,0, SEEK_SET);
        while(!feof (readUser)){
           
            while(fread(&user, sizeof(struct datauser), 1, readUser) == 1){
                recordCount++;
            }
        }
        fclose(readUser);
    }

    if((insertUser = fopen("users.dat", "ab")) == NULL){
        printf("File could not be opened\n");
    }else{
        

        
            printf("Masukan name: ");
            scanf("%s", user.username);
            printf("Masukan password: ");
            scanf("%s", user.password);
            recordCount++;
            user.id = recordCount;
            //printf("%d, %s, %s", user.id, user.name, user.password);
            fseek(insertUser,0, SEEK_END);
            fwrite(&user, sizeof(struct datauser), 1, insertUser);
        
        fclose(insertUser);
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
    printf("Apakah Anda ingin kembali ke menu home? (Y/N)");
    getchar();
    char pilihanKembali;
    scanf("%c", &pilihanKembali);
    if (pilihanKembali == 'Y' || pilihanKembali == 'y') {
        menu_user();
    } else if (pilihanKembali == 'N' || pilihanKembali == 'n') {
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
    printf("Total harga yang harus dibayar: %d\n", Harga);

    printf("Masukkan lokasi Anda tinggal: ");
    scanf(" %99[^\n]", Lokasi);

    char pilihanPembayaran;
    printf("************ Masuk Ke Proses Pembayaran ************\n");

    if (saldo >= Harga) {
        printf("Saldo anda cukup\n");
        printf("Apakah Anda ingin melanjutkan pembayaran? (Y/N)\n");
        scanf(" %c", &pilihanPembayaran);
        if (pilihanPembayaran == 'Y' || pilihanPembayaran == 'y') {
            saldo = saldo - Harga;
            printf("************ Struk Pembayaran ************\n");
            printf("Nama: %s\n", user.username); //custom ganti ke user
            printf("Lokasi: %s\n", Lokasi);
            printf("Total harga: %d\n", Harga);
            printf("Uang yang dibayarkan: %d\n", Harga);
            printf("Sisa Saldo %d\n", saldo);
            printf("************ Terima Kasih ************\n");

            printf("Apakah Anda ingin kembali ke menu home? (Y/N)");
            getchar();
            char pilihanKembali;
            scanf("%c", &pilihanKembali);
            if (pilihanKembali == 'Y' || pilihanKembali == 'y') {
                menu_user();
            } else if (pilihanKembali == 'N' || pilihanKembali == 'n') {
                printf("Terima kasih telah menggunakan jasa kami.\n");
                return;
            } else {
                printf("Pilihan tidak tersedia.\n");
                printf("Anda akan otomatis masuk ke menu\n");
                menu_user();
            }
            return;
        } else if (pilihanPembayaran == 'N' || pilihanPembayaran == 'n') {
            printf("Pembayaran dibatalkan.\n");
            printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
            printf("1. Kembali ke menu home\n2. Melanjutkan pembayaran\n");
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
        printf("Apakah Anda ingin melanjutkan pembayaran? Jika Y maka Anda akan dialihkan ke menu top up saldo. (Y/N) \n");
        getchar();
        char pilihanKembali;
        scanf("%c", &pilihanKembali);
        if (pilihanKembali == 'Y' || pilihanKembali == 'y') {
            topup_saldo_menu();
            dailyCleaning();
        } else if (pilihanKembali == 'N' || pilihanKembali == 'n') {
            printf("Pembayaran dibatalkan.\n");
            printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
            printf("1. Kembali ke menu home\n2. Melanjutkan pembayaran\n");
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
}

void sedotTungu() {
    int pilihan, Harga;
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
    char pilihanPembayaran;
    printf("************ Masuk Ke Proses Pembayaran ************\n");

    if (saldo >= Harga) {
        printf("Saldo Anda cukup.\n");
        printf("Apakah Anda ingin melanjutkan pembayaran? (Y/N)\n");
        scanf(" %c", &pilihanPembayaran);
        if (pilihanPembayaran == 'Y' || pilihanPembayaran == 'y') {
            saldo -= Harga;
            printf("************ Struk Pembayaran ************\n");
            printf("Nama: %s\n", user.username); // Ganti dengan nama pengguna
            printf("Jenis ukuran: %d\n", pilihan);
            printf("Total harga: %d\n", Harga);
            printf("Sisa Saldo %d\n", saldo);
            printf("************ Terima Kasih ************\n");

            printf("Apakah Anda ingin kembali ke menu home? (Y/N)");
            getchar();
            char pilihanKembali;
            scanf(" %c", &pilihanKembali);
            if (pilihanKembali == 'Y' || pilihanKembali == 'y') {
                menu_user();
            } else if (pilihanKembali == 'N' || pilihanKembali == 'n') {
                printf("Terima kasih telah menggunakan jasa kami.\n");
                return;
            } else {
                printf("Pilihan tidak tersedia.\n");
                printf("Anda akan otomatis masuk ke menu\n");
                menu_user();
            }
            return;
        } else if (pilihanPembayaran == 'N' || pilihanPembayaran == 'n') {
            printf("Pembayaran dibatalkan.\n");
            printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
            printf("1. Kembali ke menu home\n2. Melanjutkan pembayaran\n");
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
        printf("Apakah Anda ingin melanjutkan pembayaran? Jika Y maka Anda akan dialihkan ke menu top up saldo. (Y/N)\n");
        getchar();
        char pilihanKembali;
        scanf(" %c", &pilihanKembali);
        if (pilihanKembali == 'Y' || pilihanKembali == 'y') {
            topup_saldo_menu();
            sedotTungu();
        } else if (pilihanKembali == 'N' || pilihanKembali == 'n') {
            printf("Pembayaran dibatalkan.\n");
            printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
            printf("1. Kembali ke menu home\n2. Melanjutkan pembayaran\n");
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
    char pilihanPembayaran;
    printf("************ Masuk Ke Proses Pembayaran ************\n");

    if (saldo >= Harga) {
        printf("Saldo Anda cukup.\n");
        printf("Apakah Anda ingin melanjutkan pembayaran? (Y/N)\n");
        scanf(" %c", &pilihanPembayaran);
        if (pilihanPembayaran == 'Y' || pilihanPembayaran == 'y') {
            saldo -= Harga;
            printf("************ Struk Pembayaran ************\n");
            printf("Nama: %s\n", user.username); // Ganti dengan nama pengguna
            printf("Jenis ukuran sofa: %d\n", pilihan);
            printf("Total harga: %d\n", Harga);
            printf("Sisa Saldo %d\n", saldo);
            printf("************ Terima Kasih ************\n");

            printf("Apakah Anda ingin kembali ke menu home? (Y/N)");
            getchar();
            char pilihanKembali;
            scanf(" %c", &pilihanKembali);
            if (pilihanKembali == 'Y' || pilihanKembali == 'y') {
                menu_user();
            } else if (pilihanKembali == 'N' || pilihanKembali == 'n') {
                printf("Terima kasih telah menggunakan jasa kami.\n");
                return;
            } else {
                printf("Pilihan tidak tersedia.\n");
                printf("Anda akan otomatis masuk ke menu\n");
                menu_user();
            }
            return;
        } else if (pilihanPembayaran == 'N' || pilihanPembayaran == 'n') {
            printf("Pembayaran dibatalkan.\n");
            printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
            printf("1. Kembali ke menu home\n2. Melanjutkan pembayaran\n");
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
        printf("Apakah Anda ingin melanjutkan pembayaran? Jika Y maka Anda akan dialihkan ke menu top up saldo. (Y/N)\n");
        getchar();
        char pilihanKembali;
        scanf(" %c", &pilihanKembali);
        if (pilihanKembali == 'Y' || pilihanKembali == 'y') {
            topup_saldo_menu();
            sofaCleaning();
        } else if (pilihanKembali == 'N' || pilihanKembali == 'n') {
            printf("Pembayaran dibatalkan.\n");
            printf("Apakah Anda ingin kembali ke menu home atau melanjutkan pembayaran lagi?\n");
            printf("1. Kembali ke menu home\n2. Melanjutkan pembayaran\n");
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

void menu_user() {
    int pilihan;
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
            main();
            break;
        default:
            printf("Pilihan Anda tidak tersedia!\n");
            break;
    }
}

void menu_admin(){
    int pilihan;
    printf("MENU\n");
    printf("1. melihat data user\n2. melihat penghasilan\n3. hapus user\n4. ganti password\n0. keluar\n");
    printf("masukkan pilihan : ");
    scanf("%d", &pilihan);
    
    switch (pilihan)
    {
    case 1:
        lihat_datauser();
        break;
    case 2:
        lihat_penghasilan();
        break;
    case 3:
        hapus_user();
        break;
    case 4:
        ganti_password();
    case 0:
        main();
    default:
        printf("pilihan anda tidak tersedia!\n\n");
        menu_admin();
        break;
    }
}