<script setup lang="ts">
import { useEmailsStore } from '@/stores/emails';
import { computed } from 'vue';
import { Icon } from '@iconify/vue';

const emailsStore = useEmailsStore()

const props = defineProps<{
    x_from?: string,
    from?: string,
    date?: string,
    to?: string,
    x_to?: string,
    cc?: string,
    x_cc?: string,
    bcc?: string,
    x_bcc?: string,
    subject?: string,
    content?: string,
    isShownedHalfScreen?: boolean
}>()

const emit = defineEmits(['go-back', 'close'])

const goBack = () => {
    emit('go-back')
}

const close = () => {
    emit('close')
}

const shouldShowCC = computed(() => {
    return props.cc !== undefined && props.cc.trim() !== ""
})

const shouldShowXCC = computed(() => {
    return props.x_cc !== undefined && props.x_cc.trim() !== ""
})

const shouldShowBCC = computed(() => {
    return props.bcc !== undefined && props.bcc.trim() !== ""
})

const shouldShowXBCC = computed(() => {
    return props.x_bcc !== undefined && props.x_bcc.trim() !== ""
})


const highlightedContent = computed(() => {
    const content = props.content || '';
    const regex = new RegExp(`(${escapeRegExp(emailsStore.searchWord)})`, 'gi');
    return content.replace(regex, '<span style="background-color: yellow;">$1</span>');
});

function escapeRegExp(string: string) {
    return string.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}

</script>

<template>
    <section :class="[{'mx-auto w-95/100 h-9/10 mt-4 overflow-auto scrollbar-thin scrollbar-thumb-indigo-400 scrollbar-track-white scrollbar-thumb-rounded-md': !isShownedHalfScreen}, {'min-h-full': isShownedHalfScreen}, 'p-4 bg-white rounded-md shadow-lg']">
        <div class="flex gap-4 items-center">
            <div v-if="!isShownedHalfScreen" class="cursor-pointer py-2 px-2 hover:bg-indigo-300 hover:rounded-full" @click="goBack">
                <Icon icon="ep:back" width="24" height="24"/>
            </div>
            <div v-if="isShownedHalfScreen" class="cursor-pointer py-2 px-2 hover:bg-indigo-300 hover:rounded-full" @click="close">
                <Icon icon="iconoir:cancel" width="24" height="24" />
            </div>
            <p class="text-lg md:text-2xl font-medium">{{ x_from }}, {{ from }}</p>
        </div>
        <div class="p-4">
            <p class="text-lg"><span class="font-medium">Date:</span> {{ date }}</p>
            <p class="text-lg"><span class="font-medium">To:</span> {{ to }}</p>
            <p class="text-lg"><span class="font-medium">X_to:</span> {{ x_to }}</p>
            <p v-if="shouldShowCC" class="md:text-lg"><span class="font-medium">CC:</span> {{ cc }}</p>
            <p v-if="shouldShowXCC" class="md:text-lg"><span class="font-medium">X_cc:</span> {{ x_cc }}</p>
            <p v-if="shouldShowBCC" class="md:text-lg"><span class="font-medium">BCC:</span> {{ bcc }} </p>
            <p v-if="shouldShowXBCC" class="md:text-lg"><span class="font-medium">X_bcc:</span> {{ x_bcc }}</p>
            <p class="whitespace-pre-line text-lg md:text-xl text-gray-700 my-6 font-medium">{{ emailsStore.currentEmail?.subject }}</p>
            <p v-html="highlightedContent" class="whitespace-pre-line"></p>
        </div>
    </section>
</template>