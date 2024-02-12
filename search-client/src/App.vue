<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { Icon } from '@iconify/vue';
</script>

<template>
    <header class="w-full bg-white absolute z-1">
      <nav class="flex items-center">
          <RouterLink to="/" class="basis-5/100 py-2 px-5" active-class="bg-zinc-100">
              <Icon icon="tabler:mail-search" class="m-auto w-6 h-6 md:w-7 md:h-7"/>
          </RouterLink>
          <RouterLink to="/emails" class="basis-5/100 py-2 px-5 md:text-lg" active-class="bg-zinc-100">Emails</RouterLink>
      </nav>
    </header>
    <RouterView />
</template>
