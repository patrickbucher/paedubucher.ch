---
title: Jahresrückblick 2021
subtitle: Gar nicht so schlecht?
author: Patrick Bucher
date: 2021-12-17T13:00:00
lang: de
---

Selten war meine Laune am Jahresende so schlecht. Dass ich am Jahresende
erschöpft bin und mich nach einer Pause sehne, ist nichts Neues, und war zu
erwarten. Leider habe ich aber auch das Gefühl, dass dieses Jahr für mich extrem
unproduktiv war, und ich stehenbleibe, ja gar Rückschritte mache ‒ in allen
möglichen Lebensbereichen.

Ich möchte hier auf mein Jahr zurückblicken, soweit es mein Erinnerungsvermögen
erlaubt, um diesen Eindruck zu prüfen.

Was Fremdsprachen und Literatur betrifft, ist der Rückblick schnell abgehandelt.
Französisch und Russisch habe ich kaum verwendet; Englisch höchstens für die
Informatik. An Literatur kann ich mich dieses Jahr kaum erinnern. Ayn Rands
_Anthem_ habe ich auf dem Kindle gelesen. Viele andere Romane habe ich wohl
_an-_ aber nicht _durchgelesen_. Den Kindle habe ich schliesslich wieder in den
Keller verbannt, da ich damit nur oberflächlich lese.

Was meinen Beruf und die Informatik angeht, weiss ich doch etwas mehr zu
berichten.

# Motiviert mit MOOCs

Ich weiss nicht, ob es Ende 2020 oder Anfang 2021 war, als ich den Kurs [Machine
Learning](https://www.coursera.org/learn/machine-learning?) abgeschlossen habe,
aber mit der [Machine Learning
Specialization](https://www.coursera.org/specializations/deep-learning) habe ich
wohl erst 2021 angefangen. Hier bin ich schnell vorangekommen und war lange Zeit
motiviert ‒ bis es schliesslich in einem Mooc zum Framework
[TensorFlow](https://www.tensorflow.org/) kam, das immer noch in der Version 1.x
verwendet worden ist. Der Umstieg auf TensorFlow 2.x muss wohl ein grösseres
Problem sein, zumindest ist er mir nicht einmal für die Kursbeispiele gelungen.

Immerhin habe ich selber noch einige Beispiele entwickelt:
[x-o-classifier](https://github.com/patrickbucher/x-o-classifier) zur optischen
Unterscheidung der Zeichen `x` und `o`,
[digit-detection](https://github.com/patrickbucher/digit-detection) als
erweitertes OCR-Beispiel für die Ziffern von 0 bis 9, die [Titanic
Challenge](https://github.com/patrickbucher/titanic-challenge) und die [House
Prices](https://github.com/patrickbucher/house-prices) von Kaggle ‒ wenn auch
mit mässigem Erfolg). Weiter sind mir kleinere Beispiele und eine
Zusammenfassung des ersten Kurses in meinem
[machine-learning-Repository](https://github.com/patrickbucher/machine-learning)
geblieben.

Für mich war das Thema Machine Learning dann vorerst erledigt.

# Funktionale Programmierung

Weiter ging es mit der funktionalen Programmierung, wozu ich den recht
gelungenen
[Scala-Kurs](https://de.coursera.org/learn/scala-functional-programming) von
Martin Odersky durchgearbeitet habe. Der meiner Meinung nach am besten
aufgebaute Kurs zu funktionaler Programmierung, [Programming Languages, Part
A](https://www.coursera.org/learn/programming-languages), konnte ich leider
nicht abschliessen: das Lernen via Videos hat mich zu sehr erschöpft, und ich
wollte weg vom Bildschirm.

So machte ich mich an
[SICP](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book.html)
heran, worin ich leider im ersten Kapitel steckengeblieben bin. Immerhin habe
ich die Übungen bis dahin [seriös
gelöst](https://github.com/patrickbucher/sicp/tree/master/ch01), sodass ich hier
bei späterer Gelegenheit wieder anküpfen kann.

Eigentlich wollte ich mich 2021 mit Erlang befassen, bin aber dann aus
pragmatischen Gründen doch auf Elixir eingeschwenkt. Ich versuchte das Buch
[Learn Functional Programming with
Elixir](https://pragprog.com/titles/cdc-elixir/learn-functional-programming-with-elixir/)
durchzuarbeiten, gab aber auf Seite 129 (von 198) auf, wovon der Kommentar ganz
unten in meinen
[Notizen](https://github.com/patrickbucher/learning-elixir/blob/master/learn-functional-programming-with-elixir/notes.md)
zeugt. Die Beispiele waren mir wohl einfach zu kindisch und praxisfern. Ich
versuchte es dann noch mit [Elixir in
Action](https://www.manning.com/books/elixir-in-action-second-edition), doch
hatte ich im Frühling dann andere Prioritäten.

# Servermigrationen

In der Firma habe ich mich die ersten vier Monate v.a. mit der Migration von
Servern von Ubuntu 16.04 auf 20.04 befasst. U.a. musste ich einen Mailserver und
den ganzen Monitoring-Stack migrieren, was wohl das Schwierigste an der Aufgabe
war.

Für die Ingestion der Metriken via [Riemann](http://riemann.io/) musste ich, um
von Python 2 wegzukommen, von einer bestehenden Datenbank auf
[InfluxDB](https://www.influxdata.com/) wechseln. Hierzu durfte ich sogar etwas
Clojure lernen, wobei mir [Getting
Clojure](https://pragprog.com/titles/roclojure/getting-clojure/) von Russ Olsen
geholfen hat. Ich benötigte nur ca. zwei Kapitel. Im Frühling arbeitete ich das
Buch dann komplett durch, wovon mein Repository
[learning-clojure](https://github.com/patrickbucher/learning-clojure/) zeugt.

