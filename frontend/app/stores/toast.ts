import { type Toast, type ToastType } from '../types/toast';

export const useToastsStore = defineStore('toast', () =>
	{
		const toasts  = ref<Toast[]>([]);
		const innerID = ref(0);

		const addToast = (type: ToastType, text: string) =>
		{
			const toast: Toast =
			{
				id: getNextID(),
				type: type,
				text: text
			};

			toasts.value.push(toast);

			setTimeout(() => removeToast(toast.id), 3000);
		}

		const removeToast = (id: number) => toasts.value = toasts.value.filter((toast: Toast) => toast.id !== id);
		const getNextID   = () => innerID.value++;

		return {
			toasts,

			addToast,
			removeToast
		};
	}
)