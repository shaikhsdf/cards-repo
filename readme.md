# First Golang Project

Hi, this is just a practise project using golang.

The project is about a card game, where we do perform the following tasks:


* Creating a new deck of cards (normal and from file)
* Shuffling the cards
* Dealing the cards
* Creating a hand
* Printing the deck of cards
* Testing every method

# Setup to run this project

To start off you will need to clone the GIT repository:

```bash
git clone https://github.com/shaikhsdf/cards-repo.git
```

You can execute the Dockerfile by running:

```bash
docker image build -t cards .
```

This will create a container with the name *cards*, which you can view by the following command:

```bash
docker image ls
```

This will create a container with the name *cards*, which you can view by the following command:

```bash
docker container run --net=host -p 5000:5000 cards
```


