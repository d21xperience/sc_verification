<script setup>
import Navbar from './components/Navbar.vue';
</script>


<template>
    <Navbar />
    <div class="font-[sans-serif] bg-white md:h-screen">
        <div class="grid md:grid-cols-2 items-center gap-8 h-full">
            <div class="max-md:order-1 p-4">
                <img src="../../tes.webp" class="lg:w-full h-full object-contain block mx-auto" alt="login-image" />
            </div>

            <div class="flex items-center md:p-8 p-6 bg-[#0C172C] h-full lg:w-11/12 lg:ml-auto">
                <form class="max-w-lg w-full mx-auto">
                    <div class="mb-12">
                        <h3 class="text-3xl font-bold text-yellow-400">Verval Ijazah</h3>
                    </div>

                    <div>
                        <label class="text-white text-xs block mb-2">Nama Sekolah/NPSN</label>
                        <div class="relative flex items-center">
                            <input name="name" type="text" required
                                class="w-full bg-transparent text-sm text-white border-b border-gray-300 focus:border-yellow-400 px-2 py-3 outline-none"
                                placeholder="Isi Nama Sekolah/NPSN" />
                            <svg xmlns="http://www.w3.org/2000/svg" fill="#bbb" stroke="#bbb"
                                class="w-[18px] h-[18px] absolute right-2" viewBox="0 0 24 24">
                                <circle cx="10" cy="7" r="6" data-original="#000000"></circle>
                                <path
                                    d="M14 15H6a5 5 0 0 0-5 5 3 3 0 0 0 3 3h12a3 3 0 0 0 3-3 5 5 0 0 0-5-5zm8-4h-2.59l.3-.29a1 1 0 0 0-1.42-1.42l-2 2a1 1 0 0 0 0 1.42l2 2a1 1 0 0 0 1.42 0 1 1 0 0 0 0-1.42l-.3-.29H22a1 1 0 0 0 0-2z"
                                    data-original="#000000"></path>
                            </svg>
                        </div>
                    </div>
                    <!-- <div class="mt-8">
                        <label class="text-white text-xs block mb-2">Email</label>
                        <div class="relative flex items-center">
                            <input name="email" type="text" required
                                class="w-full bg-transparent text-sm text-white border-b border-gray-300 focus:border-yellow-400 px-2 py-3 outline-none"
                                placeholder="Enter email" />
                            <svg xmlns="http://www.w3.org/2000/svg" fill="#bbb" stroke="#bbb"
                                class="w-[18px] h-[18px] absolute right-2" viewBox="0 0 682.667 682.667">
                                <defs>
                                    <clipPath id="a" clipPathUnits="userSpaceOnUse">
                                        <path d="M0 512h512V0H0Z" data-original="#000000"></path>
                                    </clipPath>
                                </defs>
                                <g clip-path="url(#a)" transform="matrix(1.33 0 0 -1.33 0 682.667)">
                                    <path fill="none" stroke-miterlimit="10" stroke-width="40"
                                        d="M452 444H60c-22.091 0-40-17.909-40-40v-39.446l212.127-157.782c14.17-10.54 33.576-10.54 47.746 0L492 364.554V404c0 22.091-17.909 40-40 40Z"
                                        data-original="#000000"></path>
                                    <path
                                        d="M472 274.9V107.999c0-11.027-8.972-20-20-20H60c-11.028 0-20 8.973-20 20V274.9L0 304.652V107.999c0-33.084 26.916-60 60-60h392c33.084 0 60 26.916 60 60v196.653Z"
                                        data-original="#000000"></path>
                                </g>
                            </svg>
                        </div>
                    </div> -->
                    <div class="mt-8">
                        <label class="text-white text-xs block mb-2">Nomor Ijazah</label>
                        <div class="relative flex items-center">
                            <input name="password" type="password" required
                                class="w-full bg-transparent text-sm text-white border-b border-gray-300 focus:border-yellow-400 px-2 py-3 outline-none"
                                placeholder="Masukan No.Ijazah" />
                            <svg xmlns="http://www.w3.org/2000/svg" fill="#bbb" stroke="#bbb"
                                class="w-[18px] h-[18px] absolute right-2 cursor-pointer" viewBox="0 0 128 128">
                                <path
                                    d="M64 104C22.127 104 1.367 67.496.504 65.943a4 4 0 0 1 0-3.887C1.367 60.504 22.127 24 64 24s62.633 36.504 63.496 38.057a4 4 0 0 1 0 3.887C126.633 67.496 105.873 104 64 104zM8.707 63.994C13.465 71.205 32.146 96 64 96c31.955 0 50.553-24.775 55.293-31.994C114.535 56.795 95.854 32 64 32 32.045 32 13.447 56.775 8.707 63.994zM64 88c-13.234 0-24-10.766-24-24s10.766-24 24-24 24 10.766 24 24-10.766 24-24 24zm0-40c-8.822 0-16 7.178-16 16s7.178 16 16 16 16-7.178 16-16-7.178-16-16-16z"
                                    data-original="#000000"></path>
                            </svg>
                        </div>
                    </div>

                    <div class="flex items-center mt-8">
                        <input id="remember-me" name="remember-me" type="checkbox" class="h-4 w-4 shrink-0 rounded" />
                        <label for="remember-me" class="text-white ml-3 block text-sm">
                            I accept the <a href="javascript:void(0);"
                                class="text-yellow-500 font-semibold hover:underline ml-1">Terms and Conditions</a>
                        </label>
                    </div>

                    <div class="mt-12">
                        <button type="button"
                            class="w-max shadow-xl py-3 px-6 text-sm text-gray-800 font-semibold rounded-md bg-transparent bg-yellow-400 hover:bg-yellow-500 focus:outline-none">
                            Cek Ijazah
                        </button>
                        <!-- <p class="text-sm text-white mt-8">Already have an account? <a href="javascript:void(0);"
                                class="text-yellow-400 font-semibold hover:underline ml-1">Login here</a></p> -->
                    </div>
                </form>
            </div>
        </div>
    </div>
    <div class="container">
        <h2 class="font-bold text-3xl mt-2">Apa itu Verifikasi Ijazah Blockchain</h2>
        <div class="min-h-[450px]">
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ipsam dicta ducimus neque labore ea totam, omnis
            vitae, reprehenderit laborum iusto error provident quae quam natus architecto blanditiis sint praesentium
            veritatis.
        </div>
    </div>
    <div>
        <h2 class="font-bold text-3xl mt-2">Sekolah Yang Telah Bergabung</h2>
        <div class="min-h-[450px]">
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ipsam dicta ducimus neque labore ea totam, omnis
            vitae, reprehenderit laborum iusto error provident quae quam natus architecto blanditiis sint praesentium
            veritatis.
        </div>
    </div>


    <footer class="bg-slate-900 p-10 font-[sans-serif] tracking-wide">
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
            <div class="lg:flex lg:items-center">
                <a href="javascript:void(0)">
                    <img src="../../logo.webp" alt="logo" class="w-52" />
                </a>
            </div>

            <div class="lg:flex lg:items-center">
                <ul class="flex space-x-6">
                    <li>
                        <a href="javascript:void(0)">
                            <svg xmlns="http://www.w3.org/2000/svg" class="fill-gray-300 hover:fill-white w-7 h-7"
                                viewBox="0 0 24 24">
                                <path fill-rule="evenodd"
                                    d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7v-7h-2v-3h2V8.5A3.5 3.5 0 0 1 15.5 5H18v3h-2a1 1 0 0 0-1 1v2h3v3h-3v7h4a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2z"
                                    clip-rule="evenodd" />
                            </svg>
                        </a>
                    </li>
                    <li>
                        <a href="javascript:void(0)">
                            <svg xmlns="http://www.w3.org/2000/svg" class="fill-gray-300 hover:fill-white w-7 h-7"
                                viewBox="0 0 24 24">
                                <path fill-rule="evenodd"
                                    d="M21 5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5zm-2.5 8.2v5.3h-2.79v-4.93a1.4 1.4 0 0 0-1.4-1.4c-.77 0-1.39.63-1.39 1.4v4.93h-2.79v-8.37h2.79v1.11c.48-.78 1.47-1.3 2.32-1.3 1.8 0 3.26 1.46 3.26 3.26zM6.88 8.56a1.686 1.686 0 0 0 0-3.37 1.69 1.69 0 0 0-1.69 1.69c0 .93.76 1.68 1.69 1.68zm1.39 1.57v8.37H5.5v-8.37h2.77z"
                                    clip-rule="evenodd" />
                            </svg>
                        </a>
                    </li>
                    <li>
                        <a href="javascript:void(0)">
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none"
                                class="fill-gray-300 hover:fill-white w-7 h-7" viewBox="0 0 24 24">
                                <path
                                    d="M22.92 6c-.77.35-1.6.58-2.46.69.88-.53 1.56-1.37 1.88-2.38-.83.5-1.75.85-2.72 1.05C18.83 4.5 17.72 4 16.46 4c-2.35 0-4.27 1.92-4.27 4.29 0 .34.04.67.11.98-3.56-.18-6.73-1.89-8.84-4.48-.37.63-.58 1.37-.58 2.15 0 1.49.75 2.81 1.91 3.56-.71 0-1.37-.2-1.95-.5v.03c0 2.08 1.48 3.82 3.44 4.21a4.22 4.22 0 0 1-1.93.07 4.28 4.28 0 0 0 4 2.98 8.521 8.521 0 0 1-5.33 1.84c-.34 0-.68-.02-1.02-.06C3.9 20.29 6.16 21 8.58 21c7.88 0 12.21-6.54 12.21-12.21 0-.19 0-.37-.01-.56.84-.6 1.56-1.36 2.14-2.23z" />
                            </svg>
                        </a>
                    </li>
                </ul>
            </div>

            <div>
                <h4 class="text-lg font-semibold mb-6 text-white">Contact Us</h4>
                <ul class="space-y-4">
                    <li>
                        <a href="javascript:void(0)" class="text-gray-300 hover:text-white text-sm">Email</a>
                    </li>
                    <li>
                        <a href="javascript:void(0)" class="text-gray-300 hover:text-white text-sm">Phone</a>
                    </li>
                    <li>
                        <a href="javascript:void(0)" class="text-gray-300 hover:text-white text-sm">Address</a>
                    </li>
                </ul>
            </div>

            <div>
                <h4 class="text-lg font-semibold mb-6 text-white">Information</h4>
                <ul class="space-y-4">
                    <li>
                        <a href="javascript:void(0)" class="text-gray-300 hover:text-white text-sm">About Us</a>
                    </li>
                    <li>
                        <a href="javascript:void(0)" class="text-gray-300 hover:text-white text-sm">Terms &amp;
                            Conditions</a>
                    </li>
                    <li>
                        <a href="javascript:void(0)" class="text-gray-300 hover:text-white text-sm">Privacy
                            Policy</a>
                    </li>
                </ul>
            </div>
        </div>

        <p class='text-gray-300 text-sm mt-10'>© ReadymadeUI. All rights reserved.
        </p>
    </footer>
</template>