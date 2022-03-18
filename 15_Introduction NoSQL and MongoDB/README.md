# (15) Introduction NoSQL and MongoDB

- [Summary](#Summary)
- [Praktikum](#Praktikum)

# Summary

## A. Preface

- SQL memiliki property `ACID`
  - Atomic, terjadi atau tidak sama sekali
  - Cosistency, data yang masuk merupakan data yang sudah memenuhi aturan
  - Isolation, ketika terjadi transaksi secara sekuensial maka transaksi akan dilakukan secara berurut
  - Durability, terjamin bahwa data yang sudah tersimpan pasti tersimpan

## B. NoSQL

- Bukan merupakan relasional database
- Memberikan flexibilitas
- Perubahan (penambahan/pengurangan) struktur entitas tidak dipermasalkan
- Kelebihan lain:
  - Tanpa skema
  - Mendukung 
  - Konfigurasi clustering
- When not to use:
  - Bisnis finansial
  - ACID compliant
- Jasa jasa NoSQL mempunyai arsitektur berbeda-beda.

|Tipe NoSQL|Deskripsi|
|---|---|
|`key`:`value`|Menggunakan prisip `hash` biasa diguanakan untuk penyimpanan sementara misal cache|
|`column`:`family`|Penyimpanan data dilakukan dengan struktur kolom, misal On Line Analitic Processing|
|`graph`|Struktur data dengan network yang dihubungkan dengan `node`, misal Social Network Analysis|
|`document-based`|Kumpulan dokumen (data sebagai dokumen) sejenis, misal unstructured data: loging (elastic search,MongoDB)|

_redis, elastic, dan mongo biasa digunakan (cassandra untuk logging)_

## C. MongoDB

- Membuat Database

```js
use `database`
```

- Membuat Collection 

```js
db.createCollection(`nama_collection`)
```

- Memasukan dokumen kedalam collection

```js
db.collectionName.insert[One||Many](`json`||`jsonArray`)
```

- Mencari dokumen pada collection

```js
db.collectionName.find()[.pretty()]
```
_pretty() menjadikan tampilan dokumen pada output menjadi lebih mudah dibaca_

|Sintaks|Deskripsi|
|---|---|
|`$set`|update data|
|`$inc`|_increment_ data|
|`$unset`|menghilangkan data|

- Update dokumen

```js
db.collectionName.udpate(One||Many)({query},{$set{atribToChange}})
```

- Hapus dokumen

```js
db.collectionName.delete(One||Many)({query})
```
- Tambahan:
  - Terdapat fungsi angregat dan operasi yang dapat menunjang operasi
  - Advance query (dokumen dalam dokumen) menggunakan kutip untuk mengakses atribut.

# Praktikum

## 1. Create

### a. 5 operator

```
db.operators.insertMany([
  {
    _id:1,
    name:'AA',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:2,
    name:'BB',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:3,
    name:'CC',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:4,
    name:'DD',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:5,
    name:'EE',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
]);
```

### b. 3 product type

```
db.product_types.insertMany([
  {
    _id:1,
    name:'Makanan',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:2,
    name:'Minuman',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:3,
    name:'Snack',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
]);
```

### c. 2 product dengan type id 1 dan operator 3

```
db.products.insertMany([
  {
    _id:1,
    name:'Nasi Goreng',
    type_id: 1,
    operator_id: 3,
    code: '101',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:2,
    name:'Mie Goreng',
    type_id: 1,
    operator_id: 3,
    code: '102',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
]);
```

### d. 3 product dengan type id 2 dan operator 1

```
db.products.insertMany([
  {
    _id:3,
    name:'Aqua',
    type_id: 2,
    operator_id: 1,
    code: '201',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:4,
    name:'Ades',
    type_id: 2,
    operator_id: 1,
    code: '202',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:5,
    name:'Teh Botol',
    type_id: 2,
    operator_id: 1,
    code: '203',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
]);
```

### e. 3 product dengan type id 3 dan operator 4

```
db.products.insertMany([
  {
    _id:6,
    name:'Lays',
    type_id: 3,
    operator_id: 4,
    code: '301',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:7,
    name:'Doritos',
    type_id: 3,
    operator_id: 4,
    code: '302',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:8,
    name:'Citato',
    type_id: 3,
    operator_id: 4,
    code: '303',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
]);
```
### f. deskripsi produk

```
db.product_descriptions.insertMany([
  {
    _id:1,
    product_id: 1,
    description:'Nasi goreng kampung',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:2,
    product_id: 2,
    description:'Mie goreng jawa dengan telur',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:3,
    product_id: 3,
    description:'Air mineral 600ml',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:4,
    product_id: 4,
    description:'Air mineral 300ml',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:5,
    product_id: 5,
    description:'Teh botol 50ml',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:6,
    product_id: 6,
    description:'Kripik kentang rumput laut',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:7,
    product_id: 7,
    description:'Kripik singkong barberkyu',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:8,
    product_id: 8,
    description:'Kripik kentang barberkyu',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
]);
```
### g. 3 payment method

```
db.payment_methods.insertMany([
  {
    _id:1,
    name:'Transfer',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:2,
    name:'Call order delivery',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:3,
    name:'Virtual account',
    status: 1,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
]);
```
### h. 5 user

```
db.users.insertMany([
  {
    _id:1,
    name: 'AB',
    status:1,
    dob: new Date("2000-07-30"),
    gender: 'M',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:2,
    name: 'BC',
    status:1,
    dob: new Date("2000-07-30"),
    gender: 'F',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:3,
    name: 'CD',
    status:1,
    dob: new Date("2000-07-30"),
    gender: 'M',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:4,
    name: 'DE',
    status:1,
    dob: new Date("2000-07-30"),
    gender: 'F',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:5,
    name: 'EF',
    status:1,
    dob: new Date("2000-07-30"),
    gender: 'M',
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
]);
```
### i. 3 transkasi per user

```
db.transactions.insertMany([
  {
    _id:1,
    user_id:1,
    payment_method_id:1,
    status: 'barhasil',
    total_quantity: 3,
    total_price: 30000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:2,
    user_id:1,
    payment_method_id:2,
    status: 'gagal',
    total_quantity: 6,
    total_price: 50000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:3,
    user_id:1,
    payment_method_id:3,
    status: 'menunggu pembayaran',
    total_quantity: 15,
    total_price: 75000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:4,
    user_id:2,
    payment_method_id:1,
    status: 'barhasil',
    total_quantity: 3,
    total_price: 30000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:5,
    user_id:2,
    payment_method_id:2,
    status: 'gagal',
    total_quantity: 6,
    total_price: 50000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:6,
    user_id:2,
    payment_method_id:3,
    status: 'menunggu pembayaran',
    total_quantity: 15,
    total_price: 75000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:7,
    user_id:3,
    payment_method_id:1,
    status: 'barhasil',
    total_quantity: 3,
    total_price: 30000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:8,
    user_id:3,
    payment_method_id:2,
    status: 'gagal',
    total_quantity: 6,
    total_price: 50000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:9,
    user_id:3,
    payment_method_id:3,
    status: 'menunggu pembayaran',
    total_quantity: 15,
    total_price: 75000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:10,
    user_id:4,
    payment_method_id:1,
    status: 'barhasil',
    total_quantity: 3,
    total_price: 30000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:11,
    user_id:4,
    payment_method_id:2,
    status: 'gagal',
    total_quantity: 6,
    total_price: 50000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:12,
    user_id:4,
    payment_method_id:3,
    status: 'menunggu pembayaran',
    total_quantity: 15,
    total_price: 75000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:13,
    user_id:5,
    payment_method_id:1,
    status: 'barhasil',
    total_quantity: 3,
    total_price: 30000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:14,
    user_id:5,
    payment_method_id:2,
    status: 'gagal',
    total_quantity: 6,
    total_price: 50000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:15,
    user_id:5,
    payment_method_id:3,
    status: 'menunggu pembayaran',
    total_quantity: 15,
    total_price: 75000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
]);
```
### j. 3 produk per transaksi

```
db.transaction_details.insertMany([
  {
    _id:1,
    transaction_id: 1,
    product_id: 1,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:2,
    transaction_id: 1,
    product_id: 3,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:3,
    transaction_id: 1,
    product_id: 6,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:4,
    transaction_id: 2,
    product_id: 2,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:5,
    transaction_id: 2,
    product_id: 4,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:6,
    transaction_id: 2,
    product_id: 7,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:7,
    transaction_id: 3,
    product_id: 1,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:8,
    transaction_id: 3,
    product_id: 5,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:9,
    transaction_id: 3,
    product_id: 8,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:10,
    transaction_id: 4,
    product_id: 1,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:11,
    transaction_id: 4,
    product_id: 3,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:12,
    transaction_id: 4,
    product_id: 6,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:13,
    transaction_id: 5,
    product_id: 2,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:14,
    transaction_id: 5,
    product_id: 4,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:15,
    transaction_id: 5,
    product_id: 7,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:16,
    transaction_id: 6,
    product_id: 1,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:17,
    transaction_id: 6,
    product_id: 5,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:18,
    transaction_id: 6,
    product_id: 8,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:19,
    transaction_id: 7,
    product_id: 1,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:20,
    transaction_id: 7,
    product_id: 3,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:21,
    transaction_id: 7,
    product_id: 6,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:22,
    transaction_id: 8,
    product_id: 2,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:23,
    transaction_id: 8,
    product_id: 4,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:24,
    transaction_id: 8,
    product_id: 7,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:25,
    transaction_id: 9,
    product_id: 1,
    status: 'tersedia',
    qty: 5,
    price: 50000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:26,
    transaction_id: 9,
    product_id: 5,
    status: 'tersedia',
    qty: 5,
    price: 50000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:27,
    transaction_id: 9,
    product_id: 8,
    status: 'tersedia',
    qty: 5,
    price: 50000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:28,
    transaction_id: 10,
    product_id: 1,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:29,
    transaction_id: 10,
    product_id: 3,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:30,
    transaction_id: 10,
    product_id: 6,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:31,
    transaction_id: 11,
    product_id: 2,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:32,
    transaction_id: 11,
    product_id: 4,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:33,
    transaction_id: 11,
    product_id: 7,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:34,
    transaction_id: 12,
    product_id: 1,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:35,
    transaction_id: 12,
    product_id: 5,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:36,
    transaction_id: 12,
    product_id: 8,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:37,
    transaction_id: 13,
    product_id: 1,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:38,
    transaction_id: 13,
    product_id: 3,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:39,
    transaction_id: 13,
    product_id: 6,
    status: 'tersedia',
    qty: 1,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:40,
    transaction_id: 14,
    product_id: 2,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:41,
    transaction_id: 14,
    product_id: 4,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:42,
    transaction_id: 14,
    product_id: 7,
    status: 'tersedia',
    qty: 2,
    price: 10000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:43,
    transaction_id: 15,
    product_id: 1,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:44,
    transaction_id: 15,
    product_id: 5,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
  {
    _id:45,
    transaction_id: 15,
    product_id: 8,
    status: 'tersedia',
    qty: 5,
    price: 5000,
    created_at: new Timestamp(),
    updated_at: new Timestamp(),
  },
]);
```

## 2. Read

### a. Tampilkan user laki

```
db.users.find({gender: {$eq:'M'}});
```

![hasil](./screenshots/2-a.png)  

### b. produk dengan id 3

```
db.products.find({_id: {$eq:3}});
```

![hasil](./screenshots/2-b.png)

### c. jumlah user perempuan

```
db.users.count({gender: {$eq: 'F'}});
```

![hasil](./screenshots/2-c.png)

### d. urutkan pengguna berdasarkan nama

```
db.users.find().sort( { name: 1 } );
```

![hasil](./screenshots/2-d.png)

### e. tampilkan 5 data produk

```
db.products.find().limit(5);
```

![hasil](./screenshots/2-e.png)

## 3. Update

### a. update produk id 1 dengan nama baru

```
db.products.update({_id:{$eq:1}},{$set:{name:'product dummy'}});
```

![hasil](./screenshots/3-a.png)

### b. ubah jumalh produc id 1 pada detail transaksi menjadi 3

```
db.transaction_details.updateMany({product_id:{$eq:3}},{$set:{qty:3}});
```
![hasil](./screenshots/3-b.png)

## 4. Delete

### a. dokumen produk id 1

```
db.products.deleteOne({_id:{$eq:1}});
```
![hasil](./screenshots/4-a.png)

### b. hapus dokume produk denga tipe 1

```
db.products.deleteMany({type_id:{$eq:1}});
```
![hasil](./screenshots/4-b.png)
