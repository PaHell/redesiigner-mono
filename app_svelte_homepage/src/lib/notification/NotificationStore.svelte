<script lang="ts" context="module">
	import { onDestroy, onMount } from 'svelte';
	import { notificationStore } from '$lib/../store';
	import Notification from './Notification.svelte';
</script>

<script lang="ts">
	let notifications: (TNotification | null)[] = [];

	const unsub = notificationStore.subscribe((value) => {
		notifications = value;
	});

	onMount(() => {
		notifications = $notificationStore;
	});
	onDestroy(unsub);
</script>

<template>
	<div class="app-notification-store">
		{#each notifications as nt}
			{#if nt != null}
				<Notification {...nt} />
			{/if}
		{/each}
	</div>
</template>

<style global lang="postcss">
	.app-notification-store {
		@apply fixed z-30
        w-full h-0
        sm:w-0 sm:h-full
        top-0 right-0
        sm:top-4 sm:right-4
        overflow-visible
        sm:flex flex-col justify-start items-end;

		& > .app-nt-wrapper {
			@apply w-screen sm:w-96;
		}
	}
</style>
