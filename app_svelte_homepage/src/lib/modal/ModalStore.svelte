<script lang="ts" context="module">
    import { onDestroy, onMount } from 'svelte';
    import { modalStore, soundStore } from '$lib/../store';
    import Modal from './Modal.svelte';

    const SOUND_OPEN = "window_3";
</script>

<script lang="ts">
    let modal : TModal|null = null;
    let hide = true;
    
    const unsub = modalStore.subscribe((value) => {
        modal = value[value.findIndex(m => !!m)];
        if (!modal) setTimeout(() => hide = true, 250);
        else {
            soundStore.play("window_3");
            hide = false;
        }
    });

    onMount(() => {
        modal = $modalStore[0];
        hide = !!modal;
        if (modal) soundStore.play("window_3");
    });
    onDestroy(unsub);
</script>

<template>
    {#if !hide}
        <div class="app-modal-store">
            <div class="app-modal-bg animate__animated {modal ? 'animate__fadeIn' : 'animate__fadeOut'}"></div>
            <div class="app-modal-ct">
                {#if modal}
                    <Modal {...modal}/>
                {/if}
            </div>
        </div>
    {/if}
</template>

<style global lang="postcss">
    .app-modal-store {
        @apply fixed z-10
        overflow-y-auto overflow-x-hidden
        flex justify-center items-center
        w-full h-full p-4 pb-24 inset-0;

        & > .app-modal-bg {
            @apply fixed inset-0
            bg-gray-300 bg-opacity-50;
        }

        & > .app-modal-ct {
            @apply w-full max-w-lg;
        }
    }

    html.dark .app-modal-store {
        & > .app-modal-bg {
            @apply bg-gray-900 bg-opacity-75;
        }
    }
</style>
