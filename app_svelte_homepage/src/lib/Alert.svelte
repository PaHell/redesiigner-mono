<script lang="ts" context="module">
      import { onMount } from 'svelte';
	import Icon from '$lib/Icon.svelte';
</script>

<script lang="ts">
	export let text : string = "";
	export let icon : string | undefined;
	export let color: string | undefined;
	export let bg: string | undefined;
</script>

<template>
      {#if text}
      <div class={`app-alert app-alert-color-${color} app-alert-bg-${bg}`}>
                  {#if icon}
                        <Icon name={icon}/>
                  {/if}
                  <p>{text}</p>
            </div>
      {/if}
</template>

<style global lang="postcss">
      .app-alert {
            @apply flex items-center
            rounded-md border p-2;

            & > .app-icon {
                  @apply mr-2 text-inherit;
            }

            & > p {
                  @apply text-sm;
            }

            &.app-alert-color-default { @apply text-gray-600; }
            &.app-alert-color-accent { @apply text-accent-500; }
            &.app-alert-color-success { @apply text-success-500; }
            &.app-alert-color-warning { @apply text-warning-600; }
            &.app-alert-color-danger { @apply text-danger-500; }
            &.app-alert-color-info { @apply text-info-500; }
      }
</style>
