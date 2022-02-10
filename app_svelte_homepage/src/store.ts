import { browser } from '$app/env';
import { writable } from 'svelte/store';

const SOUND_NOTIFICATION = 'window_3';
const SOUND_BANNER = 'window_2';
const SOUND_MODAL = 'window_1';

type theme = 'light' | 'dark' | 'os';

type TMenu = {
	menu: Svelte2TsxComponent;
	sourceRect: DOMRect;
}

export const theme = (function () {
	const { subscribe, update } = writable<theme>('light');
	return {
		subscribe,
		set: (theme: theme) => {
			update(val => {
				if (theme == 'os') {
					// get os theme
					val = window.matchMedia('(prefers-color-scheme: dark)').matches
						? 'dark' : 'light';
					if (browser) localStorage.removeItem('theme');
				}
				else if (browser) {
					val = theme;
					localStorage.setItem('theme', val);
				}
				// update class
				document.documentElement.classList.add('disable-transitions');
				if (val != 'dark') document.documentElement.classList.remove('dark');
				else document.documentElement.classList.add('dark');
				document.documentElement.classList.remove('disable-transitions');
				return val;
			});
		}
	};
})();

export const menuStore = (function () {
	const { subscribe, update } = writable<() => void>(() => {});
	return {
		subscribe,
		open: (close: () => void) => {
			update(val => {
				val();
				return close;
			});
		}
	};
})();

export const soundStore = (function () {
	const { subscribe, set } = writable<string>('');
	return {
		subscribe,
		play: (sound: string) => {
			set('');
			set(sound);
		}
	};
})();

export const bannerStore = (function () {
	const { subscribe, set, update } = writable<(TBanner | null)[]>([]);
	return {
		subscribe,
		add: (banner: TBanner) => {
			soundStore.play(SOUND_BANNER);
			update((value) => {
				const last = value.push(null) - 1;
				// deep copy
				value[last] = {
					...banner,
					buttons: [...banner.buttons],
					remove: () =>
						update((all) => {
							all[last] = null;
							return all;
						})
				};
				return value;
			});
		},
		clear: () => {
			set([]);
		}
	};
})();

export const notificationStore = (function () {
	const { subscribe, set, update } = writable<(TNotification | null)[]>([]);
	return {
		subscribe,
		add: (notification: TNotification) => {
			soundStore.play(SOUND_NOTIFICATION);
			update((value) => {
				const last = value.push(null) - 1;
				// deep copy
				value[last] = {
					...notification,
					buttons: [...notification.buttons],
					remove: () =>
						update((all) => {
							all[last] = null;
							return all;
						})
				};
				return value;
			});
		},
		clear: () => {
			set([]);
		}
	};
})();

export const modalStore = (function () {
	const { subscribe, set, update } = writable<(TModal | null)[]>([]);
	return {
		subscribe,
		add: (modal: TModal) => {
			update((value) => {
				const last = value.push(null) - 1;
				// deep copy
				value[last] = {
					...modal,
					buttons: [...modal.buttons],
					remove: () =>
						update((all) => {
							all[last] = null;
							return all;
						})
				};
				return value;
			});
		},
		clear: () => {
			set([]);
		}
	};
})();
