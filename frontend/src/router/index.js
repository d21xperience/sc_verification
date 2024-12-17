import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

const router = createRouter({
  linkExactActiveClass: "bg-blue-100",
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
      meta: { title: "Home" },
    },
    {
      path: "/auth",
      // component: () => import("../views/auth/Login.vue"),
      children: [
        {
          path: "login",
          name: "login",
          component: () => import("../views/auth/Login.vue"),
          meta: { title: "Sign In" },
        },
        {
          path: "register",
          name: "register",
          component: () => import("../views/auth/Register.vue"),
          meta: { title: "Sign Up - Register" },
        },
      ],
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/AboutView.vue"),
    },
    {
      path: "/admin",

      component: () => import("../views/admin/Main.vue"),

      children: [
        {
          path: "",
          component: () => import("../views/admin/Dashboard.vue"),
          name: "admin",
          meta: { title: "Dashboard" },
        },
        {
          path: "profile",
          name: "profile",
          component: () => import("../views/admin/Profile.vue"),
          meta: { title: "Profile" },
        },

        // ----------------------------------------------------------
        // BLOCKHAIN
        // ----------------------------------------------------------
        {
          path: "blockchain",
          name: "blockchain",
          component: () => import("../views/sc_ijazah/Main.vue"),
          children: [
            {
              path: "setting",
              name: "setingBlockchain",
              meta: { title: "Blockhain setting" },
              component: () =>
                import("../views/sc_ijazah/BlockchainSettings.vue"),
              children: [
                {
                  path: "send-krypto",
                  name: "sendKrypto",
                  component: () => import("../views/sc_ijazah/SendTrx.vue"),
                },
              ],
            },
            {
              path: "list-bcnetwork",
              name: "listBCNetwork",
              meta: { title: "Daftar Blockhain" },
              component: () => import("../views/sc_ijazah/ListBCNetwork.vue"),
            },
            {
              path: "add-bcnetworks",
              name: "addBCNetworks",
              component: () => import("../views/sc_ijazah/AddBCNetwork.vue"),
            },
            {
              path: "sc-ijazah",
              name: "scIjazah",
              component: () => import("../views/sc_ijazah/SCIjazah.vue"),
            },
          ],
        },

        {
          path: "input-ijazah",
          name: "inputIjazah",
          component: () => import("../views/dapodik/DataSiswa.vue"),
          meta: { title: "Data Ijazah" },
        },

        // Data DAPODIK
        {
          path: "seting-dapodik",
          name: "syncDapodik",
          component: () => import("../views/dapodik/SetingDapodik.vue"),
        },
        {
          path: "data-sekolah",
          name: "dapodikSekolah",
          component: () => import("../views/dapodik/DataSekolah.vue"),
          meta: { title: "Data Sekolah" },
        },
        {
          path: "data-guru",
          name: "dapodikGuru",
          component: () => import("../views/dapodik/DataGuru.vue"),
          meta: { title: "Data Guru" },
        },
        {
          path: "data-siswa",
          name: "dapodikSiswa",
          component: () => import("../views/dapodik/DataSiswa.vue"),
          meta: { title: "Data Siswa" },
        },
        {
          path: "data-kelas",
          name: "dapodikKelas",
          component: () => import("../views/dapodik/DataKelas.vue"),
          meta: { title: "Data Kelas" },
        },

        // Data akademik siswa
        {
          path: "ketuntasan-rapor",
          name: "ketuntasanRapor",
          component: () => import("../views/data_akademik/KetuntasanRapor.vue"),
          meta: { title: "Ketuntasan Rapor" },
        },
        {
          path: "data-ijazah",
          name: "dataIjazah",
          component: () => import("../views/data_akademik/DataIjazah.vue"),
          meta: { title: "Data Ijazah" },
        },
      ],
    },
  ],
});

router.beforeEach((to, from, next) => {
  // document.title = to.name;
  document.title = to.meta.title || "Default Title";
  next();
});
export default router;
