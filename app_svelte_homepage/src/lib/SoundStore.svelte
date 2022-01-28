<script lang="ts" context="module">
	import { onDestroy, onMount } from 'svelte';
	import { soundStore } from '$lib/../store';
	const sounds = [
		'beep_1',
		'beep_2',
		'beep_3',
		'button_1',
		'button_2',
		'click_1',
		'cling_1',
		'cling_2',
		'error_1',
		'error_2',
		'error_3',
		'gun_reload',
		'modem',
		'send_boomerang',
		'send',
		'shot',
		'splash',
		'window_1',
		'window_2',
		'window_3'
	];
</script>

<script lang="ts">
	import { Howl, Howler } from 'howler';

	let audioMap: { [key: string]: HTMLAudioElement } = {};

	onMount(() => {
		// create audio elements
		sounds.forEach((sound) => {
			audioMap[sound] = new Howl({
				src: [`/sounds/${sound}.mp3`],
				autoplay: false,
				loop: false,
				volume: 0.5,
				onend: () => {
					//audioMap[sound].currentTime = 0;
				}
			});
			audioMap[sound].load();
		});
	});

	const unsub = soundStore.subscribe((value) => {
		audioMap?.[value]?.play();
		return;
		const elem: HTMLAudioElement = audioMap[value];
		if (elem != null && value.length) {
			if (elem.played && elem.currentTime > 0.25) {
				elem.pause();
				elem.currentTime = 0;
			}
			console.log('playing sound');
			elem.play();
		}
	});

	onDestroy(unsub);
</script>
