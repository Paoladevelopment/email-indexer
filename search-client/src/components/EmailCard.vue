<script setup lang="ts">

defineProps<{
    from: string,
    to: string,
    subject: string,
    content: string,
    highlight?: string[]
}>()

</script>

<template>
    <div class="h-[7.25rem]">
        <div class="mb-2">
            <p class="truncate text-lg"><span class="font-medium">From:</span>{{ from }}</p>
            <p class="truncate text-lg"><span class="font-medium">To:</span>{{ to }}</p>
        </div>
        <p class="truncate font-medium text-gray-900 mb-2">{{ subject }}</p>
        <p v-if="highlight" class="truncate text-sm text-gray-500" v-html="highlight[0]"></p>
        <p v-else class="truncate text-sm text-gray-500">{{ content }}</p>
    </div>
</template>