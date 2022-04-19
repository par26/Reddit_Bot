import discord
import time
import os

os.startfile(os.getcwd() + "\..\Executable\main")
print(os.getcwd() + "\..\main")
infoFile = open("../fakeData.txt")
data = infoFile.readlines()
running = False

class MyClient(discord.Client):
    async def on_ready(self):
        print('Logged on as {0}!'.format(self.user))
        self.running = False
        self.counter = 0

    async def on_message(self, message):
        if message.author == client.user:
            return

        if message.content == "~break":
            self.running = False

        if message.content == "~start":
            self.running = True

            while self.running:
                await message.channel.send('**' + data[self.counter] + '**')
                self.counter += 1
                if self.running:
                    time.sleep(5)

client = MyClient()

client.run(os.environ.get('TOKEN'))