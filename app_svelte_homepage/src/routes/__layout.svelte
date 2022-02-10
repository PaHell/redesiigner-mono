<script lang="ts" context="module">
	import '../styles/app.css';
	import Header from '$lib/Header.svelte';
	import ModalStore from '$lib/modal/ModalStore.svelte';
	import NotificationStore from '$lib/notification/NotificationStore.svelte';
	import SoundStore from '$lib/SoundStore.svelte';
	import BannerStore from '$lib/banner/BannerStore.svelte';
</script>

<template>
	<BannerStore />
	<ModalStore />
	<SoundStore />
	<NotificationStore />
	<Header />
	<slot />
</template>

<style global lang="postcss">
	#svelte {
	}
</style>
