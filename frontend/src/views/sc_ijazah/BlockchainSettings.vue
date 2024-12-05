<script setup>
// Dialog


// import Menu from 'primevue/menu';

import Dialog from 'primevue/dialog';

import DynamicDialog from 'primevue/dynamicdialog';
import { useDialog } from 'primevue/usedialog';

const dialog = useDialog();

import Button from 'primevue/button';
// import InputText from 'primevue/inputtext';


// import { faEthereum } from '@fortawesome/free-brands-svg-icons';
import { computed, ref, watch, onMounted } from 'vue';

// const checked = ref(false);
const BCNetworkSelected = ref(0)
const BCPlatform = ref()
// ambil data dari database
// const BCPlatform = ref([
//     {
//         ID: 1,
//         name: "Ethereum-Local",
//         unit: "ETH",
//         accounts: [
//             {
//                 ID: 0,
//                 name: "Account 1",
//                 address: "39TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp",
//                 amount: 100
//             },
//             {
//                 ID: 1,
//                 name: "Account 2",
//                 address: "9TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU",
//                 amount: 200
//             }
//         ]
//     },
//     {
//         ID: 2,
//         name: "Quorum",
//         unit: "ETH",
//         accounts: [
//             {
//                 ID: 0,
//                 name: "quo 1",
//                 address: "quuro9TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp",
//                 amount: 500
//             },
//             {
//                 ID: 1,
//                 name: "quo 2",
//                 address: "quuro2TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU",
//                 amount: 120
//             }
//         ]
//     },
//     {
//         ID: 3,
//         name: "Hyperledger Fabric",
//         unit: "No limit",
//         accounts: [
//             {
//                 ID: 0,
//                 name: "HF-01",
//                 address: "39TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp",
//                 amount: 350
//             },
//             {
//                 ID: 1,
//                 name: "HF-02",
//                 address: "10TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU",
//                 amount: 130
//             }
//         ]
//     },
// ])
// const BCPlatformSelected = ref(0)
const dialogSelectNetwork = ref(false)

// Ambil dari database
// const myBCAccount = ref([
//     {
//         BCPlatformID: 1,
//         BCAddress: "9TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU",
//         BCMainNetwork: {
//             URL: "infura",
//             Port: ""
//         },
//         BCTestNetwork: {
//             URL: "localhost",
//             Port: "5774"
//         },
//         BCLocalNetwork: {
//             URL: "",
//             Port: ""
//         },
//     },
//     {
//         BCPlatformID: 2,
//         BCAddress: "quorum",
//         BCMainNetwork: {
//             URL: "",
//             Port: ""
//         },
//         BCTestNetwork: {
//             URL: "",
//             Port: ""
//         },
//         BCLocalNetwork: {
//             URL: "",
//             Port: ""
//         },
//     },
//     {
//         BCPlatformID: 3,
//         BCAddress: "hyper ledger",
//         BCMainNetwork: {
//             URL: "",
//             Port: ""
//         },
//         BCTestNetwork: {
//             URL: "",
//             Port: ""
//         },
//         BCLocalNetwork: {
//             URL: "",
//             Port: ""
//         },
//     },
// ])

const dialogListAccounts = ref(false)
const dialogCreateAccount = ref(false)
const dialogCreateNewAccount = ref(false)
const dialogImportAccount = ref(false)

import Tes from "./Tes.vue"
import router from '@/router';
import QrCodeAccount from './QrCodeAccount.vue';
import axios from 'axios';


function openDialog() {
    dialog.open(Tes, {
        props: {
            header: 'Product List',
            style: {
                width: '50vw',
            },
            breakpoints: {
                '960px': '75vw',
                '640px': '90vw'
            },
            modal: true
        }
    })
}

function func_openMyQrCodeAccount() {
    dialog.open(QrCodeAccount, {
        props: {
            // header: 'Product List',
            style: {
                width: '30vw',
            },
            breakpoints: {
                '960px': '75vw',
                '640px': '90vw'
            },
            modal: true,
            // maximizable: true
        }
    })
}

// -----------
// const tes = () => {
//     alert('hello world!')
// }

// const addNewAccount = () => {
//     alert('new account')
// }

// const coba = (e) => {
//     alert(console.log(e))
// }


const sendKrypto = () => {
    router.push({ name: "sendKrypto" })
}
const menu = ref();
const items = ref([
    {
        label: 'Options',
        items: [
            {
                label: 'Refresh',
                icon: 'pi pi-refresh'
            },
            {
                label: 'Export',
                icon: 'pi pi-upload'
            },
            {
                label: 'Pengaturan',
                icon: 'pi pi-upload'
            }
        ]
    }
]);
const toggle = (event) => {
    menu.value.toggle(event);
};





