<script lang="ts" context="module">
    import Icon from '$lib/Icon.svelte';
import { onMount } from 'svelte';
</script>

<script lang="ts">
    export let text : string | undefined = undefined;
    export let icon : string | undefined = undefined;
    export let color : string | undefined = undefined;
    export let disabled : boolean = false;

    let ref : HTMLButtonElement | undefined;

    onMount(() => {
        ref?.focus();
    });
</script>

<template>
    <button bind:this={ref} style="outline: none;" class="app-button app-button-{color ? color : 'default'}"
         on:click>
        {#if icon}
            <Icon name="{icon}"/>
        {/if}
        {#if text}
            <p>{text}</p>
        {/if}
    </button>
</template>

<style global lang="postcss">
    
    .app-bar-button {
        @apply flex
        flex-col-reverse sm:flex-row
        sm:justify-end
        sm:p-3 p-4
        bg-gray-50;

        &.app-bar-reverse {
            @apply sm:justify-start;
        }

        & .app-button {
            &:not(:last-child) {
                @apply mt-2 sm:mt-0 mr-0 sm:mr-2;
            }
        }

        & > p {
            @apply text-gray-500 px-2 self-center;
        }
    }

    .app-button {
        @apply flex items-center
        justify-center sm:justify-start
        px-4 py-2 border rounded-md
        text-base sm:text-sm font-medium
        transition-all
        focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-50;
        
        & > *:nth-child(2) {
            @apply ml-1;
        }

        & > p {
            @apply mb-[1px];
        }

        & > *:not(.app-button-default) {
            @apply border-transparent;
        }

        &.app-button-default {
            @apply text-gray-700
            shadow-sm
            bg-gray-50
            hover:bg-white
            active:bg-gray-100
            focus:ring-accent-500;

            &:hover,
            &:focus,
            &:active {
                @apply text-gray-900;
            }
        }

        &.app-button-white {
            @apply text-accent-600
            bg-gray-100 border-gray-50
            hover:bg-gray-50 hover:border-gray-50
            active:bg-white active:border-white
            shadow-sm;
            
            &:hover > p,
            &:focus > p,
            &:active > p {
                @apply text-accent-500;
            }
        }

        &.app-button-accent {
            @apply text-white
            bg-accent-600 border-accent-700
            hover:bg-accent-500 hover:border-accent-600
            active:bg-accent-700 active:border-accent-800
            focus:ring-accent-500;
        }

        &.app-button-warning {
            @apply text-white
            bg-warning-600 border-warning-700
            hover:bg-warning-500 hover:border-warning-600
            active:bg-warning-700 active:border-warning-800
            focus:ring-warning-500;
        }

        &.app-button-success {
            @apply text-white
            bg-success-600 border-success-700
            hover:bg-success-500 hover:border-success-600
            active:bg-success-700 active:border-success-800
            focus:ring-success-500;
        }

        &.app-button-danger {
            @apply text-white
            bg-danger-600 border-danger-700
            hover:bg-danger-500 hover:border-danger-600
            active:bg-danger-700 active:border-danger-800
            focus:ring-danger-500;
        }

        &.app-button-info {
            @apply text-white
            bg-info-600 border-info-700
            hover:bg-info-500 hover:border-info-600
            active:bg-info-700 active:border-info-800
            focus:ring-info-500;
        }

        &[class*="app-button-transparent"] {
            @apply shadow-none border border-transparent;

            &:hover {
                @apply bg-gray-100 border-gray-200;
            }
            &:active {
                @apply bg-gray-200 border-gray-300;
            }
        }

        &.app-button-transparent-default {
            @apply text-gray-500;
            &:hover,
            &:active {
                @apply text-gray-700;
            }
        }

        &.app-button-transparent-accent {
            @apply text-accent-500;
        }

        &.app-button-transparent-warning {
            @apply text-warning-500;
        }

        &.app-button-transparent-success {
            @apply text-success-500;
        }

        &.app-button-transparent-danger {
            @apply text-danger-500;
        }

        &.app-button-transparent-info {
            @apply text-info-500;
        }
    }

    html.dark {
        & .app-bar-button {
            @apply bg-transparent;
            
            & > p {
                @apply text-gray-400;
            }
        }
        & .app-button {
            @apply border-2 focus:ring-offset-gray-800;

            &.app-button-default {
                @apply text-gray-300
                bg-gray-700 border-gray-700
                hover:bg-gray-700 hover:border-gray-600
                active:bg-gray-600 active:border-gray-600
                focus:ring-accent-400
                shadow-none;
                
                &:hover > p,
                &:focus > p,
                &:active > p {
                    @apply text-gray-100;
                }
            }

            &.app-button-accent {
                @apply
                bg-accent-700 border-accent-700
                hover:bg-accent-700 hover:border-accent-600
                active:bg-accent-600 active:border-accent-600
                focus:ring-accent-400;
            }

            &.app-button-warning {
                @apply
                bg-warning-700 border-warning-700
                hover:bg-warning-700 hover:border-warning-600
                active:bg-warning-600 active:border-warning-600
                focus:ring-warning-400;
            }

            &.app-button-success {
                @apply
                bg-success-700 border-success-700
                hover:bg-success-700 hover:border-success-600
                active:bg-success-600 active:border-success-600
                focus:ring-success-400;
            }

            &.app-button-danger {
                @apply
                bg-danger-700 border-danger-700
                hover:bg-danger-700 hover:border-danger-600
                active:bg-danger-600 active:border-danger-600
                focus:ring-danger-400;
            }

            &.app-button-info {
                @apply
                bg-info-700 border-info-700
                hover:bg-info-700 hover:border-info-600
                active:bg-info-600 active:border-info-600
                focus:ring-info-400;
            }

            &[class*="app-button-transparent"] {
                &:hover {
                    @apply bg-gray-700 border-gray-600;
                }
                &:active {
                    @apply bg-gray-600 border-gray-600;
                }
            }

            &.app-button-transparent-default {
                @apply text-gray-400;
                &:hover {
                    @apply text-gray-300;
                }
                &:active {
                    @apply text-gray-100;
                }
            }
        }
    }

</style>