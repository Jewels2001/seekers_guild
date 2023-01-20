<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import TheWelcome from './components/TheWelcome.vue'
import NavBar from './components/NavBar.vue'
import Banner from './components/Banner.vue'
</script>

<template>
  <div class="body background bg-slate-900 h-max w-max">
    <NavBar />
    <!-- <header>
      
      <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />

      <div class="wrapper"> -->
      <div class="mx-1">
        <RouterView />
      </div>
      <!-- </div>
    </header> -->
    
    <!-- <main>
      <TheWelcome />
    </main> -->
    <div class="absolute max-h-10% bottom-0 w-screen">
      <Banner />
    </div>
  </div>
</template>

<style scoped>
</style>
