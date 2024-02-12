<script setup lang="ts">
import { useEmailsStore } from '@/stores/emails';
import { Icon } from '@iconify/vue';
import { computed } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()
const emailsStore = useEmailsStore()

const goBack = () => {
    router.go(-1)
}

const shouldShowCC = computed(() => {
    return emailsStore.currentEmail?.cc.trim() != ""
})

const shouldShowXCC = computed(() => {
    return emailsStore.currentEmail?.x_cc.trim() != ""
})

const shouldShowBCC = computed(() => {
    return emailsStore.currentEmail?.bcc.trim() != ""
})

const shouldShowXBCC = computed(() => {
    return emailsStore.currentEmail?.x_bcc.trim() != ""
})

const highlightedContent = computed(() => {
    const content = emailsStore.currentEmail?.content || '';
    const regex = new RegExp(`(${escapeRegExp(emailsStore.searchWord)})`, 'gi');
    return content.replace(regex, '<span style="background-color: yellow;">$1</span>');
});

function escapeRegExp(string: string) {
    return string.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}
</script>

<template>
    <main class="h-screen w-full pt-16 bg-zinc-100">
        <section class="w-95/100 h-9/10 mt-4 p-4 mx-auto bg-white rounded-md shadow-lg overflow-auto scrollbar-thin scrollbar-thumb-indigo-400 
            scrollbar-track-white scrollbar-thumb-rounded-md">
            <div class="flex gap-4 items-center">
                <div class="cursor-pointer py-2 px-2 hover:bg-indigo-300 hover:rounded-full" @click="goBack">
                    <Icon icon="ep:back" width="24" height="24"/>
                </div>
                <p class="text-lg md:text-2xl font-medium">{{ emailsStore.currentEmail?.x_from }}, {{ emailsStore.currentEmail?.from }}</p>
            </div>
            <div class="p-4">
                <p class="text-lg"><span class="font-medium">Date:</span> {{ emailsStore.currentEmail?.date }}</p>
                <p class="text-lg"><span class="font-medium">To:</span> {{ emailsStore.currentEmail?.to }}</p>
                <p class="text-lg"><span class="font-medium">X_to:</span> {{ emailsStore.currentEmail?.x_to }}</p>
                <p v-if="shouldShowCC" class="md:text-lg"><span class="font-medium">CC:</span> {{ emailsStore.currentEmail?.cc }}</p>
                <p v-if="shouldShowXCC" class="md:text-lg"><span class="font-medium">X_cc:</span> {{ emailsStore.currentEmail?.x_cc }}</p>
                <p v-if="shouldShowBCC" class="md:text-lg"><span class="font-medium">BCC:</span> {{ emailsStore.currentEmail?.bcc }}</p>
                <p v-if="shouldShowXBCC" class="md:text-lg"><span class="font-medium">X_bcc:</span> {{ emailsStore.currentEmail?.x_bcc }}</p>
                <p class="whitespace-pre-line text-lg md:text-xl text-gray-700 my-6 font-medium">{{ emailsStore.currentEmail?.subject }}</p>
                <p v-html="highlightedContent" class="whitespace-pre-line"></p>
            </div>
        </section>
    </main>
</template>