<template>
    <div>
        <div class="max-w-4xl mx-auto p-6 bg-white shadow-md rounded-lg mt-10">
            <div class="flex flex-wrap">
                <div class="w-1/2">
                    <div class="flex items-center space-x-6">
                        <div class="relative">
                            <img alt="Profile picture of a person" class="w-32 h-32 rounded-full block" height="150"
                                src="https://storage.googleapis.com/a1aa/image/bS2nVV2XXD7IPNInU3xBCiOkojwX5fyy6Ne1rlaANQ9Rqs0TA.jpg"
                                width="150" />
                            <div
                                class="absolute  w-full h-full top-0 left-0 flex items-center justify-center opacity-0 rounded-full hover:opacity-100 bg-black bg-opacity-50 text-gray-100 transition ease-in-out duration-500">
                                <i @click="triggerFileInput" class="pi pi-camera cursor-pointer "
                                    style="font-size: 1.5rem;"></i>
                                <!-- Input file (disembunyikan) -->
                                <input id="fileInput" type="file" accept="image/*" style="display: none;"
                                    @change="handleFileSelect" />
                            </div>
                        </div>
                        <div class="text-center">
                            <h1 class="text-2xl font-bold">
                                John Doe
                            </h1>
                            <p class="text-gray-600">
                                Admin
                            </p>
                        </div>
                    </div>
                    <div class="mt-6">
                        <h2 class="text-xl font-semibold">
                            Personal Information
                        </h2>
                        <div class="mt-4">
                            <p>
                                <span class="font-bold">
                                    Email:
                                </span>
                                john.doe@example.com
                            </p>
                            <p>
                                <span class="font-bold">
                                    Phone:
                                </span>
                                +123 456 7890
                            </p>
                            <p>
                                <span class="font-bold">
                                    Date of Birth:
                                </span>
                                January 1, 1990
                            </p>
                        </div>
                    </div>
                </div>
                <div>
                    <div class="mt-6">
                        <h2 class="text-xl font-semibold">
                            Address
                        </h2>
                        <div class="mt-4">
                            <p>
                                <span class="font-bold">
                                    Street:
                                </span>
                                123 Main St
                            </p>
                            <p>
                                <span class="font-bold">
                                    City:
                                </span>
                                Springfield
                            </p>
                            <p>
                                <span class="font-bold">
                                    State:
                                </span>
                                IL
                            </p>
                            <p>
                                <span class="font-bold">
                                    Zip Code:
                                </span>
                                62704
                            </p>
                        </div>
                    </div>
                    <div class="mt-6">
                        <h2 class="text-xl font-semibold">
                            Parents
                        </h2>
                        <div class="mt-4">
                            <p>
                                <span class="font-bold">
                                    Father's Name:
                                </span>
                                Richard Doe
                            </p>
                            <p>
                                <span class="font-bold">
                                    Mother's Name:
                                </span>
                                Jane Doe
                            </p>
                            <p>
                                <span class="font-bold">
                                    Father's Phone:
                                </span>
                                +123 456 7891
                            </p>
                            <p>
                                <span class="font-bold">
                                    Mother's Phone:
                                </span>
                                +123 456 7892
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="max-w-4xl mx-auto p-2 bg-white shadow-md rounded-lg mt-10">
            <h3 class="font-bold text-xl">Change email & password</h3>
            <form action="" method="post">
                <div>
                    <label for="email">Email</label>
                    <input class="block border w-full" type="email" name="email" id="email">
                </div>
                <div>
                    <label for="password">Password</label>
                    <input class="block border w-full" type="password" name="password" id="password">
                </div>
                <div>
                    <button>Batal</button>
                    <button>Simpan</button>
                </div>

            </form>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

// import Tabs from 'primevue/tabs';
// import TabList from 'primevue/tablist';
// import Tab from 'primevue/tab';
// import TabPanels from 'primevue/tabpanels';
// import TabPanel from 'primevue/tabpanel';



// State untuk menyimpan file yang dipilih
const selectedFile = ref(null);

// Fungsi untuk menangani pemilihan file
const handleFileSelect = (event) => {
    const file = event.target.files[0];
    if (file) {
        selectedFile.value = file; // Simpan file yang dipilih
        console.log("Selected file:", file);
    }
};

// Fungsi untuk memicu input file secara programatis
const triggerFileInput = () => {
    const fileInput = document.getElementById("fileInput");
    if (fileInput) fileInput.click();
};
</script>

<style lang="scss" scoped></style>