const addBCNetwork = () => {
    router.push({ name: 'addBCNetworks' })
    dialogSelectNetwork.value = false
}

const networkActive = ref(false)
const networkName = ref(null)
const networkIndex = ref(0)
const selectNetwork = (e) => {
    console.log(e.target.value)
    networkIndex.value = e.target.value
    const text = e.target.textContent.trim();
    console.log(text);
    networkName.value = text
    networkActive.value = true
    dialogSelectNetwork.value = false

    // ambil data akun dari backend
    // getBcPlatform()
}
onMounted(() => {
    fetchBlockchainNetworks();
});

// Mengambil bcplatform dari backend
const fetchBlockchainNetworks = async () => {
    // menggunakan axios
    try {
        const result = await axios.get('http://localhost:8081/api/v1/accounts')
            .then(response => {
                console.log(response.data)
                BCPlatform.value = response.data
            })
    } catch (error) {
        console.error('Error fetching Ethereum price:', error);

    }
}






// Mempersingkat address
const shortenText = (text) => {
    if (text.length <= 10) return text; // Tidak dipersingkat jika terlalu pendek
    return `${text.slice(0, 5)}...${text.slice(-5)}`;
};

// Fungsi untuk mengonversi ETH ke Fiat (USD)
const fiatAmount = ref(null); // Menyimpan hasil konversi
async function convertToFiatCurrency(amount) {
    try {
        const response = await fetch(
            'https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd'
        );
        const data = await response.json();
        const ethPrice = data.ethereum.usd;
        return (amount * ethPrice).toFixed(2);
    } catch (error) {
        console.error('Error fetching Ethereum price:', error);
        return null;
    }
}

// Watcher untuk memperbarui fiatAmount saat networkIndex berubah
// watch(networkIndex, async () => {
//     const amount = BCPlatform[networkIndex.value].account[0].amount;
//     fiatAmount.value = await convertToFiatCurrency(amount);
// });

// Inisialisasi nilai awal
// (async () => {
//   const amount = BCPlatform[0].account[0].amount;
//   fiatAm

// copy to clipboard
const textToCopy = ref(null); // Reference for the text element
const isCopied = ref(false); // State for success message visibility

const copyText = async () => {
    try {
        if (textToCopy.value) {
            // Use clipboard API to copy text
            await navigator.clipboard.writeText(textToCopy.value.textContent);
            isCopied.value = true;

            // Hide the success message after 2 seconds
            setTimeout(() => {
                isCopied.value = false;
            }, 2000);
        }
    } catch (err) {
        console.error("Failed to copy text:", err);
    }
};


</script>


