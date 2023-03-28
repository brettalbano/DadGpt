<script lang='ts'>
    import { page } from '$app/stores';
	import ChatMessage from '../../../components/ChatMessage.svelte';
    import { ConverseWithDad } from '../../../api/conversation';

	let nameMe='Me';
	let profilePicMe='https://p0.pikist.com/photos/474/706/boy-portrait-outdoors-facial-men-s-young-t-shirt-hair-person-thumbnail.jpg';

	let nameChatPartner='Dad GPT';
	let profilePicChatPartner='https://storage.needpix.com/rsynced_images/male-teacher-cartoon.jpg';

    let todayMessages = [{
        "messageId": 0,
        "message": "Welcome! What topic would you like me to create a joke for you to roll your eyes on?",
        "sentByMe": false,
        }
    ]

    let newMessage = ''
    let currentMessageId = 1
    function HandleClick() {
        let newMessageStruct = {
            messageId: currentMessageId++,
            message: newMessage,
            sentByMe: true,

            // "messageId": currentMessageId++,
            // "message": newMessage,
            // "sentByMe": true,
        }
        ConverseWithDad(newMessageStruct, Number($page.params['userId'])).then(
            response => {
                todayMessages = [...todayMessages, {
                    messageId: currentMessageId++,
                    message: response['message'],
                    sentByMe: false
                }]
            }
        )

        todayMessages = [...todayMessages, newMessageStruct]

        newMessage = ''
    }

</script>



<style>
	.direct-chat .card-body {
		overflow-x: hidden;
		padding: 0;
		position: relative;
	}

	.direct-chat-messages {
		-webkit-transform: translate(0, 0);
		transform: translate(0, 0);
		height: 400px;
		overflow: auto;
		padding: 10px;
		transition: -webkit-transform .5s ease-in-out;
		transition: transform .5s ease-in-out;
		transition: transform .5s ease-in-out, -webkit-transform .5s ease-in-out;
	}

	.contacts-img {
		border-radius: 50%;
		width: 40px;
		height: 40px;
	}
	.contacts-name {
		margin-left: 15px;
		font-weight: 600;
	}
</style>

<svelte:head>
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
</svelte:head>

<div class="card card-danger direct-chat direct-chat-danger">
    <div class="card-header">
        <div class="card-tools d-flex">
            <!-- <aside class="absolute w-full h-full bg-gray-200 border-r-2 shadow-lg" class:open>

            </aside> -->
            <img class="contacts-img" src={profilePicChatPartner} alt="profilePic">
            <span class="contacts-name">{nameChatPartner}</span>
            <span class="mr-auto"></span>
        </div>
    </div>
    <div class="card-body">
        <div class="direct-chat-messages">
			{#each todayMessages as todayMessage}
                <ChatMessage
                    nameMe = {nameMe}
                    profilePicMe = {profilePicMe}
                    nameChatPartner = {nameChatPartner}
                    profilePicChatPartner = {profilePicChatPartner}
					message={todayMessage.message}
					sentByMe={todayMessage.sentByMe}
					isToday={true}
				/>
            {/each}

        </div>
    </div>
    <div class="card-footer">
        <div class="input-group">
            <input on:keypress={(e) => e.key === 'Enter' && HandleClick()} type="text" placeholder="Type Message ..." class="form-control" bind:value={newMessage}>
            <span class="input-group-append">
                <button on:click={()=> HandleClick()} type="button" class="btn btn-primary">Send</button>
            </span>
        </div>
    </div>
</div>
