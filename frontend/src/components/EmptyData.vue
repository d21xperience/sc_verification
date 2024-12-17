<template>
    <div class="text-center flex justify-center items-center space-x-2 min-h-[350px]">
        <h3 class="text-xl">Data belum ditarik, silakan lakukan tarik data terlebih dahulu</h3>
        <button type="button" class="bg-blue-200 p-2 rounded-lg" @click="router.push({ name: 'syncDapodik' })">Tarik
            Data Sekarang</button>
    </div>
</template>

<script setup>
import router from '@/router';

</script>

<style lang="scss" scoped></style>