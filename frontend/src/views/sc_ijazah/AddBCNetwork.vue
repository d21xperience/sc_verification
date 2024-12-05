<template>
    <div class="flex justify-between border-b-2 pb-3">
        <h2 class="text-2xl font-bold">Jaringan</h2>
        <button class="bg-red-400 py-2 px-3 rounded-full hover:opacity-70 flex items-center gap-2"
            @click="router.push({ name: 'setingBlockchain' })"><i class="pi pi-times"></i> Tutup</button>
        <!-- <button class="rounded-full border w-10 p-2 hover:bg-slate-400 hover:text-white"
            @click="router.push({ name: 'setingBlockchain' })">X</button> -->
    </div>
    <div class="mt-4">
        <form action="">
            <div>
                <label for="nama-jaringan" class="font-bold">Nama Jaringan</label>
                <input id="nama-jaringan" type="text" class="block border p-2">
            </div>
            <div>
                <label for="url-rpc" class="font-bold">URL RPC</label>
                <input id="url-rpc" type="text" class="block border">
            </div>
            <div>
                <label for="id-chain" class="font-bold">ID Chain</label>
                <input id="id-chain" type="text" class="block border">
            </div>
            <div>
                <label for="id-chain" class="font-bold">Simbol mata uang</label>
                <input id="id-chain" type="text" class="block border">
            </div>
            <div>
                <label for="id-chain" class="font-bold">ID Chain</label>
                <input id="id-chain" type="text" class="block border">
            </div>
            <div>
                <label for="id-chain"><span class="font-bold">URL block explorer</span> (opsional)</label>
                <input id="id-chain" type="text" class="block border">
            </div>

            <div class="mt-2">
                <button @click="router.push({ name: 'setingBlockchain' })"
                    class="mr-2 p-2 border rounded-full w-40 bg-blue-500 text-white hover:bg-slate-400 hover:text-black">Batal</button>
                <button class="p-2 border rounded-full w-40 bg-blue-500 text-white">Simpan &amp; gunakan</button>
            </div>
        </form>
    </div>
</template>

<script setup>
import router from '@/router';


</script>

<style lang="scss" scoped></style>