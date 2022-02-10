<script lang="ts" context="module">
	import { onDestroy, onMount } from 'svelte';
	import { bannerStore, soundStore } from '$lib/../store';
	import Banner from '../banner/Banner.svelte';
</script>

<script lang="ts">
	let banners: (TBanner | null)[] = [];

	const unsub = bannerStore.subscribe((value) => {
		banners = value;
	});

	onMount(() => {
		banners = $bannerStore;
	});
	onDestroy(unsub);
</script>

<template>
	<div class="app-banner-store">
		{#each banners as banner}
			{#if banner != null}
				<Banner {...banner} />
			{/if}
		{/each}
	</div>
</template>

<style global lang="postcss">
	.app-banner-store {
		@apply w-full h-auto overflow-visible
		flex flex-col-reverse
		transition-all;

		& > .app-banner {
			/**/
		}
	}
</style>
