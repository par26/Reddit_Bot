import discord
import time
import os


infoFile = open("../questions.txt")
data = infoFile.read()
running = False

class MyClient(discord.Client):
    async def on_ready(self):
        print('Logged on as {0}!'.format(self.user))
        self.running = False

    async def on_message(self, message):
        if message.author == client.user:
            return

        if message.content == "~break":
            self.running = False

        if message.content == "~start":
            self.running = True
            while self.running:
                await message.channel.send('**' + data + '**')
                time.sleep(5)

client = MyClient()

client.run('OTYwNjQ4MTY5ODYyMDk4OTY2.YktfUw.HvT3sMWtKC1BI8wHuMyEsUTc4uA')
