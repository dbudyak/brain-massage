import matplotlib.pyplot as plt
import tensorflow as tf

""" 
Week 2 example 
https://colab.research.google.com/github/lmoroney/dlaicourse/blob/master/Course%201%20-%20Part%204%20-%20Lesson%202%20-%20Notebook.ipynb
"""


class myCallback(tf.keras.callbacks.Callback):
    def on_epoch_end(self, epoch, logs={}):
        if (logs.get('loss') < 0.4):
            print("\nReached 60% accuracy so cancelling training!")
            self.model.stop_training = True


mnist = tf.keras.datasets.fashion_mnist
(training_images, training_labels), (test_images, test_labels) = mnist.load_data()
plt.imshow(training_images[0])
print(training_labels[0])
print(training_images[0])

# normilise (0..255 -> 0..1)
training_images = training_images / 255.0
test_images = test_images / 255.0

model = tf.keras.models.Sequential([  # That defines a SEQUENCE of layers in the neural network
    tf.keras.layers.Flatten(),  # Flatten just takes that square and turns it into a 1 dimensional set.
    # Adds a layer of neurons with activation function Relu - effectively means "If X>0 return X, else return 0" -
    # - so what it does it only passes values 0 or greater to the next layer in the network.
    tf.keras.layers.Dense(1024, activation=tf.nn.relu),
    # Adds a layer with activation function Softman - takes a set of values, and effectively picks the biggest one,
    # so, for example, if the output of the last layer looks like [0.1, 0.1, 0.05, 0.1, 9.5, 0.1, 0.05, 0.05, 0.05],
    # it saves you from fishing through it looking for the biggest value, and turns it into [0,0,0,0,1,0,0,0,0] -- The goal is to save a lot of coding!
    tf.keras.layers.Dense(10, activation=tf.nn.softmax)
])

model.compile(optimizer=tf.optimizers.Adam(), loss="sparse_categorical_crossentropy", metrics=['accuracy'])
model.fit(training_images, training_labels, epochs=10, callbacks=[myCallback()])
model.evaluate(test_images, test_labels)

classifications = model.predict(test_images)
print(classifications[0])
print(test_labels[0])

""" 
Week 2 task

# GRADED FUNCTION: train_mnist
def train_mnist():
    # Please write your code only where you are indicated.
    # please do not remove # model fitting inline comments.

    # YOUR CODE SHOULD START HERE
    class myCallback(tf.keras.callbacks.Callback):
        def on_epoch_end(self, epoch, logs={}):
            loss = logs.get('loss')
            print(loss)
            if loss < 0.01:
                print("\nReached 99% accuracy so cancelling training!")
                self.model.stop_training = True
    # YOUR CODE SHOULD END HERE

    mnist = tf.keras.datasets.mnist

    (x_train, y_train), (x_test, y_test) = mnist.load_data()
    # YOUR CODE SHOULD START HERE
    x_test = y_test / 255.0
    # YOUR CODE SHOULD END HERE
    model = tf.keras.models.Sequential([
        # YOUR CODE SHOULD START HERE
        tf.keras.layers.Flatten(),
        tf.keras.layers.Dense(512, activation=tf.nn.relu),
        tf.keras.layers.Dense(10, activation=tf.nn.softmax)
        # YOUR CODE SHOULD END HERE
    ])

    model.compile(optimizer='adam',
                  loss='sparse_categorical_crossentropy',
                  metrics=['accuracy'])

    # model fitting
    history = model.fit(  # YOUR CODE SHOULD START HERE
        x_test, y_test, epochs=10, callbacks=[myCallback()]
        # YOUR CODE SHOULD END HERE
    )
    # model fitting
    return history.epoch, history.history['acc'][-1]

train_mnist()
"""
