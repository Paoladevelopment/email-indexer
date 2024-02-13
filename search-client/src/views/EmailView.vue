<script setup lang="ts">
import EmailComponent from '@/components/EmailComponent.vue';
import { useEmailsStore } from '@/stores/emails';
import { useRouter } from 'vue-router';

const router = useRouter()
const emailsStore = useEmailsStore()

const goBack = () => {
    router.go(-1)
}
</script>

<template>
    <main class="h-screen pt-16 w-full bg-zinc-100">
        <EmailComponent :x_from="emailsStore.currentEmail?.x_from"  :from="emailsStore.currentEmail?.from" :date="emailsStore.currentEmail?.date"
        :to="emailsStore.currentEmail?.to" :x_to="emailsStore.currentEmail?.x_to" :cc="emailsStore.currentEmail?.cc" :x_cc="emailsStore.currentEmail?.x_cc" 
        :bcc="emailsStore.currentEmail?.x_cc" :x_bcc="emailsStore.currentEmail?.x_bcc" :subject="emailsStore.currentEmail?.subject" :content="emailsStore.currentEmail?.content"
        :is-showned-half-screen="emailsStore.showHalfScreen" @go-back="goBack"/>
    </main>
</template>