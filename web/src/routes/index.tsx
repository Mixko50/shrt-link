import { createSignal, Match, Switch } from 'solid-js';
import {
	Dialog,
	DialogPanel,
	DialogTitle,
	Transition,
	TransitionChild,
	DialogOverlay
} from 'solid-headless';
import { BasedResponse } from '~/types/response';
import { RetrieveRequest, ShrtRequest } from '~/types/request';
import { A } from 'solid-start';
import LoadingIndicator from '~/components/LoadingIndicator';

const Home = () => {
	const [slug, setSlug] = createSignal<string>('');
	const [url, setUrl] = createSignal<string>('');

	const [response, setResponse] = createSignal<BasedResponse>();

	const [isOpen, setIsOpen] = createSignal(false);
	const [isFetching, setIsFecthing] = createSignal<boolean>(false);

	const closeModal = () => {
		setIsOpen(false);
	};

	const openModal = () => {
		setIsOpen(true);
	};

	const createShortUrl = async () => {
		setIsFecthing(true);
		const req: ShrtRequest = {
			full_url: url(),
			slug: slug()
		};
		const requestOptions = {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(req)
		};
		fetch(import.meta.env.VITE_BASE_URL + '/api/create', requestOptions)
			.then((response) => response.json())
			.then((data) => {
				setResponse(data);
				openModal();
				setIsFecthing(false);
			})
			.catch(() => {
				setIsFecthing(false);
				openModal();
			});
	};

	const retrieveOriginalUrl = () => {
		setIsFecthing(true);
		const req: RetrieveRequest = {
			shrt_url: url()
		};
		const requestOptions = {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(req)
		};
		fetch(import.meta.env.VITE_BASE_URL + '/api/retrieve', requestOptions)
			.then((response) => response.json())
			.then((data) => {
				setResponse(data);
				openModal();
				setIsFecthing(false);
			})
			.catch(() => {
				setIsFecthing(false);
				openModal();
			});
	};

	return (
		<div class="bg-white">
			<div class="relative isolate px-6 pt-14 lg:px-8">
				<div class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80">
					<svg
						class="relative left-[calc(50%-11rem)] -z-10 h-[21.1875rem] max-w-none -translate-x-1/2 rotate-[30deg] sm:left-[calc(50%-30rem)] sm:h-[42.375rem]"
						viewBox="0 0 1155 678"
					>
						<path
							fill="url(#45de2b6b-92d5-4d68-a6a0-9b9b2abad533)"
							fill-opacity=".3"
							d="M317.219 518.975L203.852 678 0 438.341l317.219 80.634 204.172-286.402c1.307 132.337 45.083 346.658 209.733 145.248C936.936 126.058 882.053-94.234 1031.02 41.331c119.18 108.451 130.68 295.337 121.53 375.223L855 299l21.173 362.054-558.954-142.079z"
						/>
						<defs>
							<linearGradient
								id="45de2b6b-92d5-4d68-a6a0-9b9b2abad533"
								x1="1155.49"
								x2="-78.208"
								y1=".177"
								y2="474.645"
								gradientUnits="userSpaceOnUse"
							>
								<stop stop-color="#9089FC" />
								<stop offset="1" stop-color="#FF80B5" />
							</linearGradient>
						</defs>
					</svg>
				</div>
				{/*  */}
				<div class="mx-auto max-w-2xl py-32 sm:py-48 lg:py-56">
					<div class="text-center">
						<h1 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">
							Shrt URL
						</h1>
						<p class="mt-6 text-lg leading-8 text-gray-600">URL shortener</p>

						<div class="mt-5">
							<div class="mt-5 md:col-span-2 md:mt-0">
								<div class="overflow-hidden shadow sm:rounded-md">
									<div class="bg-white px-4 py-5 sm:p-6">
										<div class="grid grid-cols-6 gap-6">
											<div class="col-span-6 sm:col-span-4">
												<label class="block text-sm font-medium leading-6 text-gray-900">
													URL
												</label>
												<input
													type="text"
													class="mt-2 block w-full appearance-none rounded-md border-0 py-1.5 px-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
													onInput={(e) =>
														setUrl((e.target as HTMLInputElement).value)
													}
												/>
											</div>

											<div class="col-span-6 sm:col-span-2">
												<label class="block text-sm font-medium leading-6 text-gray-900">
													Slug
												</label>
												<input
													type="text"
													class="mt-2 block w-full appearance-none rounded-md border-0 py-1.5 px-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
													onInput={(e) =>
														setSlug(
															(e.target as HTMLInputElement).value
														)
													}
												/>
											</div>
										</div>
									</div>
									<div class="space-x-2 pb-5">
										<Switch>
											<Match when={isFetching()}>
												<LoadingIndicator />
											</Match>
											<Match when={!isFetching()}>
												<button
													type="submit"
													class="h-10 rounded-md bg-[#E96479] py-2 px-3 text-sm font-semibold text-white shadow-sm hover:bg-[#C3ACD0] focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
													onClick={createShortUrl}
												>
													Get shorten url
												</button>
												<button
													type="submit"
													class="h-10 rounded-md bg-[#7286D3] py-2 px-3 text-sm font-semibold text-white shadow-sm hover:bg-[#C3ACD0] focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-[indigo-500]"
													onClick={retrieveOriginalUrl}
												>
													Retrieve original url
												</button>
											</Match>
										</Switch>
									</div>
								</div>
							</div>
						</div>
					</div>
					{/* / */}
				</div>
				<div class="absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]">
					<svg
						class="relative left-[calc(50%+3rem)] h-[21.1875rem] max-w-none -translate-x-1/2 sm:left-[calc(50%+36rem)] sm:h-[30.375rem]"
						viewBox="0 0 1155 678"
					>
						<path
							fill="url(#ecb5b0c9-546c-4772-8c71-4d3f06d544bc)"
							fill-opacity=".3"
							d="M317.219 518.975L203.852 678 0 438.341l317.219 80.634 204.172-286.402c1.307 132.337 45.083 346.658 209.733 145.248C936.936 126.058 882.053-94.234 1031.02 41.331c119.18 108.451 130.68 295.337 121.53 375.223L855 299l21.173 362.054-558.954-142.079z"
						/>
						<defs>
							<linearGradient
								id="ecb5b0c9-546c-4772-8c71-4d3f06d544bc"
								x1="1155.49"
								x2="-78.208"
								y1=".177"
								y2="474.645"
								gradientUnits="userSpaceOnUse"
							>
								<stop stop-color="#9089FC" />
								<stop offset="1" stop-color="#FF80B5" />
							</linearGradient>
						</defs>
					</svg>
				</div>
			</div>

			{/* Dialog */}
			<Transition appear show={isOpen()}>
				<Dialog isOpen class="fixed inset-0 z-10 overflow-y-auto" onClose={closeModal}>
					<div class="flex min-h-screen items-center justify-center px-4">
						<TransitionChild
							enter="ease-out duration-300"
							enterFrom="opacity-0"
							enterTo="opacity-100"
							leave="ease-in duration-200"
							leaveFrom="opacity-100"
							leaveTo="opacity-0"
						>
							<DialogOverlay class="fixed inset-0 bg-gray-900 bg-opacity-50" />
						</TransitionChild>

						{/* This element is to trick the browser into centering the modal contents. */}
						<span class="inline-block h-screen align-middle" aria-hidden="true">
							&#8203;
						</span>
						<TransitionChild
							enter="ease-out duration-300"
							enterFrom="opacity-0 scale-95"
							enterTo="opacity-100 scale-100"
							leave="ease-in duration-200"
							leaveFrom="opacity-100 scale-100"
							leaveTo="opacity-0 scale-95"
						>
							<DialogPanel class="my-8 inline-block transform overflow-hidden rounded-xl bg-gray-50 p-6 text-left align-middle shadow-xl transition-all dark:border dark:border-gray-50">
								<DialogTitle
									as="h3"
									class="text-xl font-medium leading-6 text-gray-900"
								>
									{response()?.success
										? 'Success!'
										: response()?.error?.error_message ??
										  'Something went wrong'}
								</DialogTitle>
								<Switch>
									<Match when={response()?.success}>
										<div class="flex-col">
											<div class="mt-2 flex">
												<p class="text-lg font-medium">
													Your shrt link:&nbsp
												</p>
												&nbsp
												<p class="text-lg">
													{import.meta.env.VITE_BASE_URL +
														'/' +
														response()?.data.slug}
												</p>
											</div>
											<div class="flex justify-center">
												<img
													src={`https://chart.googleapis.com/chart?cht=qr&chs=300x300&chl=${encodeURIComponent(
														import.meta.env.VITE_BASE_URL +
															'/' +
															response()?.data.slug
													)}`}
													alt="qr"
													class="mt-2"
												/>
											</div>
										</div>
									</Match>
									<Match when={!response()?.success}>
										<div class="mt-2 flex flex-col">
											<p class="text-lg">
												{response()?.error?.detail ??
													'An error occurred while retrieving data'}
											</p>
										</div>
									</Match>
								</Switch>

								<div class="mt-4 flex justify-center space-x-2">
									<Switch>
										<Match when={response()?.success}>
											<button
												type="button"
												class="h-10 w-full justify-center rounded-md bg-[#C780FA] px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto "
												onClick={() => {
													navigator.clipboard.writeText(
														`${import.meta.env.VITE_BASE_URL}/${
															response()?.data.slug
														}`
													);
													closeModal();
												}}
											>
												Copy
											</button>
											<A
												class="h-10 w-full justify-center rounded-md bg-[#FFB26B] px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto "
												onClick={closeModal}
												href={`${import.meta.env.VITE_BASE_URL}/${
													response()?.data.slug ?? url()
												}`}
												target="_blank"
											>
												Open url
											</A>
										</Match>
										<Match when={!response()?.success}>
											<button
												type="button"
												class="h-10 w-full justify-center rounded-md bg-[#C780FA] px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto "
												onClick={closeModal}
											>
												Ok
											</button>
										</Match>
									</Switch>
								</div>
							</DialogPanel>
						</TransitionChild>
					</div>
				</Dialog>
			</Transition>
		</div>
	);
};

export default Home;