<template>

    <DynamicDialog />

    <div class=" w-full bg-white shadow-md rounded-lg my-4">
        <div class="flex justify-between">
            <div class="flex items-center justify-between p-4 border-b">
                <button @click="dialogSelectNetwork = true"
                    class="rounded-full bg-slate-300 py-2 px-4 hover:opacity-80">
                    <span v-if="!networkActive">Pilih Jaringan</span>
                    <span v-else>{{ BCPlatform[networkIndex].network_name }}</span>
                    <i class="ml-2 pi pi-angle-up "></i></button>
            </div>
            <div v-show="networkActive" class="flex items-center">
                <button class="bg-red-400 py-2 px-3 rounded-full hover:opacity-70 flex items-center gap-2"
                    @click="networkActive = false"><i class="pi pi-times"></i> Tutup</button>
            </div>
        </div>
    </div>
    <div class="p-4">
        <template v-if="networkActive">
            <div v-if="BCPlatform[networkIndex].accounts.length >=1" class="text-center">
                <button type="button" @click="openDialog" class="hover:opacity-70 border-b-2 shadow-sm">
                    <span>{{ BCPlatform[networkIndex].accounts[0].name }}</span> <i class="pi pi-angle-up "></i>
                </button>
                <div class="text-sm text-gray-500 flex justify-center">
                    <p> {{ shortenText(BCPlatform[networkIndex].accounts[0].address) }}</p>
                    <button @click="copyText"><i class="pi pi-copy "></i></button>
                    <div class="relative ">
                        <p class="hidden" ref="textToCopy"> {{ BCPlatform[networkIndex].accounts[0].address }}</p>
                        <p v-if="isCopied"
                            class="font-bold text-green-500 mt-2 absolute top-0 transition ease-in-out duration-500">
                            Text copied!</p>
                    </div>
                </div>
                <h1 class="text-3xl font-bold mt-2">
                    {{ BCPlatform[networkIndex].accounts[0].ammount }} {{ BCPlatform[networkIndex].unit }}
                </h1>
                <p class="text-gray-500">
                    <!-- konversikan ke mata uang -->
                    <!-- $ {{ convertToFiatCurrency(BCPlatform[networkIndex].account[0].amount) }} USD -->
                    $ {{ fiatAmount ?? 'Loading...' }} USD
                </p>
            </div>
            <div class="flex justify-center space-x-4 mt-4">
                <div class="flex flex-col items-center">
                    <Button icon="pi pi-plus" severity="info" aria-label="User" rounded />
                    <p class="text-sm mt-1">
                        Jual &amp;...
                    </p>
                </div>
                <div class="flex flex-col items-center">
                    <Button icon="pi pi-arrow-right" severity="info" aria-label="Send" rounded @click="sendKrypto" />
                    <p class="text-sm mt-1">
                        Kirim
                    </p>
                </div>
                <div class="flex flex-col items-center">
                    <Button icon="pi pi-exchange" severity="info" aria-label="pertukaran" rounded />
                    <p class="text-sm mt-1">
                        Pertukaran
                    </p>
                </div>
                <div class="flex flex-col items-center">
                    <Button icon="pi pi-random" severity="info" aria-label="User" rounded />
                    <p class="text-sm mt-1">
                        Bridge
                    </p>
                </div>
                <div class="flex flex-col items-center">
                    <Button icon="pi pi-briefcase" severity="info" aria-label="User" rounded />
                    <p class="text-sm mt-1">
                        Portofolio
                    </p>
                </div>
                <div v-show="networkIndex <= 1" class="flex flex-col items-center">
                    <Button icon="pi pi-qrcode" severity="info" aria-label="User" rounded
                        @click="func_openMyQrCodeAccount" />
                    <p class="text-sm mt-1">
                        Terima
                    </p>
                </div>

            </div>
            <div class="flex justify-center mt-6">
                <div class="flex space-x-8">
                    <a class="text-blue-500 border-b-2 border-blue-500 pb-1" href="#">
                        Token
                    </a>
                    <a class="text-gray-500" href="#">
                        NFT
                    </a>
                    <a class="text-gray-500" href="#">
                        Aktivitas
                    </a>
                </div>
            </div>
            <div class="mt-4">
                <RouterView></RouterView>
                <!-- <div class="flex items-center justify-between py-2 border-b">
                    <div class="flex items-center space-x-2">
                        <img alt="BNB logo" class="w-6 h-6" height="24"
                            src="https://storage.googleapis.com/a1aa/image/LUcCNgXFGubHChpik2eW4Mc5yj1CDFfeEHNfe7AQt4Wx8SneE.jpg"
                            width="24" />
                        <div>
                            <p class="font-medium">
                                BNB
                            </p>
                            <p class="text-sm text-gray-500">
                                BNB
                            </p>
                        </div>
                    </div>
                    <div class="text-right">
                        <p class="font-medium">
                            0 BNB
                        </p>
                        <p class="text-sm text-gray-500">
                            $0.00 USD
                        </p>
                    </div>
                </div>
                <div class="flex items-center justify-between py-2 border-b">
                    <div class="flex items-center space-x-2">
                        <font-awesome-icon :icon="faEthereum" />
                        <div>
                            <p class="font-medium">
                                USDT
                            </p>
                            <p class="text-sm text-gray-500">
                                Tether USD
                            </p>
                        </div>
                    </div>
                    <div class="text-right">
                        <p class="font-medium">
                            $4.76 USD
                        </p>
                        <p class="text-sm text-gray-500">
                            4.78 USDT
                        </p>
                    </div>
                </div>
                <div class="flex items-center justify-between py-2">
                    <div class="flex items-center space-x-2">
                        <img alt="JMPT logo" class="w-6 h-6" height="24"
                            src="https://storage.googleapis.com/a1aa/image/wl0kcOq8C9bCFxGqmEAaMHBDutEqt5nebLUmpeeTsfqkeSneE.jpg"
                            width="24" />
                        <div>
                            <p class="font-medium">
                                JMPT
                            </p>
                            <p class="text-sm text-gray-500">
                                Tether USD
                            </p>
                        </div>
                    </div>
                    <div class="text-right">
                        <p class="font-medium">
                            $4.74 USD
                        </p>
                        <p class="text-sm text-gray-500">
                            4.74 USD
                        </p>
                    </div>
                </div> -->
            </div>
        </template>
        <template v-else>
            <div class="text-center">
                <h2 class="text-3xl uppercase">Anda Belum terhubung ke jaringan Blockahain</h2>
                <p>Silahkankan <span class="font-semibold">pilih jaringan</span> yang ingin Anda gunakan</p>
                <p class="mt-5"><i class="pi pi-sort-alt-slash" style="font-size: 3rem;"></i></p>
            </div>
        </template>
        <!-- </div> -->
    </div>

    <!-- <Dialog v-model:visible="dialogListAccounts" modal header="Pilih Akun" :style="{ width: '25rem' }" position="top">
       
    </Dialog> -->
    <Dialog v-model:visible="dialogCreateAccount" modal header="Tambahkan Akun" :style="{ width: '25rem' }"
        position="top">
        <div class="flex flex-col">
            <ul>
                <li class="mb-2 hover:text-blue-500 py-2 rounded-sm ">
                    <button @click="dialogCreateNewAccount = true">
                        <i class="pi pi-plus mr-2">
                        </i>
                        Tambah akun
                    </button>
                </li>
                <li @click="dialogImportAccount = true" class="mb-2 hover:text-blue-500 py-2 rounded-sm"><button>
                        <i class="pi pi-cloud-download mr-2">
                        </i>
                        Impor akun
                    </button></li>
            </ul>
        </div>
    </Dialog>



    <!-- Dialog Create New Acount -->
    <Dialog v-model:visible="dialogCreateNewAccount" modal header="Tambahkan Akun" :style="{ width: '25rem' }"
        position="top">
        <form action="" method="post">
            <div class="flex flex-col">
                <div class="mb-2  p-2 rounded-lg cursor-pointer">
                    <label for="">Nama akun</label>
                    <input type="text" placeholder="Account - 2" class="w-full p-2">
                </div>

                <div class="flex justify-between gap-2 mt-2">
                    <Button label="Batal" icon="pi pi-times" class="w-full hover:opacity-75"
                        style="background-color: white; color: black;" @click="dialogCreateNewAccount = false" />
                    <Button label="Buat" icon="pi pi-check" class="w-full" />
                </div>
            </div>
        </form>
    </Dialog>
    <!-- ########################################### -->
    <!-- Dialog Import Akun -->
    <Dialog v-model:visible="dialogImportAccount" modal header="Impor Akun" :style="{ width: '25rem' }" position="top">
        <div class="flex flex-col">
            <p>Akun yang diimpor tidak akan ditautkan dengan Frasa Pemulihan Rahasia Akun ini.</p>
            <p>Pelajari selengkapnya</p>
            <div class="flex justify-between mt-2">
                <label for="jenis-kunci">Pilih jenis kunci</label>
                <select name="jenis-kunci" id="jenis-kunci" class="p-2">
                    <option value="0">Kebijakan pribadi</option>
                    <option value="0">File JSON</option>
                </select>
            </div>
            <div class="flex flex-col mb-4">
                <label for="string-kunci">Tempel string kunci</label>
                <div>
                    <input name="string-kunci" id="string-kunci" class="p-2 w-full border border-solid">
                    </input>
                </div>
            </div>

            <div class="flex justify-between gap-2 mt-2">
                <Button label="Batal" icon="pi pi-times" class="w-full hover:opacity-75"
                    style="background-color: white; color: black;" @click="dialogImportAccount = false" />
                <Button label="Impor" icon="pi pi-check" class="w-full" />
            </div>
        </div>
    </Dialog>

    <!-- ########################################### -->
    <!-- Dialog Select Network -->
    <Dialog v-model:visible="dialogSelectNetwork" modal header="Pilih Jaringan" position="top">
        <ul>
            <li v-for="(bc, index) in BCPlatform" :key="index"
                class="my-2 hover:bg-slate-500 p-2 cursor-pointer hover:text-white flex justify-between items-center"
                :value="index" @click="selectNetwork($event)">{{
                    bc.network_name }} </li>
            <li class=""><button @click="addBCNetwork"
                    class="px-4 py-2 hover:bg-blue-600 hover:text-white border  rounded-lg text-blue-600 font-bold"><i
                        class="pi pi-plus"></i>Tambahkan Jaringan</button></li>
        </ul>

    </Dialog>






</template>