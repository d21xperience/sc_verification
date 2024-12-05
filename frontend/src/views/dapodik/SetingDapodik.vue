<script setup>


import InputMask from 'primevue/inputmask';
import axios from 'axios';

import InputText from 'primevue/inputtext';

import { onMounted, ref } from 'vue';


import Select from 'primevue/select';

import ProgressBar from 'primevue/progressbar';


const dataDapodik = ref({
    baseUrl: "",
    port: "",
    username: "d21xperience@gmail.com",
    password: "Pasja123*",
    semester_id: "",
    sekolah_id: "",
    semester_id: ""
})

// Data didapat dari Database
const semester = ref([
    {
        id: '20231',
        nama: '2022/2023 Ganjil',
    },
    {
        id: '20232',
        nama: '2022/2023 Ganjil',
    },
    {
        id: '20221',
        nama: '2021/2022 Ganjil',
    },
    {
        id: '20222',
        nama: '2021/2022 Genap',
    },
    {
        id: '20231',
        nama: '2022/2023 Ganjil',
    },

])

const isConnect = ref(false)

const connecting = () => {
    // isConnect.value = !isConnect.value
    if (!isConnect) {
        alert("Aplikasi dapodik sudah terhubung")
    } else {
        alert("Aplikasi dapodik belum terhubung")
    }
}


onMounted(() => {
    axios.get('/dapo')
        .then((response) => {
            console.log(response.data)
            if (response.data) {
                isConnect.value = true
            } else {

            }
        })
        .catch((error) => console.error(error));
    // var url = "http://localhost:5774";
    // var xhr = new XMLHttpRequest();
    // xhr.open("GET", url, true);
    // xhr.onload = function () {
    //     if (xhr.status == 200) {
    //         console.log("Koneksi Berhasil");
    //     } else {
    //         console.log("Koneksi Gagal");
    //     }
    // };
    // xhr.send();
}
)

</script>



<template>
    <div class="ml-2 card flex">
        <div v-show="!isConnect">
            <form action="" class="border p-2">
                <div>
                    <label for="alamat-ip">IP Dapodik</label>
                    <div>
                        <InputMask id="basic" v-model="dataDapodik.baseUrl" mask="999.999.999.999"
                            placeholder="127.0.0.1" />
                    </div>

                </div>
                <div>
                    <label for="alamat-ip">PORT Dapodik</label>
                    <div>
                        <InputMask id="basic" v-model="dataDapodik.port" mask="9999" placeholder="5774" />
                    </div>
                </div>
                <div class="mt-2">
                    <button class="bg-blue-200 p-2 rounded-lg" type="button" @click="connecting">Cek Koneksi</button>
                </div>
            </form>
        </div>
        <div v-show="isConnect">
            <form action="" class="border p-2">
                <div>
                    <label for="username">Username</label>
                    <div>
                        <!-- <input type="text" name="username" id="username" value="d21xperience@gmail.com"> -->
                        <InputText type="text" v-model="dataDapodik.username" />
                    </div>

                </div>
                <div>
                    <label for="alamat-ip">Password</label>
                    <div>
                        <InputText type="text" v-model="dataDapodik.password" />
                    </div>
                </div>
                <div class="">
                    <label for="">Semester</label>
                    <div>
                        <Select v-model="dataDapodik.semester_id" :options="semester" optionLabel="nama"
                            placeholder="Pilih Semester" class="w-full md:w-56" />

                    </div>
                </div>
                <div class="mt-2">
                    <button class="bg-blue-200 p-2 rounded-lg" type="button">Login</button>
                </div>
            </form>
        </div>
    </div>
    <div>
        <div class="card-body">
            <div class="text-end">
            </div>
            <div class="my-3 border-top"></div>

            <table class="table-auto border-collapse border border-slate-400" id="tb-koneksi">
                <thead class="bg-primary text-light">
                    <tr>
                        <th>No</th>
                        <th>Nama Koneksi</th>
                        <th>IP Address Server Rapor SP</th>
                        <th>IP Address Server Dapodik</th>
                        <th>Key</th>
                        <th>NPSN Sekolah</th>
                        <th>Opsi</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>1</td>
                        <td>RAPORSP</td>
                        <td>localhost</td>
                        <td>localhost</td>
                        <td>mkegY2Ps2W7b4HM</td>
                        <td>20254180</td>

                        <td>
                            <button type="button" class="btn btn-outline-warning btn-sm"
                                onclick="edit_koneksi('ded48229-f449-43fb-9e91-60d930ee6b74')"><i
                                    class="fa fa-edit"></i> Edit </button>
                            <button type="button" class="btn btn-outline-danger btn-sm"
                                onclick="hapus_koneksi('ded48229-f449-43fb-9e91-60d930ee6b74')"><i
                                    class="fa fa-trash"></i> Hapus </button>

                        </td>
                    </tr>

                    <tr>
                        <td>&nbsp;</td>
                        <td>&nbsp;</td>
                        <td>&nbsp;</td>
                        <td>&nbsp;</td>
                        <td>&nbsp;</td>
                        <td>&nbsp;</td>
                        <td>&nbsp;</td>
                    </tr>
                </tbody>
                <tfoot>

                    <tr>
                        <td colspan="7">
                            <div class="text-end"> <button type="button" class="btn btn-success btn-sm"
                                    onclick="tes_koneksi()"><i class="fa fa-refresh"></i> Tes Koneksi </button>
                            </div>
                        </td>
                    </tr>
                </tfoot>
            </table>
        </div>
    </div>
    <!-- Daftar syncron -->
    <div class="ml-2">
        <button class="bg-blue-200 p-2 rounded-lg" type="button">Tarik Data</button>

        <div class="mb-2 flex flex-row">
            <h2 class="">Data Semester</h2>
            <div>
                <!-- <ProgressBar :value="50"></ProgressBar> -->

            </div>
        </div>
        <ProgressBar :value="50"></ProgressBar>
        <!-- <div class="mb-2">
            <ProgressBar :value="50"></ProgressBar>
        </div>
        <div class="mb-2">
            <ProgressBar :value="50"></ProgressBar>
        </div>
        <div class="mb-2">
            <ProgressBar :value="50"></ProgressBar>
        </div>
        <div class="mb-2">
            <ProgressBar :value="50"></ProgressBar>
        </div>
        <div class="mb-2">
            <ProgressBar :value="50"></ProgressBar>
        </div> -->
        <div>
        </div>
    </div>

</template>
