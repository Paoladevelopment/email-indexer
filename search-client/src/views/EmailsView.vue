<script setup lang="ts">
import { onMounted, ref, computed } from 'vue';
import SearchComponent from '@/components/SearchComponent.vue';
import EmailCard from '@/components/EmailCard.vue';
import {useEmailsStore} from '@/stores/emails'
import { Icon } from '@iconify/vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const emailsStore = useEmailsStore()
const searchInput = ref<string>(emailsStore.searchWord)
const fromEmail = ref<number>(0)
const maxEmails = 100

interface IEmail {
    message_id: string,
    date: string,
    from: string,
    x_from: string,
    to: string,
    x_to: string,
    subject: string,
    cc: string,
    x_cc: string,
    bcc: string,
    x_bcc: string,
    content: string
}

const onSubmit = () => {
    fromEmail.value = 0
    emailsStore.FetchEmails(searchInput.value, fromEmail.value, maxEmails)
}

const goToDetails = (email: IEmail) => {
    emailsStore.setCurrentEmail(email)
    emailsStore.setSearchWord(searchInput.value)
    router.push(`/emails/${email.message_id}`)
}

const nextPage = () => {
   fromEmail.value = fromEmail.value + maxEmails
   emailsStore.FetchEmails(searchInput.value, fromEmail.value, maxEmails)
}

const prevPage = () => {
    if (fromEmail.value < maxEmails) {
        return
    }
    fromEmail.value = fromEmail.value - maxEmails
    emailsStore.FetchEmails(searchInput.value, fromEmail.value, maxEmails)
}

const fromValue = computed(() => {
    return fromEmail.value 
})

const toValue = computed(() => {
    return fromEmail.value + maxEmails
})

const isActivePrevButton = computed(() => {
    return fromEmail.value >= maxEmails
})

const isActiveNextButton = computed(() => {
    return (fromValue.value + maxEmails) <= emailsStore.totalPages
})

onMounted(async () => {
    emailsStore.FetchEmails(searchInput.value, fromEmail.value, maxEmails)
})
</script>

<template>
    <main class="h-screen w-full pt-16 bg-zinc-100">
        <form @submit.prevent="onSubmit" class="w-4/5 m-auto mb-2">
            <SearchComponent v-model="searchInput" placeholder="Search email..."/>  
        </form>
        <div class="w-95/100 mx-auto flex justify-end items-center">
            <p class="mr-2">{{fromValue}}-{{ toValue }} de {{ emailsStore.totalPages }}</p>
            <div @click="prevPage" :class="[{'text-black cursor-pointer hover:px-2 hover:bg-indigo-300 hover:rounded-full': isActivePrevButton, 'text-slate-300': !isActivePrevButton}, 'py-2']">
                <Icon icon="raphael:arrowleft" width="24" height="24" />
            </div>
            <div @click="nextPage" :class="[{'text-black cursor-pointer hover:px-2 hover:bg-indigo-300 hover:rounded-full': isActiveNextButton, 'text-slate-300': !isActiveNextButton}, 'py-2']">
                <Icon icon="raphael:arrowright" width="24" height="24"/>
            </div>
        </div>
        <section onclick="" class="w-95/100 h-4/5 mt-4 mx-auto bg-white rounded-md shadow-lg overflow-auto scrollbar-thin scrollbar-thumb-indigo-400 
            scrollbar-track-white scrollbar-thumb-rounded-md">
            <ul class="divide-y divide-gray-200">
                <li v-for="email in emailsStore.emails" :key="email.message_id" class="py-4 p-2 cursor-pointer hover:bg-indigo-50" @click="goToDetails(email)">
                    <EmailCard :from="email.from" :to="email.to" :subject="email.subject" :content="email.content"/>
                </li>
            </ul>
        </section>
    </main>
</template>