<script lang="ts" context="module">
	import Btn from '$lib/Button.svelte';
	import Sidebar from '$lib/Sidebar.svelte';
	import ProfileListEntry from '$lib/ProfileListEntry.svelte';
	import { bannerStore, modalStore, notificationStore } from '../store';
	import Banner from '$lib/banner/Banner.svelte';

	const styles = ['default', 'accent', 'success', 'warning', 'danger', 'info'];
</script>

<script lang="ts">
	const ntVariants = ['default', 'button', 'reply'];
	var item = ntVariants[Math.floor(Math.random() * ntVariants.length)];

	let sidebarVisible = false;
	let sidebarRight = false;
	let sidebarColored = false;

	function createModal(color: string): TModal {
		return {
			icon: 'cross',
			title: `Modal ${color} title`,
			text: 'Haxx0r ipsum char endif d00dz man pages protected default tcp piggyback blob else injection. Sudo strlen snarf less infinite loop break big-endian cat then hash flush overflow tunnel in ack ip semaphore Starcraft. Class wombat rm -rf bin sudo spoof crack baz highjack new vi double fatal false Dennis Ritchie throw while nak.',
			color,
			buttons: [
				{
					icon: 'cross',
					text: 'Close',
					OnClick: (close) => {
						close();
					}
				},
				{
					icon: 'cross',
					text: 'Primary',
					color,
					OnClick: (close) => {
						close();
					}
				}
			]
		};
	}

	function createNotification(color: string): TNotification {
		return {
			icon: 'cross',
			title: `Notification ${color}`,
			text: 'Haxx0r ipsum char endif d00dz man pages.',
			color,
			variant: ntVariants[Math.floor(Math.random() * ntVariants.length)],
			buttons: [
				{
					icon: 'cross',
					text: 'Primary',
					color: `transparent-${color}`,
					OnClick: () => {}
				}
			]
		};
	}

	function createBanner(color: string): TBanner {
		return {
			icon: 'cross',
			text: `This is a ${color} Banner`,
			color,
			buttons: [
				{
					icon: 'cross',
					text: 'Primary',
					color: 'white', //`white-${color}`,
					OnClick: () => {}
				}
			]
		};
	}
</script>

<template>
	<div class="">
		<div class="app-bar-button app-bar-reverse">
			<p class="w-32">Modal</p>
			<Btn
				text="all"
				on:click={(e) => styles.forEach((style) => modalStore.add(createModal(style)))}
			/>
			{#each styles as style}
				<Btn text={style} color={style} on:click={(e) => modalStore.add(createModal(style))} />
			{/each}
		</div>
		<div class="app-bar-button app-bar-reverse">
			<p class="w-32">Notification</p>
			<Btn
				text="all"
				on:click={(e) =>
					styles.forEach((style) => notificationStore.add(createNotification(style)))}
			/>
			{#each styles as style}
				<Btn
					text={style}
					color={style}
					on:click={(e) => notificationStore.add(createNotification(style))}
				/>
			{/each}
		</div>
		<div class="app-bar-button app-bar-reverse">
			<p class="w-32">Banner</p>
			<Btn
				text="all"
				on:click={(e) => styles.forEach((style) => bannerStore.add(createBanner(style)))}
			/>
			{#each styles as style}
				<Btn text={style} color={style} on:click={(e) => bannerStore.add(createBanner(style))} />
			{/each}
		</div>
		<div class="app-bar-button app-bar-reverse relative z-50">
			<p class="w-32">Sidebar</p>
			<Btn
				icon="menu"
				text="Show/Hide"
				color="default"
				on:click={(e) => (sidebarVisible = !sidebarVisible)}
			/>
			<Btn
				icon="menu"
				text="Left/Right"
				color="default"
				on:click={(e) => (sidebarRight = !sidebarRight)}
			/>
			<Btn
				icon="menu"
				text="Default/Accent"
				color="default"
				on:click={(e) => (sidebarColored = !sidebarColored)}
			/>
		</div>
	</div>
	<Sidebar
		title="Sidebar"
		img="/img/emilia.jpg"
		visible={sidebarVisible}
		colored={sidebarColored}
		right={sidebarRight}
	>
		<div slot="header">
			<Btn text="Login" color="accent" />
			<Btn text="Register" />
		</div>
		<div slot="default">
			<ProfileListEntry
				name="Elliot Alderson"
				sub="Was geht?"
				img="/img/elliot.jpg"
				status="online"
			/>
			<ProfileListEntry
				name="Tyrell Wellick"
				sub="nichts? bd"
				img="/img/tyrell.jpg"
				status="offline"
			/>
			<ProfileListEntry
				name="Darlene Alderson"
				sub="Bonsoir, Elliot."
				img="/img/darlene.jpg"
				status="away"
			/>
			<ProfileListEntry name="Whiterose" sub="did ypu get the files?" img="" status="busy" />
			<ProfileListEntry
				name="Angela Moss"
				sub="did you get the files?"
				img="/img/angela.jpg"
				status=""
			/>
		</div>
		<div slot="footer" class="app-bar-button">
			<Btn text="Login" color="accent" />
			<Btn text="Register" />
		</div>
	</Sidebar>
</template>
