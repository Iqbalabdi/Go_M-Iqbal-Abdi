# (29) CI/CD

- [Summary](#Summary)
- [Praktikum](#Praktikum)

## Summary
### Pengertian
- CI/CD disebut juga dengan Continous Integration Continous Deployment. CI/CD menjembatani kesenjangan antara aktivitas pengembangan dan operasi dan tim dengan menerapkan otomatisasi dalam membangun, menguji, dan menyebarkan aplikasi.
### CI (Continous Integration)
- CI (Continous Integration) adalah automated process. Ini diperlukan untuk memberikan perintah untuk mengintegrasikan berbagai codes dari sumber potensial yang berbeda yang akan digunakan untuk build dan test
- Berikut adaalah cycle dari proses CI  
  ![image](https://user-images.githubusercontent.com/75016595/167263328-84f6d362-3dc8-4b57-8a26-b9f4e1505a5a.png)

### CD (Continous Deployment)
- CD (Continous Deployment) adalah proces untuk melakukan deploy setiap build yang telah terverifikasi
- Berikut adaalah cycle dari proses CD  
  ![image](https://user-images.githubusercontent.com/75016595/167263346-814c139e-4d56-45e8-afae-24f52a420930.png)

## Praktikum
Pada praktikum kali ini, diberikan task untuk melakukan clone pada repo `https://github.com/goFrendiAsgard/alta-batch-3-ec2`. Lalu melakukan process CI/CD menggunakan github actions

Steps : 
- Saya membuat repo testing baru di `https://github.com/Iqbalabdi/alta-ci-cd-testing`
- Cloning repo `https://github.com/goFrendiAsgard/alta-batch-3-ec2`
- Push hasil clone ke repo
- Karena untuk koneksi SSH ke server saya menggunakan password, maka pada file `deploy.yml` ada sedikit peenyesuian  
  ![hasil](./screenshots/deploy_code.jpg)
- Edit secrets di repo testing agar bisa melakukan deployment ke server  
  ![hasil](./screenshots/secrets.jpg)
- Run jobs pada Actions  
  ![hasil](./screenshots/actions.jpg)
