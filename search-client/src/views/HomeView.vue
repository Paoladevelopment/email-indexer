<script setup lang="ts">
import SearchComponent from '@/components/SearchComponent.vue';
import { useEmailsStore } from '@/stores/emails';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const emailsStore = useEmailsStore()
const searchInput = ref<string>(emailsStore.searchWord)

const onSubmit = () => {
    emailsStore.setSearchWord(searchInput.value)
    router.push('/emails')
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
