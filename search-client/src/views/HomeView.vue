<script setup lang="ts">
import SearchComponent from '@/components/SearchComponent.vue';
import { ref } from 'vue';
const searchInput = ref<string>("")

const onSubmit = () => {
    console.log(searchInput.value);
}
</script>

<template>
    <section class="grid h-screen w-full bg-zinc-100 place-items-center">
        <div class="w-full">
            <h1 class="text-3xl mb-4 text-center md:text-5xl">En<span class="underline decoration-indigo-500">Search</span></h1>
            <form @submit.prevent="onSubmit" class="w-4/5 m-auto"> 
                <SearchComponent v-model="searchInput" placeholder="Search email..."/>  
            </form>
        </div>
    </section>
</template>
